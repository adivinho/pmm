// Copyright (C) 2023 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

// DumpFilters represents filters for artifacts list.
type DumpFilters struct {
	// Return only artifacts by specified status.
	Status BackupStatus
}

type CreateDumpParams struct {
	StartTime  time.Time
	EndTime    time.Time
	ExportQAN  bool
	IgnoreLoad bool
}

func CreateDump(q *reform.Querier, params CreateDumpParams) (*Dump, error) {
	dump := &Dump{
		ID:         "/job_id/" + uuid.New().String(),
		Status:     DumpStatusInProgress, // TODO ?
		NodeIDs:    nil,                  // TODO
		StartTime:  params.StartTime,
		EndTime:    params.EndTime,
		ExportQAN:  params.ExportQAN,
		IgnoreLoad: params.IgnoreLoad,
	}
	if err := q.Insert(dump); err != nil {
		return nil, errors.WithStack(err)
	}
	return dump, nil
}

// FindDumps returns dumps list sorted by creation time in DESCENDING order.
func FindDumps(q *reform.Querier, filters DumpFilters) ([]*Dump, error) {
	var conditions []string
	var args []interface{}
	var idx int

	if filters.Status != "" {
		idx++
		conditions = append(conditions, fmt.Sprintf("status = %s", q.Placeholder(idx)))
		args = append(args, filters.Status)
	}

	var whereClause string
	if len(conditions) != 0 {
		whereClause = fmt.Sprintf("WHERE %s", strings.Join(conditions, " AND "))
	}
	rows, err := q.SelectAllFrom(DumpTable, fmt.Sprintf("%s ORDER BY created_at DESC", whereClause), args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select dumps")
	}

	dumps := make([]*Dump, 0, len(rows))
	for _, r := range rows {
		dumps = append(dumps, r.(*Dump)) //nolint:forcetypeassert
	}

	return dumps, nil
}

// FindDumpByID returns dump by given ID if found, ErrNotFound if not.
func FindDumpByID(q *reform.Querier, id string) (*Dump, error) {
	if id == "" {
		return nil, errors.New("provided dump id is empty")
	}

	dump := &Dump{ID: id}
	err := q.Reload(dump)
	if err != nil {
		if errors.Is(err, reform.ErrNoRows) {
			return nil, errors.Wrapf(ErrNotFound, "dump by id '%s'", id)
		}
		return nil, errors.WithStack(err)
	}

	return dump, nil
}

// DeleteDump removes dump by ID.
func DeleteDump(q *reform.Querier, id string) error {
	if _, err := FindDumpByID(q, id); err != nil {
		return err
	}

	if err := q.Delete(&Dump{ID: id}); err != nil {
		return errors.Wrapf(err, "failed to delete dump by id '%s'", id)
	}
	return nil
}

// IsDumpFinalStatus checks if dump status is one of the final ones.
func IsDumpFinalStatus(dumpStatus DumpStatus) bool {
	switch dumpStatus {
	case DumpStatusSuccess, DumpStatusError:
		return true
	default:
		return false
	}
}

// CreateDumpLogParams are params for creating a new pmm-dump log.
type CreateDumpLogParams struct {
	DumpID    string
	ChunkID   int
	Data      string
	LastChunk bool
}

// CreateDumpLog inserts new chunk log.
func CreateDumpLog(q *reform.Querier, params CreateDumpLogParams) (*DumpLog, error) {
	log := &DumpLog{
		DumpID:    params.DumpID,
		ChunkID:   params.ChunkID,
		Data:      params.Data,
		LastChunk: params.LastChunk,
	}
	if err := q.Insert(log); err != nil {
		return nil, errors.WithStack(err)
	}
	return log, nil
}
