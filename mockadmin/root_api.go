// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20240805002/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RootApi is an autogenerated mock type for the RootApi type
type RootApi struct {
	mock.Mock
}

type RootApi_Expecter struct {
	mock *mock.Mock
}

func (_m *RootApi) EXPECT() *RootApi_Expecter {
	return &RootApi_Expecter{mock: &_m.Mock}
}

// GetSystemStatus provides a mock function with given fields: ctx
func (_m *RootApi) GetSystemStatus(ctx context.Context) admin.GetSystemStatusApiRequest {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetSystemStatus")
	}

	var r0 admin.GetSystemStatusApiRequest
	if rf, ok := ret.Get(0).(func(context.Context) admin.GetSystemStatusApiRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(admin.GetSystemStatusApiRequest)
	}

	return r0
}

// RootApi_GetSystemStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSystemStatus'
type RootApi_GetSystemStatus_Call struct {
	*mock.Call
}

// GetSystemStatus is a helper method to define mock.On call
//   - ctx context.Context
func (_e *RootApi_Expecter) GetSystemStatus(ctx interface{}) *RootApi_GetSystemStatus_Call {
	return &RootApi_GetSystemStatus_Call{Call: _e.mock.On("GetSystemStatus", ctx)}
}

func (_c *RootApi_GetSystemStatus_Call) Run(run func(ctx context.Context)) *RootApi_GetSystemStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RootApi_GetSystemStatus_Call) Return(_a0 admin.GetSystemStatusApiRequest) *RootApi_GetSystemStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RootApi_GetSystemStatus_Call) RunAndReturn(run func(context.Context) admin.GetSystemStatusApiRequest) *RootApi_GetSystemStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetSystemStatusExecute provides a mock function with given fields: r
func (_m *RootApi) GetSystemStatusExecute(r admin.GetSystemStatusApiRequest) (*admin.SystemStatus, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetSystemStatusExecute")
	}

	var r0 *admin.SystemStatus
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetSystemStatusApiRequest) (*admin.SystemStatus, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetSystemStatusApiRequest) *admin.SystemStatus); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.SystemStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetSystemStatusApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetSystemStatusApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RootApi_GetSystemStatusExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSystemStatusExecute'
type RootApi_GetSystemStatusExecute_Call struct {
	*mock.Call
}

// GetSystemStatusExecute is a helper method to define mock.On call
//   - r admin.GetSystemStatusApiRequest
func (_e *RootApi_Expecter) GetSystemStatusExecute(r interface{}) *RootApi_GetSystemStatusExecute_Call {
	return &RootApi_GetSystemStatusExecute_Call{Call: _e.mock.On("GetSystemStatusExecute", r)}
}

func (_c *RootApi_GetSystemStatusExecute_Call) Run(run func(r admin.GetSystemStatusApiRequest)) *RootApi_GetSystemStatusExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetSystemStatusApiRequest))
	})
	return _c
}

func (_c *RootApi_GetSystemStatusExecute_Call) Return(_a0 *admin.SystemStatus, _a1 *http.Response, _a2 error) *RootApi_GetSystemStatusExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RootApi_GetSystemStatusExecute_Call) RunAndReturn(run func(admin.GetSystemStatusApiRequest) (*admin.SystemStatus, *http.Response, error)) *RootApi_GetSystemStatusExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetSystemStatusWithParams provides a mock function with given fields: ctx, args
func (_m *RootApi) GetSystemStatusWithParams(ctx context.Context, args *admin.GetSystemStatusApiParams) admin.GetSystemStatusApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetSystemStatusWithParams")
	}

	var r0 admin.GetSystemStatusApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetSystemStatusApiParams) admin.GetSystemStatusApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetSystemStatusApiRequest)
	}

	return r0
}

// RootApi_GetSystemStatusWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSystemStatusWithParams'
type RootApi_GetSystemStatusWithParams_Call struct {
	*mock.Call
}

// GetSystemStatusWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetSystemStatusApiParams
func (_e *RootApi_Expecter) GetSystemStatusWithParams(ctx interface{}, args interface{}) *RootApi_GetSystemStatusWithParams_Call {
	return &RootApi_GetSystemStatusWithParams_Call{Call: _e.mock.On("GetSystemStatusWithParams", ctx, args)}
}

func (_c *RootApi_GetSystemStatusWithParams_Call) Run(run func(ctx context.Context, args *admin.GetSystemStatusApiParams)) *RootApi_GetSystemStatusWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetSystemStatusApiParams))
	})
	return _c
}

func (_c *RootApi_GetSystemStatusWithParams_Call) Return(_a0 admin.GetSystemStatusApiRequest) *RootApi_GetSystemStatusWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RootApi_GetSystemStatusWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetSystemStatusApiParams) admin.GetSystemStatusApiRequest) *RootApi_GetSystemStatusWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnAllControlPlaneIPAddresses provides a mock function with given fields: ctx
func (_m *RootApi) ReturnAllControlPlaneIPAddresses(ctx context.Context) admin.ReturnAllControlPlaneIPAddressesApiRequest {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ReturnAllControlPlaneIPAddresses")
	}

	var r0 admin.ReturnAllControlPlaneIPAddressesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context) admin.ReturnAllControlPlaneIPAddressesApiRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(admin.ReturnAllControlPlaneIPAddressesApiRequest)
	}

	return r0
}

// RootApi_ReturnAllControlPlaneIPAddresses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnAllControlPlaneIPAddresses'
type RootApi_ReturnAllControlPlaneIPAddresses_Call struct {
	*mock.Call
}

