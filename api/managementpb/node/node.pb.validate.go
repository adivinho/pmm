// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: managementpb/node/node.proto

package nodev1beta1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	inventorypb "github.com/percona/pmm/api/inventorypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort

	_ = inventorypb.NodeType(0)
)

// Validate checks the field values on UniversalNode with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UniversalNode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UniversalNode with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UniversalNodeMultiError, or
// nil if none found.
func (m *UniversalNode) ValidateAll() error {
	return m.validate(true)
}

func (m *UniversalNode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NodeId

	// no validation rules for NodeType

	// no validation rules for NodeName

	// no validation rules for MachineId

	// no validation rules for Distro

	// no validation rules for NodeModel

	// no validation rules for ContainerId

	// no validation rules for ContainerName

	// no validation rules for Address

	// no validation rules for Region

	// no validation rules for Az

	// no validation rules for CustomLabels

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UniversalNodeValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UniversalNodeValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UniversalNodeValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UniversalNodeValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UniversalNodeValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UniversalNodeValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Status

	for idx, item := range m.GetAgents() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UniversalNodeValidationError{
						field:  fmt.Sprintf("Agents[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UniversalNodeValidationError{
						field:  fmt.Sprintf("Agents[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UniversalNodeValidationError{
					field:  fmt.Sprintf("Agents[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetServices() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UniversalNodeValidationError{
						field:  fmt.Sprintf("Services[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UniversalNodeValidationError{
						field:  fmt.Sprintf("Services[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UniversalNodeValidationError{
					field:  fmt.Sprintf("Services[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UniversalNodeMultiError(errors)
	}

	return nil
}

// UniversalNodeMultiError is an error wrapping multiple validation errors
// returned by UniversalNode.ValidateAll() if the designated constraints
// aren't met.
type UniversalNodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UniversalNodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UniversalNodeMultiError) AllErrors() []error { return m }

// UniversalNodeValidationError is the validation error returned by
// UniversalNode.Validate if the designated constraints aren't met.
type UniversalNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UniversalNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UniversalNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UniversalNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UniversalNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UniversalNodeValidationError) ErrorName() string { return "UniversalNodeValidationError" }

// Error satisfies the builtin error interface
func (e UniversalNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUniversalNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UniversalNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UniversalNodeValidationError{}

// Validate checks the field values on ListNodeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListNodeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListNodeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListNodeRequestMultiError, or nil if none found.
func (m *ListNodeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListNodeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NodeType

	if len(errors) > 0 {
		return ListNodeRequestMultiError(errors)
	}

	return nil
}

// ListNodeRequestMultiError is an error wrapping multiple validation errors
// returned by ListNodeRequest.ValidateAll() if the designated constraints
// aren't met.
type ListNodeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListNodeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListNodeRequestMultiError) AllErrors() []error { return m }

// ListNodeRequestValidationError is the validation error returned by
// ListNodeRequest.Validate if the designated constraints aren't met.
type ListNodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNodeRequestValidationError) ErrorName() string { return "ListNodeRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListNodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNodeRequestValidationError{}

// Validate checks the field values on ListNodeResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListNodeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListNodeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListNodeResponseMultiError, or nil if none found.
func (m *ListNodeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListNodeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetNodes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListNodeResponseValidationError{
						field:  fmt.Sprintf("Nodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListNodeResponseValidationError{
						field:  fmt.Sprintf("Nodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListNodeResponseValidationError{
					field:  fmt.Sprintf("Nodes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListNodeResponseMultiError(errors)
	}

	return nil
}

// ListNodeResponseMultiError is an error wrapping multiple validation errors
// returned by ListNodeResponse.ValidateAll() if the designated constraints
// aren't met.
type ListNodeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListNodeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListNodeResponseMultiError) AllErrors() []error { return m }

// ListNodeResponseValidationError is the validation error returned by
// ListNodeResponse.Validate if the designated constraints aren't met.
type ListNodeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNodeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNodeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNodeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNodeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNodeResponseValidationError) ErrorName() string { return "ListNodeResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListNodeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNodeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNodeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNodeResponseValidationError{}

// Validate checks the field values on GetNodeRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetNodeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNodeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetNodeRequestMultiError,
// or nil if none found.
func (m *GetNodeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNodeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NodeId

	if len(errors) > 0 {
		return GetNodeRequestMultiError(errors)
	}

	return nil
}

// GetNodeRequestMultiError is an error wrapping multiple validation errors
// returned by GetNodeRequest.ValidateAll() if the designated constraints
// aren't met.
type GetNodeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNodeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNodeRequestMultiError) AllErrors() []error { return m }

// GetNodeRequestValidationError is the validation error returned by
// GetNodeRequest.Validate if the designated constraints aren't met.
type GetNodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNodeRequestValidationError) ErrorName() string { return "GetNodeRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetNodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNodeRequestValidationError{}

// Validate checks the field values on GetNodeResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetNodeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNodeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNodeResponseMultiError, or nil if none found.
func (m *GetNodeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNodeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetNodeResponseValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetNodeResponseValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNodeResponseValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetNodeResponseMultiError(errors)
	}

	return nil
}

// GetNodeResponseMultiError is an error wrapping multiple validation errors
// returned by GetNodeResponse.ValidateAll() if the designated constraints
// aren't met.
type GetNodeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNodeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNodeResponseMultiError) AllErrors() []error { return m }

// GetNodeResponseValidationError is the validation error returned by
// GetNodeResponse.Validate if the designated constraints aren't met.
type GetNodeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNodeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNodeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNodeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNodeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNodeResponseValidationError) ErrorName() string { return "GetNodeResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetNodeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNodeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNodeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNodeResponseValidationError{}

// Validate checks the field values on UniversalNode_Service with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UniversalNode_Service) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UniversalNode_Service with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UniversalNode_ServiceMultiError, or nil if none found.
func (m *UniversalNode_Service) ValidateAll() error {
	return m.validate(true)
}

func (m *UniversalNode_Service) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ServiceId

	// no validation rules for ServiceType

	// no validation rules for ServiceName

	if len(errors) > 0 {
		return UniversalNode_ServiceMultiError(errors)
	}

	return nil
}

// UniversalNode_ServiceMultiError is an error wrapping multiple validation
// errors returned by UniversalNode_Service.ValidateAll() if the designated
// constraints aren't met.
type UniversalNode_ServiceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UniversalNode_ServiceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UniversalNode_ServiceMultiError) AllErrors() []error { return m }

// UniversalNode_ServiceValidationError is the validation error returned by
// UniversalNode_Service.Validate if the designated constraints aren't met.
type UniversalNode_ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UniversalNode_ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UniversalNode_ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UniversalNode_ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UniversalNode_ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UniversalNode_ServiceValidationError) ErrorName() string {
	return "UniversalNode_ServiceValidationError"
}

