// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20250219001/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// AWSClustersDNSApi is an autogenerated mock type for the AWSClustersDNSApi type
type AWSClustersDNSApi struct {
	mock.Mock
}

type AWSClustersDNSApi_Expecter struct {
	mock *mock.Mock
}

func (_m *AWSClustersDNSApi) EXPECT() *AWSClustersDNSApi_Expecter {
	return &AWSClustersDNSApi_Expecter{mock: &_m.Mock}
}

// GetAwsCustomDns provides a mock function with given fields: ctx, groupId
func (_m *AWSClustersDNSApi) GetAwsCustomDns(ctx context.Context, groupId string) admin.GetAwsCustomDnsApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for GetAwsCustomDns")
	}

	var r0 admin.GetAwsCustomDnsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.GetAwsCustomDnsApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.GetAwsCustomDnsApiRequest)
	}

	return r0
}

// AWSClustersDNSApi_GetAwsCustomDns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAwsCustomDns'
type AWSClustersDNSApi_GetAwsCustomDns_Call struct {
	*mock.Call
}

// GetAwsCustomDns is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *AWSClustersDNSApi_Expecter) GetAwsCustomDns(ctx any, groupId any) *AWSClustersDNSApi_GetAwsCustomDns_Call {
	return &AWSClustersDNSApi_GetAwsCustomDns_Call{Call: _e.mock.On("GetAwsCustomDns", ctx, groupId)}
}

func (_c *AWSClustersDNSApi_GetAwsCustomDns_Call) Run(run func(ctx context.Context, groupId string)) *AWSClustersDNSApi_GetAwsCustomDns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AWSClustersDNSApi_GetAwsCustomDns_Call) Return(_a0 admin.GetAwsCustomDnsApiRequest) *AWSClustersDNSApi_GetAwsCustomDns_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AWSClustersDNSApi_GetAwsCustomDns_Call) RunAndReturn(run func(context.Context, string) admin.GetAwsCustomDnsApiRequest) *AWSClustersDNSApi_GetAwsCustomDns_Call {
	_c.Call.Return(run)
	return _c
}

// GetAwsCustomDnsExecute provides a mock function with given fields: r
func (_m *AWSClustersDNSApi) GetAwsCustomDnsExecute(r admin.GetAwsCustomDnsApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetAwsCustomDnsExecute")
	}

	var r0 *admin.AWSCustomDNSEnabled
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetAwsCustomDnsApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetAwsCustomDnsApiRequest) *admin.AWSCustomDNSEnabled); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.AWSCustomDNSEnabled)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetAwsCustomDnsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetAwsCustomDnsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AWSClustersDNSApi_GetAwsCustomDnsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAwsCustomDnsExecute'
type AWSClustersDNSApi_GetAwsCustomDnsExecute_Call struct {
	*mock.Call
}

// GetAwsCustomDnsExecute is a helper method to define mock.On call
//   - r admin.GetAwsCustomDnsApiRequest
func (_e *AWSClustersDNSApi_Expecter) GetAwsCustomDnsExecute(r any) *AWSClustersDNSApi_GetAwsCustomDnsExecute_Call {
	return &AWSClustersDNSApi_GetAwsCustomDnsExecute_Call{Call: _e.mock.On("GetAwsCustomDnsExecute", r)}
}

func (_c *AWSClustersDNSApi_GetAwsCustomDnsExecute_Call) Run(run func(r admin.GetAwsCustomDnsApiRequest)) *AWSClustersDNSApi_GetAwsCustomDnsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetAwsCustomDnsApiRequest))
	})
	return _c
}

func (_c *AWSClustersDNSApi_GetAwsCustomDnsExecute_Call) Return(_a0 *admin.AWSCustomDNSEnabled, _a1 *http.Response, _a2 error) *AWSClustersDNSApi_GetAwsCustomDnsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *AWSClustersDNSApi_GetAwsCustomDnsExecute_Call) RunAndReturn(run func(admin.GetAwsCustomDnsApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error)) *AWSClustersDNSApi_GetAwsCustomDnsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetAwsCustomDnsWithParams provides a mock function with given fields: ctx, args
func (_m *AWSClustersDNSApi) GetAwsCustomDnsWithParams(ctx context.Context, args *admin.GetAwsCustomDnsApiParams) admin.GetAwsCustomDnsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetAwsCustomDnsWithParams")
	}

	var r0 admin.GetAwsCustomDnsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetAwsCustomDnsApiParams) admin.GetAwsCustomDnsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetAwsCustomDnsApiRequest)
	}

	return r0
}

