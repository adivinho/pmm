// Code generated by mockery v2.35.1. DO NOT EDIT.

package dump

import (
	mock "github.com/stretchr/testify/mock"

	servicesdump "github.com/percona/pmm/managed/services/dump"
)

// mockDumpService is an autogenerated mock type for the dumpService type
type mockDumpService struct {
	mock.Mock
}

// StartDump provides a mock function with given fields: params
func (_m *mockDumpService) StartDump(params *servicesdump.Params) (string, error) {
	ret := _m.Called(params)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*servicesdump.Params) (string, error)); ok {
		return rf(params)
	}
	if rf, ok := ret.Get(0).(func(*servicesdump.Params) string); ok {
		r0 = rf(params)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*servicesdump.Params) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockDumpService creates a new instance of mockDumpService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockDumpService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockDumpService {
	mock := &mockDumpService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
