// Code generated by mockery v1.0.0. DO NOT EDIT.

package run

import (
	context "context"
	io "io"
	time "time"

	types "github.com/docker/docker/api/types"
	container "github.com/docker/docker/api/types/container"
	client "github.com/docker/docker/client"
	mock "github.com/stretchr/testify/mock"

	docker "github.com/percona/pmm/admin/pkg/docker"
)

// mockContainerManager is an autogenerated mock type for the containerManager type
type mockContainerManager struct {
	mock.Mock
}

// ContainerInspect provides a mock function with given fields: ctx, containerID
func (_m *mockContainerManager) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error) {
	ret := _m.Called(ctx, containerID)

	var r0 types.ContainerJSON
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ContainerJSON); ok {
		r0 = rf(ctx, containerID)
	} else {
		r0 = ret.Get(0).(types.ContainerJSON)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, containerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerList provides a mock function with given fields: ctx, options
func (_m *mockContainerManager) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	ret := _m.Called(ctx, options)

	var r0 []types.Container
	if rf, ok := ret.Get(0).(func(context.Context, types.ContainerListOptions) []types.Container); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.ContainerListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerStop provides a mock function with given fields: ctx, containerID, timeout
func (_m *mockContainerManager) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error {
	ret := _m.Called(ctx, containerID, timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *time.Duration) error); ok {
		r0 = rf(ctx, containerID, timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerUpdate provides a mock function with given fields: ctx, containerID, updateConfig
func (_m *mockContainerManager) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error) {
	ret := _m.Called(ctx, containerID, updateConfig)

	var r0 container.ContainerUpdateOKBody
	if rf, ok := ret.Get(0).(func(context.Context, string, container.UpdateConfig) container.ContainerUpdateOKBody); ok {
		r0 = rf(ctx, containerID, updateConfig)
	} else {
		r0 = ret.Get(0).(container.ContainerUpdateOKBody)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, container.UpdateConfig) error); ok {
		r1 = rf(ctx, containerID, updateConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindServerContainers provides a mock function with given fields: ctx
func (_m *mockContainerManager) FindServerContainers(ctx context.Context) ([]types.Container, error) {
	ret := _m.Called(ctx)

	var r0 []types.Container
	if rf, ok := ret.Get(0).(func(context.Context) []types.Container); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDockerClient provides a mock function with given fields:
func (_m *mockContainerManager) GetDockerClient() *client.Client {
	ret := _m.Called()

	var r0 *client.Client
	if rf, ok := ret.Get(0).(func() *client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Client)
		}
	}

	return r0
}

// HaveDockerAccess provides a mock function with given fields: ctx
func (_m *mockContainerManager) HaveDockerAccess(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ImageInspectWithRaw provides a mock function with given fields: ctx, imageID
func (_m *mockContainerManager) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error) {
	ret := _m.Called(ctx, imageID)

	var r0 types.ImageInspect
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ImageInspect); ok {
		r0 = rf(ctx, imageID)
	} else {
		r0 = ret.Get(0).(types.ImageInspect)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string) []byte); ok {
		r1 = rf(ctx, imageID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, imageID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PullImage provides a mock function with given fields: ctx, dockerImage, opts
func (_m *mockContainerManager) PullImage(ctx context.Context, dockerImage string, opts types.ImagePullOptions) (io.Reader, error) {
	ret := _m.Called(ctx, dockerImage, opts)

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImagePullOptions) io.Reader); ok {
		r0 = rf(ctx, dockerImage, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ImagePullOptions) error); ok {
		r1 = rf(ctx, dockerImage, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunContainer provides a mock function with given fields: ctx, config, hostConfig, containerName
func (_m *mockContainerManager) RunContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, containerName string) (string, error) {
	ret := _m.Called(ctx, config, hostConfig, containerName)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *container.Config, *container.HostConfig, string) string); ok {
		r0 = rf(ctx, config, hostConfig, containerName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *container.Config, *container.HostConfig, string) error); ok {
		r1 = rf(ctx, config, hostConfig, containerName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForHealthyContainer provides a mock function with given fields: ctx, containerID
func (_m *mockContainerManager) WaitForHealthyContainer(ctx context.Context, containerID string) <-chan docker.WaitHealthyResponse {
	ret := _m.Called(ctx, containerID)

	var r0 <-chan docker.WaitHealthyResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) <-chan docker.WaitHealthyResponse); ok {
		r0 = rf(ctx, containerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan docker.WaitHealthyResponse)
		}
	}

	return r0
}