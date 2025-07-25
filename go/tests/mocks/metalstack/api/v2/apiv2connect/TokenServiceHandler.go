// Code generated by mockery v2.53.4. DO NOT EDIT.

package apiv2connect

import (
	connect "connectrpc.com/connect"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TokenServiceHandler is an autogenerated mock type for the TokenServiceHandler type
type TokenServiceHandler struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *TokenServiceHandler) Create(_a0 context.Context, _a1 *connect.Request[apiv2.TokenServiceCreateRequest]) (*connect.Response[apiv2.TokenServiceCreateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *connect.Response[apiv2.TokenServiceCreateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceCreateRequest]) (*connect.Response[apiv2.TokenServiceCreateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceCreateRequest]) *connect.Response[apiv2.TokenServiceCreateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.TokenServiceCreateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.TokenServiceCreateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *TokenServiceHandler) Get(_a0 context.Context, _a1 *connect.Request[apiv2.TokenServiceGetRequest]) (*connect.Response[apiv2.TokenServiceGetResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *connect.Response[apiv2.TokenServiceGetResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceGetRequest]) (*connect.Response[apiv2.TokenServiceGetResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceGetRequest]) *connect.Response[apiv2.TokenServiceGetResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.TokenServiceGetResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.TokenServiceGetRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *TokenServiceHandler) List(_a0 context.Context, _a1 *connect.Request[apiv2.TokenServiceListRequest]) (*connect.Response[apiv2.TokenServiceListResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *connect.Response[apiv2.TokenServiceListResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceListRequest]) (*connect.Response[apiv2.TokenServiceListResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceListRequest]) *connect.Response[apiv2.TokenServiceListResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.TokenServiceListResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.TokenServiceListRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Refresh provides a mock function with given fields: _a0, _a1
func (_m *TokenServiceHandler) Refresh(_a0 context.Context, _a1 *connect.Request[apiv2.TokenServiceRefreshRequest]) (*connect.Response[apiv2.TokenServiceRefreshResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Refresh")
	}

	var r0 *connect.Response[apiv2.TokenServiceRefreshResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceRefreshRequest]) (*connect.Response[apiv2.TokenServiceRefreshResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceRefreshRequest]) *connect.Response[apiv2.TokenServiceRefreshResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.TokenServiceRefreshResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.TokenServiceRefreshRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Revoke provides a mock function with given fields: _a0, _a1
func (_m *TokenServiceHandler) Revoke(_a0 context.Context, _a1 *connect.Request[apiv2.TokenServiceRevokeRequest]) (*connect.Response[apiv2.TokenServiceRevokeResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Revoke")
	}

	var r0 *connect.Response[apiv2.TokenServiceRevokeResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceRevokeRequest]) (*connect.Response[apiv2.TokenServiceRevokeResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceRevokeRequest]) *connect.Response[apiv2.TokenServiceRevokeResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.TokenServiceRevokeResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.TokenServiceRevokeRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *TokenServiceHandler) Update(_a0 context.Context, _a1 *connect.Request[apiv2.TokenServiceUpdateRequest]) (*connect.Response[apiv2.TokenServiceUpdateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *connect.Response[apiv2.TokenServiceUpdateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceUpdateRequest]) (*connect.Response[apiv2.TokenServiceUpdateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.TokenServiceUpdateRequest]) *connect.Response[apiv2.TokenServiceUpdateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.TokenServiceUpdateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.TokenServiceUpdateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenServiceHandler creates a new instance of TokenServiceHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenServiceHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenServiceHandler {
	mock := &TokenServiceHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
