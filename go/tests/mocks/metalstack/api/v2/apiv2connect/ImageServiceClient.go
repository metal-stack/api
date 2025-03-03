// Code generated by mockery v2.53.0. DO NOT EDIT.

package apiv2connect

import (
	connect "connectrpc.com/connect"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ImageServiceClient is an autogenerated mock type for the ImageServiceClient type
type ImageServiceClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *ImageServiceClient) Get(_a0 context.Context, _a1 *connect.Request[apiv2.ImageServiceGetRequest]) (*connect.Response[apiv2.ImageServiceGetResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *connect.Response[apiv2.ImageServiceGetResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.ImageServiceGetRequest]) (*connect.Response[apiv2.ImageServiceGetResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.ImageServiceGetRequest]) *connect.Response[apiv2.ImageServiceGetResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.ImageServiceGetResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.ImageServiceGetRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *ImageServiceClient) List(_a0 context.Context, _a1 *connect.Request[apiv2.ImageServiceListRequest]) (*connect.Response[apiv2.ImageServiceListResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *connect.Response[apiv2.ImageServiceListResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.ImageServiceListRequest]) (*connect.Response[apiv2.ImageServiceListResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.ImageServiceListRequest]) *connect.Response[apiv2.ImageServiceListResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.ImageServiceListResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.ImageServiceListRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewImageServiceClient creates a new instance of ImageServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewImageServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ImageServiceClient {
	mock := &ImageServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
