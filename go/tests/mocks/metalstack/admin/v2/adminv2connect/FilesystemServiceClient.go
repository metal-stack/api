// Code generated by mockery v2.53.0. DO NOT EDIT.

package adminv2connect

import (
	connect "connectrpc.com/connect"
	adminv2 "github.com/metal-stack/api/go/metalstack/admin/v2"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// FilesystemServiceClient is an autogenerated mock type for the FilesystemServiceClient type
type FilesystemServiceClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *FilesystemServiceClient) Create(_a0 context.Context, _a1 *connect.Request[adminv2.FilesystemServiceCreateRequest]) (*connect.Response[adminv2.FilesystemServiceCreateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *connect.Response[adminv2.FilesystemServiceCreateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.FilesystemServiceCreateRequest]) (*connect.Response[adminv2.FilesystemServiceCreateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.FilesystemServiceCreateRequest]) *connect.Response[adminv2.FilesystemServiceCreateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[adminv2.FilesystemServiceCreateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[adminv2.FilesystemServiceCreateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *FilesystemServiceClient) Delete(_a0 context.Context, _a1 *connect.Request[adminv2.FilesystemServiceDeleteRequest]) (*connect.Response[adminv2.FilesystemServiceDeleteResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *connect.Response[adminv2.FilesystemServiceDeleteResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.FilesystemServiceDeleteRequest]) (*connect.Response[adminv2.FilesystemServiceDeleteResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.FilesystemServiceDeleteRequest]) *connect.Response[adminv2.FilesystemServiceDeleteResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[adminv2.FilesystemServiceDeleteResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[adminv2.FilesystemServiceDeleteRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *FilesystemServiceClient) Update(_a0 context.Context, _a1 *connect.Request[adminv2.FilesystemServiceUpdateRequest]) (*connect.Response[adminv2.FilesystemServiceUpdateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *connect.Response[adminv2.FilesystemServiceUpdateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.FilesystemServiceUpdateRequest]) (*connect.Response[adminv2.FilesystemServiceUpdateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.FilesystemServiceUpdateRequest]) *connect.Response[adminv2.FilesystemServiceUpdateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[adminv2.FilesystemServiceUpdateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[adminv2.FilesystemServiceUpdateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFilesystemServiceClient creates a new instance of FilesystemServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFilesystemServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *FilesystemServiceClient {
	mock := &FilesystemServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
