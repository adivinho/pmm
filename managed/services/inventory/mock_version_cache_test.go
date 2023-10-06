// Code generated by mockery v2.35.1. DO NOT EDIT.

package inventory

import mock "github.com/stretchr/testify/mock"

// mockVersionCache is an autogenerated mock type for the versionCache type
type mockVersionCache struct {
	mock.Mock
}

// RequestSoftwareVersionsUpdate provides a mock function with given fields:
func (_m *mockVersionCache) RequestSoftwareVersionsUpdate() {
	_m.Called()
}

// newMockVersionCache creates a new instance of mockVersionCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockVersionCache(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockVersionCache {
	mock := &mockVersionCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