// AWSClustersDNSApi_GetAwsCustomDnsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAwsCustomDnsWithParams'
type AWSClustersDNSApi_GetAwsCustomDnsWithParams_Call struct {
	*mock.Call
}

// GetAwsCustomDnsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetAwsCustomDnsApiParams
func (_e *AWSClustersDNSApi_Expecter) GetAwsCustomDnsWithParams(ctx any, args any) *AWSClustersDNSApi_GetAwsCustomDnsWithParams_Call {
	return &AWSClustersDNSApi_GetAwsCustomDnsWithParams_Call{Call: _e.mock.On("GetAwsCustomDnsWithParams", ctx, args)}
}

func (_c *AWSClustersDNSApi_GetAwsCustomDnsWithParams_Call) Run(run func(ctx context.Context, args *admin.GetAwsCustomDnsApiParams)) *AWSClustersDNSApi_GetAwsCustomDnsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetAwsCustomDnsApiParams))
	})
	return _c
}

func (_c *AWSClustersDNSApi_GetAwsCustomDnsWithParams_Call) Return(_a0 admin.GetAwsCustomDnsApiRequest) *AWSClustersDNSApi_GetAwsCustomDnsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AWSClustersDNSApi_GetAwsCustomDnsWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetAwsCustomDnsApiParams) admin.GetAwsCustomDnsApiRequest) *AWSClustersDNSApi_GetAwsCustomDnsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleAwsCustomDns provides a mock function with given fields: ctx, groupId, aWSCustomDNSEnabled
func (_m *AWSClustersDNSApi) ToggleAwsCustomDns(ctx context.Context, groupId string, aWSCustomDNSEnabled *admin.AWSCustomDNSEnabled) admin.ToggleAwsCustomDnsApiRequest {
	ret := _m.Called(ctx, groupId, aWSCustomDNSEnabled)

	if len(ret) == 0 {
		panic("no return value specified for ToggleAwsCustomDns")
	}

	var r0 admin.ToggleAwsCustomDnsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.AWSCustomDNSEnabled) admin.ToggleAwsCustomDnsApiRequest); ok {
		r0 = rf(ctx, groupId, aWSCustomDNSEnabled)
	} else {
		r0 = ret.Get(0).(admin.ToggleAwsCustomDnsApiRequest)
	}

	return r0
}

// AWSClustersDNSApi_ToggleAwsCustomDns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleAwsCustomDns'
type AWSClustersDNSApi_ToggleAwsCustomDns_Call struct {
	*mock.Call
}

// ToggleAwsCustomDns is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - aWSCustomDNSEnabled *admin.AWSCustomDNSEnabled
func (_e *AWSClustersDNSApi_Expecter) ToggleAwsCustomDns(ctx any, groupId any, aWSCustomDNSEnabled any) *AWSClustersDNSApi_ToggleAwsCustomDns_Call {
	return &AWSClustersDNSApi_ToggleAwsCustomDns_Call{Call: _e.mock.On("ToggleAwsCustomDns", ctx, groupId, aWSCustomDNSEnabled)}
}