// ReturnAllControlPlaneIPAddresses is a helper method to define mock.On call
//   - ctx context.Context
func (_e *RootApi_Expecter) ReturnAllControlPlaneIPAddresses(ctx interface{}) *RootApi_ReturnAllControlPlaneIPAddresses_Call {
	return &RootApi_ReturnAllControlPlaneIPAddresses_Call{Call: _e.mock.On("ReturnAllControlPlaneIPAddresses", ctx)}
}

func (_c *RootApi_ReturnAllControlPlaneIPAddresses_Call) Run(run func(ctx context.Context)) *RootApi_ReturnAllControlPlaneIPAddresses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIPAddresses_Call) Return(_a0 admin.ReturnAllControlPlaneIPAddressesApiRequest) *RootApi_ReturnAllControlPlaneIPAddresses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIPAddresses_Call) RunAndReturn(run func(context.Context) admin.ReturnAllControlPlaneIPAddressesApiRequest) *RootApi_ReturnAllControlPlaneIPAddresses_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnAllControlPlaneIPAddressesExecute provides a mock function with given fields: r
func (_m *RootApi) ReturnAllControlPlaneIPAddressesExecute(r admin.ReturnAllControlPlaneIPAddressesApiRequest) (*admin.ControlPlaneIPAddresses, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ReturnAllControlPlaneIPAddressesExecute")
	}

	var r0 *admin.ControlPlaneIPAddresses
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ReturnAllControlPlaneIPAddressesApiRequest) (*admin.ControlPlaneIPAddresses, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ReturnAllControlPlaneIPAddressesApiRequest) *admin.ControlPlaneIPAddresses); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ControlPlaneIPAddresses)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ReturnAllControlPlaneIPAddressesApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ReturnAllControlPlaneIPAddressesApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RootApi_ReturnAllControlPlaneIPAddressesExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnAllControlPlaneIPAddressesExecute'
type RootApi_ReturnAllControlPlaneIPAddressesExecute_Call struct {
	*mock.Call
}

// ReturnAllControlPlaneIPAddressesExecute is a helper method to define mock.On call
//   - r admin.ReturnAllControlPlaneIPAddressesApiRequest
func (_e *RootApi_Expecter) ReturnAllControlPlaneIPAddressesExecute(r interface{}) *RootApi_ReturnAllControlPlaneIPAddressesExecute_Call {
	return &RootApi_ReturnAllControlPlaneIPAddressesExecute_Call{Call: _e.mock.On("ReturnAllControlPlaneIPAddressesExecute", r)}
}

func (_c *RootApi_ReturnAllControlPlaneIPAddressesExecute_Call) Run(run func(r admin.ReturnAllControlPlaneIPAddressesApiRequest)) *RootApi_ReturnAllControlPlaneIPAddressesExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ReturnAllControlPlaneIPAddressesApiRequest))
	})
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIPAddressesExecute_Call) Return(_a0 *admin.ControlPlaneIPAddresses, _a1 *http.Response, _a2 error) *RootApi_ReturnAllControlPlaneIPAddressesExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIPAddressesExecute_Call) RunAndReturn(run func(admin.ReturnAllControlPlaneIPAddressesApiRequest) (*admin.ControlPlaneIPAddresses, *http.Response, error)) *RootApi_ReturnAllControlPlaneIPAddressesExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnAllControlPlaneIPAddressesWithParams provides a mock function with given fields: ctx, args
func (_m *RootApi) ReturnAllControlPlaneIPAddressesWithParams(ctx context.Context, args *admin.ReturnAllControlPlaneIPAddressesApiParams) admin.ReturnAllControlPlaneIPAddressesApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ReturnAllControlPlaneIPAddressesWithParams")
	}

	var r0 admin.ReturnAllControlPlaneIPAddressesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ReturnAllControlPlaneIPAddressesApiParams) admin.ReturnAllControlPlaneIPAddressesApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ReturnAllControlPlaneIPAddressesApiRequest)
	}

	return r0
}

// RootApi_ReturnAllControlPlaneIPAddressesWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnAllControlPlaneIPAddressesWithParams'
type RootApi_ReturnAllControlPlaneIPAddressesWithParams_Call struct {
	*mock.Call
}

// ReturnAllControlPlaneIPAddressesWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ReturnAllControlPlaneIPAddressesApiParams
func (_e *RootApi_Expecter) ReturnAllControlPlaneIPAddressesWithParams(ctx interface{}, args interface{}) *RootApi_ReturnAllControlPlaneIPAddressesWithParams_Call {
	return &RootApi_ReturnAllControlPlaneIPAddressesWithParams_Call{Call: _e.mock.On("ReturnAllControlPlaneIPAddressesWithParams", ctx, args)}
}

func (_c *RootApi_ReturnAllControlPlaneIPAddressesWithParams_Call) Run(run func(ctx context.Context, args *admin.ReturnAllControlPlaneIPAddressesApiParams)) *RootApi_ReturnAllControlPlaneIPAddressesWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ReturnAllControlPlaneIPAddressesApiParams))
	})
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIPAddressesWithParams_Call) Return(_a0 admin.ReturnAllControlPlaneIPAddressesApiRequest) *RootApi_ReturnAllControlPlaneIPAddressesWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIPAddressesWithParams_Call) RunAndReturn(run func(context.Context, *admin.ReturnAllControlPlaneIPAddressesApiParams) admin.ReturnAllControlPlaneIPAddressesApiRequest) *RootApi_ReturnAllControlPlaneIPAddressesWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewRootApi creates a new instance of RootApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRootApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *RootApi {
	mock := &RootApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
