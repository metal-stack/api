// Code generated by mockery v2.53.4. DO NOT EDIT.

package adminv2connect

import (
	connect "connectrpc.com/connect"
	adminv2 "github.com/metal-stack/api/go/metalstack/admin/v2"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TenantServiceHandler is an autogenerated mock type for the TenantServiceHandler type
type TenantServiceHandler struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *TenantServiceHandler) Create(_a0 context.Context, _a1 *connect.Request[adminv2.TenantServiceCreateRequest]) (*connect.Response[adminv2.TenantServiceCreateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *connect.Response[adminv2.TenantServiceCreateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.TenantServiceCreateRequest]) (*connect.Response[adminv2.TenantServiceCreateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.TenantServiceCreateRequest]) *connect.Response[adminv2.TenantServiceCreateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[adminv2.TenantServiceCreateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[adminv2.TenantServiceCreateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *TenantServiceHandler) List(_a0 context.Context, _a1 *connect.Request[adminv2.TenantServiceListRequest]) (*connect.Response[adminv2.TenantServiceListResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *connect.Response[adminv2.TenantServiceListResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.TenantServiceListRequest]) (*connect.Response[adminv2.TenantServiceListResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.TenantServiceListRequest]) *connect.Response[adminv2.TenantServiceListResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[adminv2.TenantServiceListResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[adminv2.TenantServiceListRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTenantServiceHandler creates a new instance of TenantServiceHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTenantServiceHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *TenantServiceHandler {
	mock := &TenantServiceHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