// Error satisfies the builtin error interface
func (e UniversalNode_ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUniversalNode_Service.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UniversalNode_ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UniversalNode_ServiceValidationError{}

// Validate checks the field values on UniversalNode_Agent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UniversalNode_Agent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UniversalNode_Agent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UniversalNode_AgentMultiError, or nil if none found.
func (m *UniversalNode_Agent) ValidateAll() error {
	return m.validate(true)
}

func (m *UniversalNode_Agent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AgentId

	// no validation rules for AgentType

	// no validation rules for Status

	// no validation rules for IsConnected

	if len(errors) > 0 {
		return UniversalNode_AgentMultiError(errors)
	}

	return nil
}

// UniversalNode_AgentMultiError is an error wrapping multiple validation
// errors returned by UniversalNode_Agent.ValidateAll() if the designated
// constraints aren't met.
type UniversalNode_AgentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UniversalNode_AgentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UniversalNode_AgentMultiError) AllErrors() []error { return m }

// UniversalNode_AgentValidationError is the validation error returned by
// UniversalNode_Agent.Validate if the designated constraints aren't met.
type UniversalNode_AgentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UniversalNode_AgentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UniversalNode_AgentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UniversalNode_AgentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UniversalNode_AgentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UniversalNode_AgentValidationError) ErrorName() string {
	return "UniversalNode_AgentValidationError"
}

// Error satisfies the builtin error interface
func (e UniversalNode_AgentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUniversalNode_Agent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UniversalNode_AgentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UniversalNode_AgentValidationError{}