func (_c *AWSClustersDNSApi_ToggleAwsCustomDns_Call) Run(run func(ctx context.Context, groupId string, aWSCustomDNSEnabled *admin.AWSCustomDNSEnabled)) *AWSClustersDNSApi_ToggleAwsCustomDns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.AWSCustomDNSEnabled))
	})
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAwsCustomDns_Call) Return(_a0 admin.ToggleAwsCustomDnsApiRequest) *AWSClustersDNSApi_ToggleAwsCustomDns_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAwsCustomDns_Call) RunAndReturn(run func(context.Context, string, *admin.AWSCustomDNSEnabled) admin.ToggleAwsCustomDnsApiRequest) *AWSClustersDNSApi_ToggleAwsCustomDns_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleAwsCustomDnsExecute provides a mock function with given fields: r
func (_m *AWSClustersDNSApi) ToggleAwsCustomDnsExecute(r admin.ToggleAwsCustomDnsApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ToggleAwsCustomDnsExecute")
	}

	var r0 *admin.AWSCustomDNSEnabled
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ToggleAwsCustomDnsApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ToggleAwsCustomDnsApiRequest) *admin.AWSCustomDNSEnabled); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.AWSCustomDNSEnabled)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ToggleAwsCustomDnsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ToggleAwsCustomDnsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AWSClustersDNSApi_ToggleAwsCustomDnsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleAwsCustomDnsExecute'
type AWSClustersDNSApi_ToggleAwsCustomDnsExecute_Call struct {
	*mock.Call
}

// ToggleAwsCustomDnsExecute is a helper method to define mock.On call
//   - r admin.ToggleAwsCustomDnsApiRequest
func (_e *AWSClustersDNSApi_Expecter) ToggleAwsCustomDnsExecute(r any) *AWSClustersDNSApi_ToggleAwsCustomDnsExecute_Call {
	return &AWSClustersDNSApi_ToggleAwsCustomDnsExecute_Call{Call: _e.mock.On("ToggleAwsCustomDnsExecute", r)}
}

func (_c *AWSClustersDNSApi_ToggleAwsCustomDnsExecute_Call) Run(run func(r admin.ToggleAwsCustomDnsApiRequest)) *AWSClustersDNSApi_ToggleAwsCustomDnsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ToggleAwsCustomDnsApiRequest))
	})
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAwsCustomDnsExecute_Call) Return(_a0 *admin.AWSCustomDNSEnabled, _a1 *http.Response, _a2 error) *AWSClustersDNSApi_ToggleAwsCustomDnsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAwsCustomDnsExecute_Call) RunAndReturn(run func(admin.ToggleAwsCustomDnsApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error)) *AWSClustersDNSApi_ToggleAwsCustomDnsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleAwsCustomDnsWithParams provides a mock function with given fields: ctx, args
func (_m *AWSClustersDNSApi) ToggleAwsCustomDnsWithParams(ctx context.Context, args *admin.ToggleAwsCustomDnsApiParams) admin.ToggleAwsCustomDnsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ToggleAwsCustomDnsWithParams")
	}

	var r0 admin.ToggleAwsCustomDnsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ToggleAwsCustomDnsApiParams) admin.ToggleAwsCustomDnsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ToggleAwsCustomDnsApiRequest)
	}

	return r0
}

// AWSClustersDNSApi_ToggleAwsCustomDnsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleAwsCustomDnsWithParams'
type AWSClustersDNSApi_ToggleAwsCustomDnsWithParams_Call struct {
	*mock.Call
}

// ToggleAwsCustomDnsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ToggleAwsCustomDnsApiParams
func (_e *AWSClustersDNSApi_Expecter) ToggleAwsCustomDnsWithParams(ctx any, args any) *AWSClustersDNSApi_ToggleAwsCustomDnsWithParams_Call {
	return &AWSClustersDNSApi_ToggleAwsCustomDnsWithParams_Call{Call: _e.mock.On("ToggleAwsCustomDnsWithParams", ctx, args)}
}

func (_c *AWSClustersDNSApi_ToggleAwsCustomDnsWithParams_Call) Run(run func(ctx context.Context, args *admin.ToggleAwsCustomDnsApiParams)) *AWSClustersDNSApi_ToggleAwsCustomDnsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ToggleAwsCustomDnsApiParams))
	})
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAwsCustomDnsWithParams_Call) Return(_a0 admin.ToggleAwsCustomDnsApiRequest) *AWSClustersDNSApi_ToggleAwsCustomDnsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAwsCustomDnsWithParams_Call) RunAndReturn(run func(context.Context, *admin.ToggleAwsCustomDnsApiParams) admin.ToggleAwsCustomDnsApiRequest) *AWSClustersDNSApi_ToggleAwsCustomDnsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewAWSClustersDNSApi creates a new instance of AWSClustersDNSApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAWSClustersDNSApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *AWSClustersDNSApi {
	mock := &AWSClustersDNSApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
