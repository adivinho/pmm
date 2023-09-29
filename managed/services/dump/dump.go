package dump

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm/managed/models"
)

var ErrDumpAlreadyRunning = errors.New("pmm-dump already running")

const (
	pmmDumpBin = "pmm-dump"
	dumpsDir   = "/srv/dump/"
)

type Service struct {
	l *logrus.Entry

	db *reform.DB

	running atomic.Bool

	rw     sync.RWMutex
	cancel context.CancelFunc
}

func New(db *reform.DB) *Service {
	return &Service{
		l:  logrus.WithField("component", "management/backup/backup"),
		db: db,
	}
}

type Params struct {
	StartTime  time.Time
	EndTime    time.Time
	ExportQAN  bool
	IgnoreLoad bool
}

func (s *Service) StartDump(params *Params) (string, error) {
	// Check is some dump already running
	if !s.running.CompareAndSwap(false, true) {
		return "", ErrDumpAlreadyRunning
	}

	dump, err := models.CreateDump(s.db.Querier, models.CreateDumpParams{
		StartTime:  params.StartTime,
		EndTime:    params.EndTime,
		ExportQAN:  params.ExportQAN,
		IgnoreLoad: params.IgnoreLoad,
	})
	if err != nil {
		s.running.Store(false)
		return "", errors.Wrap(err, "failed to create dump")
	}

	ctx, cancel := context.WithCancel(context.Background())

	s.rw.Lock()
	s.cancel = cancel
	s.rw.Unlock()

	pmmDumpCmd := exec.CommandContext(ctx,
		pmmDumpBin,
		"export",
		`--pmm-url "http://127.0.0.1"`,
		"--pmm-user admin", // TODO set valid user
		"--pmm-pass admin", // TODO and password
		"--dump-path="+dumpsDir+dump.ID,
	)

	if !params.StartTime.IsZero() {
		pmmDumpCmd.Args = append(pmmDumpCmd.Args, fmt.Sprintf("--start-ts=%s", params.StartTime.Format(time.RFC3339)))
	}

	if !params.EndTime.IsZero() {
		pmmDumpCmd.Args = append(pmmDumpCmd.Args, fmt.Sprintf("--end-ts=%s", params.EndTime.Format(time.RFC3339)))
	}

	if params.ExportQAN {
		pmmDumpCmd.Args = append(pmmDumpCmd.Args, "--dump-qan")
	}

	if params.IgnoreLoad {
		pmmDumpCmd.Args = append(pmmDumpCmd.Args, "--ignore-load")
	}
	var output bytes.Buffer
	pmmDumpCmd.Stdout = &output
	pmmDumpCmd.Stderr = &output

	go func() {
		err := s.persistLogs(ctx, dump.ID, &output)
		if err != nil && !errors.Is(err, context.Canceled) {
			s.l.Errorf("pmm-dupm logs persist failed: %v", err)
		}
	}()

	go func() {
		// Switch running flag back to false
		defer s.running.Store(false)

		if err := pmmDumpCmd.Run(); err != nil {
			s.l.Errorf("failed to execute pmm-dump: %v", err)
		}
	}()

	return dump.ID, nil
}

func (s *Service) persistLogs(ctx context.Context, dumpID string, buf *bytes.Buffer) error {
	scanner := bufio.NewScanner(buf)
	var err error
	var chunkN int
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// continue
		}

		if scanner.Scan() {
			_, err = models.CreateDumpLog(s.db.Querier, models.CreateDumpLogParams{
				DumpID:    dumpID,
				ChunkID:   chunkN,
				Data:      scanner.Text(),
				LastChunk: false,
			})

			if err != nil {
				return errors.Wrap(err, "failed to save pmm-dump log chunk")
			}
		}

		if err := scanner.Err(); err != nil {
			s.l.Warnf("failed to read pmm-dupm logs: %v", err)
		}

		_, err = models.CreateDumpLog(s.db.Querier, models.CreateDumpLogParams{
			DumpID:    dumpID,
			ChunkID:   chunkN,
			Data:      "",
			LastChunk: true,
		})

		if err != nil {
			return errors.Wrap(err, "failed to save pmm-dump last log chunk")
		}
	}
}

func (s *Service) StopDump() {
	s.rw.RLock()
	defer s.rw.RUnlock()

	s.cancel()
}
