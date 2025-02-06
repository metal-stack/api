// Code generated by mockery v2.52.1. DO NOT EDIT.

package adminv2connect

import (
	connect "connectrpc.com/connect"
	adminv2 "github.com/metal-stack/api/go/metalstack/admin/v2"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// PartitionServiceClient is an autogenerated mock type for the PartitionServiceClient type
type PartitionServiceClient struct {
	mock.Mock
}

// Capacity provides a mock function with given fields: _a0, _a1
func (_m *PartitionServiceClient) Capacity(_a0 context.Context, _a1 *connect.Request[adminv2.PartitionServiceCapacityRequest]) (*connect.Response[adminv2.PartitionServiceCapacityResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Capacity")
	}

	var r0 *connect.Response[adminv2.PartitionServiceCapacityResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.PartitionServiceCapacityRequest]) (*connect.Response[adminv2.PartitionServiceCapacityResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.PartitionServiceCapacityRequest]) *connect.Response[adminv2.PartitionServiceCapacityResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[adminv2.PartitionServiceCapacityResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[adminv2.PartitionServiceCapacityRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *PartitionServiceClient) Create(_a0 context.Context, _a1 *connect.Request[adminv2.PartitionServiceCreateRequest]) (*connect.Response[adminv2.PartitionServiceCreateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *connect.Response[adminv2.PartitionServiceCreateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.PartitionServiceCreateRequest]) (*connect.Response[adminv2.PartitionServiceCreateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.PartitionServiceCreateRequest]) *connect.Response[adminv2.PartitionServiceCreateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[adminv2.PartitionServiceCreateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[adminv2.PartitionServiceCreateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *PartitionServiceClient) Delete(_a0 context.Context, _a1 *connect.Request[adminv2.PartitionServiceDeleteRequest]) (*connect.Response[adminv2.PartitionServiceDeleteResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *connect.Response[adminv2.PartitionServiceDeleteResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.PartitionServiceDeleteRequest]) (*connect.Response[adminv2.PartitionServiceDeleteResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.PartitionServiceDeleteRequest]) *connect.Response[adminv2.PartitionServiceDeleteResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[adminv2.PartitionServiceDeleteResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[adminv2.PartitionServiceDeleteRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *PartitionServiceClient) Update(_a0 context.Context, _a1 *connect.Request[adminv2.PartitionServiceUpdateRequest]) (*connect.Response[adminv2.PartitionServiceUpdateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *connect.Response[adminv2.PartitionServiceUpdateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.PartitionServiceUpdateRequest]) (*connect.Response[adminv2.PartitionServiceUpdateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.PartitionServiceUpdateRequest]) *connect.Response[adminv2.PartitionServiceUpdateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[adminv2.PartitionServiceUpdateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[adminv2.PartitionServiceUpdateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPartitionServiceClient creates a new instance of PartitionServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPartitionServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *PartitionServiceClient {
	mock := &PartitionServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
