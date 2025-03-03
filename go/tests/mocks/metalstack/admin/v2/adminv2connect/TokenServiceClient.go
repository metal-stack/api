// Code generated by mockery v2.53.0. DO NOT EDIT.

package adminv2connect

import (
	connect "connectrpc.com/connect"
	adminv2 "github.com/metal-stack/api/go/metalstack/admin/v2"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TokenServiceClient is an autogenerated mock type for the TokenServiceClient type
type TokenServiceClient struct {
	mock.Mock
}

// List provides a mock function with given fields: _a0, _a1
func (_m *TokenServiceClient) List(_a0 context.Context, _a1 *connect.Request[adminv2.TokenServiceListRequest]) (*connect.Response[adminv2.TokenServiceListResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *connect.Response[adminv2.TokenServiceListResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.TokenServiceListRequest]) (*connect.Response[adminv2.TokenServiceListResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.TokenServiceListRequest]) *connect.Response[adminv2.TokenServiceListResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[adminv2.TokenServiceListResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[adminv2.TokenServiceListRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Revoke provides a mock function with given fields: _a0, _a1
func (_m *TokenServiceClient) Revoke(_a0 context.Context, _a1 *connect.Request[adminv2.TokenServiceRevokeRequest]) (*connect.Response[adminv2.TokenServiceRevokeResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Revoke")
	}

	var r0 *connect.Response[adminv2.TokenServiceRevokeResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.TokenServiceRevokeRequest]) (*connect.Response[adminv2.TokenServiceRevokeResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[adminv2.TokenServiceRevokeRequest]) *connect.Response[adminv2.TokenServiceRevokeResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[adminv2.TokenServiceRevokeResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[adminv2.TokenServiceRevokeRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenServiceClient creates a new instance of TokenServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenServiceClient {
	mock := &TokenServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
