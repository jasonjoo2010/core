// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import lock "github.com/projecteru2/core/lock"
import mock "github.com/stretchr/testify/mock"

import types "github.com/projecteru2/core/types"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AddContainer provides a mock function with given fields: ctx, container
func (_m *Store) AddContainer(ctx context.Context, container *types.Container) error {
	ret := _m.Called(ctx, container)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Container) error); ok {
		r0 = rf(ctx, container)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddNode provides a mock function with given fields: ctx, name, endpoint, podname, ca, cert, key, cpu, share, memory, labels
func (_m *Store) AddNode(ctx context.Context, name string, endpoint string, podname string, ca string, cert string, key string, cpu int, share int, memory int64, labels map[string]string) (*types.Node, error) {
	ret := _m.Called(ctx, name, endpoint, podname, ca, cert, key, cpu, share, memory, labels)

	var r0 *types.Node
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string, string, int, int, int64, map[string]string) *types.Node); ok {
		r0 = rf(ctx, name, endpoint, podname, ca, cert, key, cpu, share, memory, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, string, string, int, int, int64, map[string]string) error); ok {
		r1 = rf(ctx, name, endpoint, podname, ca, cert, key, cpu, share, memory, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddPod provides a mock function with given fields: ctx, name, favor, desc
func (_m *Store) AddPod(ctx context.Context, name string, favor string, desc string) (*types.Pod, error) {
	ret := _m.Called(ctx, name, favor, desc)

	var r0 *types.Pod
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *types.Pod); ok {
		r0 = rf(ctx, name, favor, desc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, favor, desc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CleanContainerData provides a mock function with given fields: ctx, ID, appname, entrypoint, nodename
func (_m *Store) CleanContainerData(ctx context.Context, ID string, appname string, entrypoint string, nodename string) error {
	ret := _m.Called(ctx, ID, appname, entrypoint, nodename)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) error); ok {
		r0 = rf(ctx, ID, appname, entrypoint, nodename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerDeployed provides a mock function with given fields: ctx, ID, appname, entrypoint, nodename, data
func (_m *Store) ContainerDeployed(ctx context.Context, ID string, appname string, entrypoint string, nodename string, data string) error {
	ret := _m.Called(ctx, ID, appname, entrypoint, nodename, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) error); ok {
		r0 = rf(ctx, ID, appname, entrypoint, nodename, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateLock provides a mock function with given fields: key, ttl
func (_m *Store) CreateLock(key string, ttl int) (lock.DistributedLock, error) {
	ret := _m.Called(key, ttl)

	var r0 lock.DistributedLock
	if rf, ok := ret.Get(0).(func(string, int) lock.DistributedLock); ok {
		r0 = rf(key, ttl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(lock.DistributedLock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(key, ttl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNode provides a mock function with given fields: ctx, node
func (_m *Store) DeleteNode(ctx context.Context, node *types.Node) {
	_m.Called(ctx, node)
}

// DeleteProcessing provides a mock function with given fields: ctx, opts, nodeInfo
func (_m *Store) DeleteProcessing(ctx context.Context, opts *types.DeployOptions, nodeInfo types.NodeInfo) error {
	ret := _m.Called(ctx, opts, nodeInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.DeployOptions, types.NodeInfo) error); ok {
		r0 = rf(ctx, opts, nodeInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllNodes provides a mock function with given fields: ctx
func (_m *Store) GetAllNodes(ctx context.Context) ([]*types.Node, error) {
	ret := _m.Called(ctx)

	var r0 []*types.Node
	if rf, ok := ret.Get(0).(func(context.Context) []*types.Node); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Node)
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

// GetAllPods provides a mock function with given fields: ctx
func (_m *Store) GetAllPods(ctx context.Context) ([]*types.Pod, error) {
	ret := _m.Called(ctx)

	var r0 []*types.Pod
	if rf, ok := ret.Get(0).(func(context.Context) []*types.Pod); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Pod)
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

// GetContainer provides a mock function with given fields: ctx, ID
func (_m *Store) GetContainer(ctx context.Context, ID string) (*types.Container, error) {
	ret := _m.Called(ctx, ID)

	var r0 *types.Container
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Container); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContainers provides a mock function with given fields: ctx, IDs
func (_m *Store) GetContainers(ctx context.Context, IDs []string) ([]*types.Container, error) {
	ret := _m.Called(ctx, IDs)

	var r0 []*types.Container
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*types.Container); ok {
		r0 = rf(ctx, IDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, IDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNode provides a mock function with given fields: ctx, podname, nodename
func (_m *Store) GetNode(ctx context.Context, podname string, nodename string) (*types.Node, error) {
	ret := _m.Called(ctx, podname, nodename)

	var r0 *types.Node
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *types.Node); ok {
		r0 = rf(ctx, podname, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, podname, nodename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeByName provides a mock function with given fields: ctx, nodename
func (_m *Store) GetNodeByName(ctx context.Context, nodename string) (*types.Node, error) {
	ret := _m.Called(ctx, nodename)

	var r0 *types.Node
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Node); ok {
		r0 = rf(ctx, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nodename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodesByPod provides a mock function with given fields: ctx, podname
func (_m *Store) GetNodesByPod(ctx context.Context, podname string) ([]*types.Node, error) {
	ret := _m.Called(ctx, podname)

	var r0 []*types.Node
	if rf, ok := ret.Get(0).(func(context.Context, string) []*types.Node); ok {
		r0 = rf(ctx, podname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, podname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPod provides a mock function with given fields: ctx, podname
func (_m *Store) GetPod(ctx context.Context, podname string) (*types.Pod, error) {
	ret := _m.Called(ctx, podname)

	var r0 *types.Pod
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Pod); ok {
		r0 = rf(ctx, podname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, podname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListContainers provides a mock function with given fields: ctx, appname, entrypoint, nodename
func (_m *Store) ListContainers(ctx context.Context, appname string, entrypoint string, nodename string) ([]*types.Container, error) {
	ret := _m.Called(ctx, appname, entrypoint, nodename)

	var r0 []*types.Container
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) []*types.Container); ok {
		r0 = rf(ctx, appname, entrypoint, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, appname, entrypoint, nodename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNodeContainers provides a mock function with given fields: ctx, nodename
func (_m *Store) ListNodeContainers(ctx context.Context, nodename string) ([]*types.Container, error) {
	ret := _m.Called(ctx, nodename)

	var r0 []*types.Container
	if rf, ok := ret.Get(0).(func(context.Context, string) []*types.Container); ok {
		r0 = rf(ctx, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nodename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeDeployStatus provides a mock function with given fields: ctx, opts, nodesInfo
func (_m *Store) MakeDeployStatus(ctx context.Context, opts *types.DeployOptions, nodesInfo []types.NodeInfo) ([]types.NodeInfo, error) {
	ret := _m.Called(ctx, opts, nodesInfo)

	var r0 []types.NodeInfo
	if rf, ok := ret.Get(0).(func(context.Context, *types.DeployOptions, []types.NodeInfo) []types.NodeInfo); ok {
		r0 = rf(ctx, opts, nodesInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.NodeInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.DeployOptions, []types.NodeInfo) error); ok {
		r1 = rf(ctx, opts, nodesInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveContainer provides a mock function with given fields: ctx, container
func (_m *Store) RemoveContainer(ctx context.Context, container *types.Container) error {
	ret := _m.Called(ctx, container)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Container) error); ok {
		r0 = rf(ctx, container)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemovePod provides a mock function with given fields: ctx, podname
func (_m *Store) RemovePod(ctx context.Context, podname string) error {
	ret := _m.Called(ctx, podname)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, podname)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveProcessing provides a mock function with given fields: ctx, opts, nodeInfo
func (_m *Store) SaveProcessing(ctx context.Context, opts *types.DeployOptions, nodeInfo types.NodeInfo) error {
	ret := _m.Called(ctx, opts, nodeInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.DeployOptions, types.NodeInfo) error); ok {
		r0 = rf(ctx, opts, nodeInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateContainer provides a mock function with given fields: ctx, container
func (_m *Store) UpdateContainer(ctx context.Context, container *types.Container) error {
	ret := _m.Called(ctx, container)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Container) error); ok {
		r0 = rf(ctx, container)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNode provides a mock function with given fields: ctx, node
func (_m *Store) UpdateNode(ctx context.Context, node *types.Node) error {
	ret := _m.Called(ctx, node)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Node) error); ok {
		r0 = rf(ctx, node)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNodeResource provides a mock function with given fields: ctx, node, cpu, quota, mem, action
func (_m *Store) UpdateNodeResource(ctx context.Context, node *types.Node, cpu types.CPUMap, quota float64, mem int64, action string) error {
	ret := _m.Called(ctx, node, cpu, quota, mem, action)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Node, types.CPUMap, float64, int64, string) error); ok {
		r0 = rf(ctx, node, cpu, quota, mem, action)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProcessing provides a mock function with given fields: ctx, opts, nodename, count
func (_m *Store) UpdateProcessing(ctx context.Context, opts *types.DeployOptions, nodename string, count int) error {
	ret := _m.Called(ctx, opts, nodename, count)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.DeployOptions, string, int) error); ok {
		r0 = rf(ctx, opts, nodename, count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WatchDeployStatus provides a mock function with given fields: ctx, appname, entrypoint, nodename
func (_m *Store) WatchDeployStatus(ctx context.Context, appname string, entrypoint string, nodename string) chan *types.DeployStatus {
	ret := _m.Called(ctx, appname, entrypoint, nodename)

	var r0 chan *types.DeployStatus
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) chan *types.DeployStatus); ok {
		r0 = rf(ctx, appname, entrypoint, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.DeployStatus)
		}
	}

	return r0
}
