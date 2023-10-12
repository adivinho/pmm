// Code generated by mockery v2.35.1. DO NOT EDIT.

package backup

import (
	mock "github.com/stretchr/testify/mock"

	agents "github.com/percona/pmm/managed/services/agents"
)

// mockVersioner is an autogenerated mock type for the versioner type
type mockVersioner struct {
	mock.Mock
}

// GetVersions provides a mock function with given fields: pmmAgentID, softwares
func (_m *mockVersioner) GetVersions(pmmAgentID string, softwares []agents.Software) ([]agents.Version, error) {
	ret := _m.Called(pmmAgentID, softwares)

	var r0 []agents.Version
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []agents.Software) ([]agents.Version, error)); ok {
		return rf(pmmAgentID, softwares)
	}
	if rf, ok := ret.Get(0).(func(string, []agents.Software) []agents.Version); ok {
		r0 = rf(pmmAgentID, softwares)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]agents.Version)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []agents.Software) error); ok {
		r1 = rf(pmmAgentID, softwares)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockVersioner creates a new instance of mockVersioner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockVersioner(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockVersioner {
	mock := &mockVersioner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
