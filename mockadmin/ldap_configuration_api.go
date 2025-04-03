// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20250312002/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// LDAPConfigurationApi is an autogenerated mock type for the LDAPConfigurationApi type
type LDAPConfigurationApi struct {
	mock.Mock
}

type LDAPConfigurationApi_Expecter struct {
	mock *mock.Mock
}

func (_m *LDAPConfigurationApi) EXPECT() *LDAPConfigurationApi_Expecter {
	return &LDAPConfigurationApi_Expecter{mock: &_m.Mock}
}

// DeleteLdapConfiguration provides a mock function with given fields: ctx, groupId
func (_m *LDAPConfigurationApi) DeleteLdapConfiguration(ctx context.Context, groupId string) admin.DeleteLdapConfigurationApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteLdapConfiguration")
	}

	var r0 admin.DeleteLdapConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.DeleteLdapConfigurationApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.DeleteLdapConfigurationApiRequest)
	}

	return r0
}

// LDAPConfigurationApi_DeleteLdapConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteLdapConfiguration'
type LDAPConfigurationApi_DeleteLdapConfiguration_Call struct {
	*mock.Call
}

// DeleteLdapConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *LDAPConfigurationApi_Expecter) DeleteLdapConfiguration(ctx any, groupId any) *LDAPConfigurationApi_DeleteLdapConfiguration_Call {
	return &LDAPConfigurationApi_DeleteLdapConfiguration_Call{Call: _e.mock.On("DeleteLdapConfiguration", ctx, groupId)}
}

func (_c *LDAPConfigurationApi_DeleteLdapConfiguration_Call) Run(run func(ctx context.Context, groupId string)) *LDAPConfigurationApi_DeleteLdapConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *LDAPConfigurationApi_DeleteLdapConfiguration_Call) Return(_a0 admin.DeleteLdapConfigurationApiRequest) *LDAPConfigurationApi_DeleteLdapConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LDAPConfigurationApi_DeleteLdapConfiguration_Call) RunAndReturn(run func(context.Context, string) admin.DeleteLdapConfigurationApiRequest) *LDAPConfigurationApi_DeleteLdapConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteLdapConfigurationExecute provides a mock function with given fields: r
func (_m *LDAPConfigurationApi) DeleteLdapConfigurationExecute(r admin.DeleteLdapConfigurationApiRequest) (*admin.UserSecurity, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DeleteLdapConfigurationExecute")
	}

	var r0 *admin.UserSecurity
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.DeleteLdapConfigurationApiRequest) (*admin.UserSecurity, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DeleteLdapConfigurationApiRequest) *admin.UserSecurity); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.UserSecurity)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DeleteLdapConfigurationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.DeleteLdapConfigurationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LDAPConfigurationApi_DeleteLdapConfigurationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteLdapConfigurationExecute'
type LDAPConfigurationApi_DeleteLdapConfigurationExecute_Call struct {
	*mock.Call
}

// DeleteLdapConfigurationExecute is a helper method to define mock.On call
//   - r admin.DeleteLdapConfigurationApiRequest
func (_e *LDAPConfigurationApi_Expecter) DeleteLdapConfigurationExecute(r any) *LDAPConfigurationApi_DeleteLdapConfigurationExecute_Call {
	return &LDAPConfigurationApi_DeleteLdapConfigurationExecute_Call{Call: _e.mock.On("DeleteLdapConfigurationExecute", r)}
}

func (_c *LDAPConfigurationApi_DeleteLdapConfigurationExecute_Call) Run(run func(r admin.DeleteLdapConfigurationApiRequest)) *LDAPConfigurationApi_DeleteLdapConfigurationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DeleteLdapConfigurationApiRequest))
	})
	return _c
}

func (_c *LDAPConfigurationApi_DeleteLdapConfigurationExecute_Call) Return(_a0 *admin.UserSecurity, _a1 *http.Response, _a2 error) *LDAPConfigurationApi_DeleteLdapConfigurationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *LDAPConfigurationApi_DeleteLdapConfigurationExecute_Call) RunAndReturn(run func(admin.DeleteLdapConfigurationApiRequest) (*admin.UserSecurity, *http.Response, error)) *LDAPConfigurationApi_DeleteLdapConfigurationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteLdapConfigurationWithParams provides a mock function with given fields: ctx, args
func (_m *LDAPConfigurationApi) DeleteLdapConfigurationWithParams(ctx context.Context, args *admin.DeleteLdapConfigurationApiParams) admin.DeleteLdapConfigurationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DeleteLdapConfigurationWithParams")
	}

	var r0 admin.DeleteLdapConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeleteLdapConfigurationApiParams) admin.DeleteLdapConfigurationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DeleteLdapConfigurationApiRequest)
	}

	return r0
}

// LDAPConfigurationApi_DeleteLdapConfigurationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteLdapConfigurationWithParams'
type LDAPConfigurationApi_DeleteLdapConfigurationWithParams_Call struct {
	*mock.Call
}

// DeleteLdapConfigurationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DeleteLdapConfigurationApiParams
func (_e *LDAPConfigurationApi_Expecter) DeleteLdapConfigurationWithParams(ctx any, args any) *LDAPConfigurationApi_DeleteLdapConfigurationWithParams_Call {
	return &LDAPConfigurationApi_DeleteLdapConfigurationWithParams_Call{Call: _e.mock.On("DeleteLdapConfigurationWithParams", ctx, args)}
}

func (_c *LDAPConfigurationApi_DeleteLdapConfigurationWithParams_Call) Run(run func(ctx context.Context, args *admin.DeleteLdapConfigurationApiParams)) *LDAPConfigurationApi_DeleteLdapConfigurationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DeleteLdapConfigurationApiParams))
	})
	return _c
}

func (_c *LDAPConfigurationApi_DeleteLdapConfigurationWithParams_Call) Return(_a0 admin.DeleteLdapConfigurationApiRequest) *LDAPConfigurationApi_DeleteLdapConfigurationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LDAPConfigurationApi_DeleteLdapConfigurationWithParams_Call) RunAndReturn(run func(context.Context, *admin.DeleteLdapConfigurationApiParams) admin.DeleteLdapConfigurationApiRequest) *LDAPConfigurationApi_DeleteLdapConfigurationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetLdapConfiguration provides a mock function with given fields: ctx, groupId
func (_m *LDAPConfigurationApi) GetLdapConfiguration(ctx context.Context, groupId string) admin.GetLdapConfigurationApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for GetLdapConfiguration")
	}

	var r0 admin.GetLdapConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.GetLdapConfigurationApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.GetLdapConfigurationApiRequest)
	}

	return r0
}

// LDAPConfigurationApi_GetLdapConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLdapConfiguration'
type LDAPConfigurationApi_GetLdapConfiguration_Call struct {
	*mock.Call
}

// GetLdapConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *LDAPConfigurationApi_Expecter) GetLdapConfiguration(ctx any, groupId any) *LDAPConfigurationApi_GetLdapConfiguration_Call {
	return &LDAPConfigurationApi_GetLdapConfiguration_Call{Call: _e.mock.On("GetLdapConfiguration", ctx, groupId)}
}

func (_c *LDAPConfigurationApi_GetLdapConfiguration_Call) Run(run func(ctx context.Context, groupId string)) *LDAPConfigurationApi_GetLdapConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfiguration_Call) Return(_a0 admin.GetLdapConfigurationApiRequest) *LDAPConfigurationApi_GetLdapConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfiguration_Call) RunAndReturn(run func(context.Context, string) admin.GetLdapConfigurationApiRequest) *LDAPConfigurationApi_GetLdapConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// GetLdapConfigurationExecute provides a mock function with given fields: r
func (_m *LDAPConfigurationApi) GetLdapConfigurationExecute(r admin.GetLdapConfigurationApiRequest) (*admin.UserSecurity, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetLdapConfigurationExecute")
	}

	var r0 *admin.UserSecurity
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetLdapConfigurationApiRequest) (*admin.UserSecurity, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetLdapConfigurationApiRequest) *admin.UserSecurity); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.UserSecurity)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetLdapConfigurationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetLdapConfigurationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LDAPConfigurationApi_GetLdapConfigurationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLdapConfigurationExecute'
type LDAPConfigurationApi_GetLdapConfigurationExecute_Call struct {
	*mock.Call
}

// GetLdapConfigurationExecute is a helper method to define mock.On call
//   - r admin.GetLdapConfigurationApiRequest
func (_e *LDAPConfigurationApi_Expecter) GetLdapConfigurationExecute(r any) *LDAPConfigurationApi_GetLdapConfigurationExecute_Call {
	return &LDAPConfigurationApi_GetLdapConfigurationExecute_Call{Call: _e.mock.On("GetLdapConfigurationExecute", r)}
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationExecute_Call) Run(run func(r admin.GetLdapConfigurationApiRequest)) *LDAPConfigurationApi_GetLdapConfigurationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetLdapConfigurationApiRequest))
	})
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationExecute_Call) Return(_a0 *admin.UserSecurity, _a1 *http.Response, _a2 error) *LDAPConfigurationApi_GetLdapConfigurationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationExecute_Call) RunAndReturn(run func(admin.GetLdapConfigurationApiRequest) (*admin.UserSecurity, *http.Response, error)) *LDAPConfigurationApi_GetLdapConfigurationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetLdapConfigurationStatus provides a mock function with given fields: ctx, groupId, requestId
func (_m *LDAPConfigurationApi) GetLdapConfigurationStatus(ctx context.Context, groupId string, requestId string) admin.GetLdapConfigurationStatusApiRequest {
	ret := _m.Called(ctx, groupId, requestId)

	if len(ret) == 0 {
		panic("no return value specified for GetLdapConfigurationStatus")
	}

	var r0 admin.GetLdapConfigurationStatusApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetLdapConfigurationStatusApiRequest); ok {
		r0 = rf(ctx, groupId, requestId)
	} else {
		r0 = ret.Get(0).(admin.GetLdapConfigurationStatusApiRequest)
	}

	return r0
}

// LDAPConfigurationApi_GetLdapConfigurationStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLdapConfigurationStatus'
type LDAPConfigurationApi_GetLdapConfigurationStatus_Call struct {
	*mock.Call
}

// GetLdapConfigurationStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - requestId string
func (_e *LDAPConfigurationApi_Expecter) GetLdapConfigurationStatus(ctx any, groupId any, requestId any) *LDAPConfigurationApi_GetLdapConfigurationStatus_Call {
	return &LDAPConfigurationApi_GetLdapConfigurationStatus_Call{Call: _e.mock.On("GetLdapConfigurationStatus", ctx, groupId, requestId)}
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationStatus_Call) Run(run func(ctx context.Context, groupId string, requestId string)) *LDAPConfigurationApi_GetLdapConfigurationStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationStatus_Call) Return(_a0 admin.GetLdapConfigurationStatusApiRequest) *LDAPConfigurationApi_GetLdapConfigurationStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationStatus_Call) RunAndReturn(run func(context.Context, string, string) admin.GetLdapConfigurationStatusApiRequest) *LDAPConfigurationApi_GetLdapConfigurationStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetLdapConfigurationStatusExecute provides a mock function with given fields: r
func (_m *LDAPConfigurationApi) GetLdapConfigurationStatusExecute(r admin.GetLdapConfigurationStatusApiRequest) (*admin.LDAPVerifyConnectivityJobRequest, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetLdapConfigurationStatusExecute")
	}

	var r0 *admin.LDAPVerifyConnectivityJobRequest
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetLdapConfigurationStatusApiRequest) (*admin.LDAPVerifyConnectivityJobRequest, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetLdapConfigurationStatusApiRequest) *admin.LDAPVerifyConnectivityJobRequest); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.LDAPVerifyConnectivityJobRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetLdapConfigurationStatusApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetLdapConfigurationStatusApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LDAPConfigurationApi_GetLdapConfigurationStatusExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLdapConfigurationStatusExecute'
type LDAPConfigurationApi_GetLdapConfigurationStatusExecute_Call struct {
	*mock.Call
}

// GetLdapConfigurationStatusExecute is a helper method to define mock.On call
//   - r admin.GetLdapConfigurationStatusApiRequest
func (_e *LDAPConfigurationApi_Expecter) GetLdapConfigurationStatusExecute(r any) *LDAPConfigurationApi_GetLdapConfigurationStatusExecute_Call {
	return &LDAPConfigurationApi_GetLdapConfigurationStatusExecute_Call{Call: _e.mock.On("GetLdapConfigurationStatusExecute", r)}
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationStatusExecute_Call) Run(run func(r admin.GetLdapConfigurationStatusApiRequest)) *LDAPConfigurationApi_GetLdapConfigurationStatusExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetLdapConfigurationStatusApiRequest))
	})
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationStatusExecute_Call) Return(_a0 *admin.LDAPVerifyConnectivityJobRequest, _a1 *http.Response, _a2 error) *LDAPConfigurationApi_GetLdapConfigurationStatusExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationStatusExecute_Call) RunAndReturn(run func(admin.GetLdapConfigurationStatusApiRequest) (*admin.LDAPVerifyConnectivityJobRequest, *http.Response, error)) *LDAPConfigurationApi_GetLdapConfigurationStatusExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetLdapConfigurationStatusWithParams provides a mock function with given fields: ctx, args
func (_m *LDAPConfigurationApi) GetLdapConfigurationStatusWithParams(ctx context.Context, args *admin.GetLdapConfigurationStatusApiParams) admin.GetLdapConfigurationStatusApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetLdapConfigurationStatusWithParams")
	}

	var r0 admin.GetLdapConfigurationStatusApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetLdapConfigurationStatusApiParams) admin.GetLdapConfigurationStatusApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetLdapConfigurationStatusApiRequest)
	}

	return r0
}

// LDAPConfigurationApi_GetLdapConfigurationStatusWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLdapConfigurationStatusWithParams'
type LDAPConfigurationApi_GetLdapConfigurationStatusWithParams_Call struct {
	*mock.Call
}

// GetLdapConfigurationStatusWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetLdapConfigurationStatusApiParams
func (_e *LDAPConfigurationApi_Expecter) GetLdapConfigurationStatusWithParams(ctx any, args any) *LDAPConfigurationApi_GetLdapConfigurationStatusWithParams_Call {
	return &LDAPConfigurationApi_GetLdapConfigurationStatusWithParams_Call{Call: _e.mock.On("GetLdapConfigurationStatusWithParams", ctx, args)}
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationStatusWithParams_Call) Run(run func(ctx context.Context, args *admin.GetLdapConfigurationStatusApiParams)) *LDAPConfigurationApi_GetLdapConfigurationStatusWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetLdapConfigurationStatusApiParams))
	})
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationStatusWithParams_Call) Return(_a0 admin.GetLdapConfigurationStatusApiRequest) *LDAPConfigurationApi_GetLdapConfigurationStatusWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationStatusWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetLdapConfigurationStatusApiParams) admin.GetLdapConfigurationStatusApiRequest) *LDAPConfigurationApi_GetLdapConfigurationStatusWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetLdapConfigurationWithParams provides a mock function with given fields: ctx, args
func (_m *LDAPConfigurationApi) GetLdapConfigurationWithParams(ctx context.Context, args *admin.GetLdapConfigurationApiParams) admin.GetLdapConfigurationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetLdapConfigurationWithParams")
	}

	var r0 admin.GetLdapConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetLdapConfigurationApiParams) admin.GetLdapConfigurationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetLdapConfigurationApiRequest)
	}

	return r0
}

// LDAPConfigurationApi_GetLdapConfigurationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLdapConfigurationWithParams'
type LDAPConfigurationApi_GetLdapConfigurationWithParams_Call struct {
	*mock.Call
}

// GetLdapConfigurationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetLdapConfigurationApiParams
func (_e *LDAPConfigurationApi_Expecter) GetLdapConfigurationWithParams(ctx any, args any) *LDAPConfigurationApi_GetLdapConfigurationWithParams_Call {
	return &LDAPConfigurationApi_GetLdapConfigurationWithParams_Call{Call: _e.mock.On("GetLdapConfigurationWithParams", ctx, args)}
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationWithParams_Call) Run(run func(ctx context.Context, args *admin.GetLdapConfigurationApiParams)) *LDAPConfigurationApi_GetLdapConfigurationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetLdapConfigurationApiParams))
	})
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationWithParams_Call) Return(_a0 admin.GetLdapConfigurationApiRequest) *LDAPConfigurationApi_GetLdapConfigurationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LDAPConfigurationApi_GetLdapConfigurationWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetLdapConfigurationApiParams) admin.GetLdapConfigurationApiRequest) *LDAPConfigurationApi_GetLdapConfigurationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// SaveLdapConfiguration provides a mock function with given fields: ctx, groupId, userSecurity
func (_m *LDAPConfigurationApi) SaveLdapConfiguration(ctx context.Context, groupId string, userSecurity *admin.UserSecurity) admin.SaveLdapConfigurationApiRequest {
	ret := _m.Called(ctx, groupId, userSecurity)

	if len(ret) == 0 {
		panic("no return value specified for SaveLdapConfiguration")
	}

	var r0 admin.SaveLdapConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.UserSecurity) admin.SaveLdapConfigurationApiRequest); ok {
		r0 = rf(ctx, groupId, userSecurity)
	} else {
		r0 = ret.Get(0).(admin.SaveLdapConfigurationApiRequest)
	}

	return r0
}

// LDAPConfigurationApi_SaveLdapConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveLdapConfiguration'
type LDAPConfigurationApi_SaveLdapConfiguration_Call struct {
	*mock.Call
}

// SaveLdapConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - userSecurity *admin.UserSecurity
func (_e *LDAPConfigurationApi_Expecter) SaveLdapConfiguration(ctx any, groupId any, userSecurity any) *LDAPConfigurationApi_SaveLdapConfiguration_Call {
	return &LDAPConfigurationApi_SaveLdapConfiguration_Call{Call: _e.mock.On("SaveLdapConfiguration", ctx, groupId, userSecurity)}
}

func (_c *LDAPConfigurationApi_SaveLdapConfiguration_Call) Run(run func(ctx context.Context, groupId string, userSecurity *admin.UserSecurity)) *LDAPConfigurationApi_SaveLdapConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.UserSecurity))
	})
	return _c
}

func (_c *LDAPConfigurationApi_SaveLdapConfiguration_Call) Return(_a0 admin.SaveLdapConfigurationApiRequest) *LDAPConfigurationApi_SaveLdapConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LDAPConfigurationApi_SaveLdapConfiguration_Call) RunAndReturn(run func(context.Context, string, *admin.UserSecurity) admin.SaveLdapConfigurationApiRequest) *LDAPConfigurationApi_SaveLdapConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// SaveLdapConfigurationExecute provides a mock function with given fields: r
func (_m *LDAPConfigurationApi) SaveLdapConfigurationExecute(r admin.SaveLdapConfigurationApiRequest) (*admin.UserSecurity, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for SaveLdapConfigurationExecute")
	}

	var r0 *admin.UserSecurity
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.SaveLdapConfigurationApiRequest) (*admin.UserSecurity, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.SaveLdapConfigurationApiRequest) *admin.UserSecurity); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.UserSecurity)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.SaveLdapConfigurationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.SaveLdapConfigurationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LDAPConfigurationApi_SaveLdapConfigurationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveLdapConfigurationExecute'
type LDAPConfigurationApi_SaveLdapConfigurationExecute_Call struct {
	*mock.Call
}

// SaveLdapConfigurationExecute is a helper method to define mock.On call
//   - r admin.SaveLdapConfigurationApiRequest
func (_e *LDAPConfigurationApi_Expecter) SaveLdapConfigurationExecute(r any) *LDAPConfigurationApi_SaveLdapConfigurationExecute_Call {
	return &LDAPConfigurationApi_SaveLdapConfigurationExecute_Call{Call: _e.mock.On("SaveLdapConfigurationExecute", r)}
}

func (_c *LDAPConfigurationApi_SaveLdapConfigurationExecute_Call) Run(run func(r admin.SaveLdapConfigurationApiRequest)) *LDAPConfigurationApi_SaveLdapConfigurationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.SaveLdapConfigurationApiRequest))
	})
	return _c
}

func (_c *LDAPConfigurationApi_SaveLdapConfigurationExecute_Call) Return(_a0 *admin.UserSecurity, _a1 *http.Response, _a2 error) *LDAPConfigurationApi_SaveLdapConfigurationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *LDAPConfigurationApi_SaveLdapConfigurationExecute_Call) RunAndReturn(run func(admin.SaveLdapConfigurationApiRequest) (*admin.UserSecurity, *http.Response, error)) *LDAPConfigurationApi_SaveLdapConfigurationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// SaveLdapConfigurationWithParams provides a mock function with given fields: ctx, args
func (_m *LDAPConfigurationApi) SaveLdapConfigurationWithParams(ctx context.Context, args *admin.SaveLdapConfigurationApiParams) admin.SaveLdapConfigurationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for SaveLdapConfigurationWithParams")
	}

	var r0 admin.SaveLdapConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.SaveLdapConfigurationApiParams) admin.SaveLdapConfigurationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.SaveLdapConfigurationApiRequest)
	}

	return r0
}

// LDAPConfigurationApi_SaveLdapConfigurationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveLdapConfigurationWithParams'
type LDAPConfigurationApi_SaveLdapConfigurationWithParams_Call struct {
	*mock.Call
}

// SaveLdapConfigurationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.SaveLdapConfigurationApiParams
func (_e *LDAPConfigurationApi_Expecter) SaveLdapConfigurationWithParams(ctx any, args any) *LDAPConfigurationApi_SaveLdapConfigurationWithParams_Call {
	return &LDAPConfigurationApi_SaveLdapConfigurationWithParams_Call{Call: _e.mock.On("SaveLdapConfigurationWithParams", ctx, args)}
}

func (_c *LDAPConfigurationApi_SaveLdapConfigurationWithParams_Call) Run(run func(ctx context.Context, args *admin.SaveLdapConfigurationApiParams)) *LDAPConfigurationApi_SaveLdapConfigurationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.SaveLdapConfigurationApiParams))
	})
	return _c
}

func (_c *LDAPConfigurationApi_SaveLdapConfigurationWithParams_Call) Return(_a0 admin.SaveLdapConfigurationApiRequest) *LDAPConfigurationApi_SaveLdapConfigurationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LDAPConfigurationApi_SaveLdapConfigurationWithParams_Call) RunAndReturn(run func(context.Context, *admin.SaveLdapConfigurationApiParams) admin.SaveLdapConfigurationApiRequest) *LDAPConfigurationApi_SaveLdapConfigurationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyLdapConfiguration provides a mock function with given fields: ctx, groupId, lDAPVerifyConnectivityJobRequestParams
func (_m *LDAPConfigurationApi) VerifyLdapConfiguration(ctx context.Context, groupId string, lDAPVerifyConnectivityJobRequestParams *admin.LDAPVerifyConnectivityJobRequestParams) admin.VerifyLdapConfigurationApiRequest {
	ret := _m.Called(ctx, groupId, lDAPVerifyConnectivityJobRequestParams)

	if len(ret) == 0 {
		panic("no return value specified for VerifyLdapConfiguration")
	}

	var r0 admin.VerifyLdapConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.LDAPVerifyConnectivityJobRequestParams) admin.VerifyLdapConfigurationApiRequest); ok {
		r0 = rf(ctx, groupId, lDAPVerifyConnectivityJobRequestParams)
	} else {
		r0 = ret.Get(0).(admin.VerifyLdapConfigurationApiRequest)
	}

	return r0
}

// LDAPConfigurationApi_VerifyLdapConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyLdapConfiguration'
type LDAPConfigurationApi_VerifyLdapConfiguration_Call struct {
	*mock.Call
}

// VerifyLdapConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - lDAPVerifyConnectivityJobRequestParams *admin.LDAPVerifyConnectivityJobRequestParams
func (_e *LDAPConfigurationApi_Expecter) VerifyLdapConfiguration(ctx any, groupId any, lDAPVerifyConnectivityJobRequestParams any) *LDAPConfigurationApi_VerifyLdapConfiguration_Call {
	return &LDAPConfigurationApi_VerifyLdapConfiguration_Call{Call: _e.mock.On("VerifyLdapConfiguration", ctx, groupId, lDAPVerifyConnectivityJobRequestParams)}
}

func (_c *LDAPConfigurationApi_VerifyLdapConfiguration_Call) Run(run func(ctx context.Context, groupId string, lDAPVerifyConnectivityJobRequestParams *admin.LDAPVerifyConnectivityJobRequestParams)) *LDAPConfigurationApi_VerifyLdapConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.LDAPVerifyConnectivityJobRequestParams))
	})
	return _c
}

func (_c *LDAPConfigurationApi_VerifyLdapConfiguration_Call) Return(_a0 admin.VerifyLdapConfigurationApiRequest) *LDAPConfigurationApi_VerifyLdapConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LDAPConfigurationApi_VerifyLdapConfiguration_Call) RunAndReturn(run func(context.Context, string, *admin.LDAPVerifyConnectivityJobRequestParams) admin.VerifyLdapConfigurationApiRequest) *LDAPConfigurationApi_VerifyLdapConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyLdapConfigurationExecute provides a mock function with given fields: r
func (_m *LDAPConfigurationApi) VerifyLdapConfigurationExecute(r admin.VerifyLdapConfigurationApiRequest) (*admin.LDAPVerifyConnectivityJobRequest, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for VerifyLdapConfigurationExecute")
	}

	var r0 *admin.LDAPVerifyConnectivityJobRequest
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.VerifyLdapConfigurationApiRequest) (*admin.LDAPVerifyConnectivityJobRequest, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.VerifyLdapConfigurationApiRequest) *admin.LDAPVerifyConnectivityJobRequest); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.LDAPVerifyConnectivityJobRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.VerifyLdapConfigurationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.VerifyLdapConfigurationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LDAPConfigurationApi_VerifyLdapConfigurationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyLdapConfigurationExecute'
type LDAPConfigurationApi_VerifyLdapConfigurationExecute_Call struct {
	*mock.Call
}

// VerifyLdapConfigurationExecute is a helper method to define mock.On call
//   - r admin.VerifyLdapConfigurationApiRequest
func (_e *LDAPConfigurationApi_Expecter) VerifyLdapConfigurationExecute(r any) *LDAPConfigurationApi_VerifyLdapConfigurationExecute_Call {
	return &LDAPConfigurationApi_VerifyLdapConfigurationExecute_Call{Call: _e.mock.On("VerifyLdapConfigurationExecute", r)}
}

func (_c *LDAPConfigurationApi_VerifyLdapConfigurationExecute_Call) Run(run func(r admin.VerifyLdapConfigurationApiRequest)) *LDAPConfigurationApi_VerifyLdapConfigurationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.VerifyLdapConfigurationApiRequest))
	})
	return _c
}

func (_c *LDAPConfigurationApi_VerifyLdapConfigurationExecute_Call) Return(_a0 *admin.LDAPVerifyConnectivityJobRequest, _a1 *http.Response, _a2 error) *LDAPConfigurationApi_VerifyLdapConfigurationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *LDAPConfigurationApi_VerifyLdapConfigurationExecute_Call) RunAndReturn(run func(admin.VerifyLdapConfigurationApiRequest) (*admin.LDAPVerifyConnectivityJobRequest, *http.Response, error)) *LDAPConfigurationApi_VerifyLdapConfigurationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyLdapConfigurationWithParams provides a mock function with given fields: ctx, args
func (_m *LDAPConfigurationApi) VerifyLdapConfigurationWithParams(ctx context.Context, args *admin.VerifyLdapConfigurationApiParams) admin.VerifyLdapConfigurationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for VerifyLdapConfigurationWithParams")
	}

	var r0 admin.VerifyLdapConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.VerifyLdapConfigurationApiParams) admin.VerifyLdapConfigurationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.VerifyLdapConfigurationApiRequest)
	}

	return r0
}

// LDAPConfigurationApi_VerifyLdapConfigurationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyLdapConfigurationWithParams'
type LDAPConfigurationApi_VerifyLdapConfigurationWithParams_Call struct {
	*mock.Call
}

// VerifyLdapConfigurationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.VerifyLdapConfigurationApiParams
func (_e *LDAPConfigurationApi_Expecter) VerifyLdapConfigurationWithParams(ctx any, args any) *LDAPConfigurationApi_VerifyLdapConfigurationWithParams_Call {
	return &LDAPConfigurationApi_VerifyLdapConfigurationWithParams_Call{Call: _e.mock.On("VerifyLdapConfigurationWithParams", ctx, args)}
}

func (_c *LDAPConfigurationApi_VerifyLdapConfigurationWithParams_Call) Run(run func(ctx context.Context, args *admin.VerifyLdapConfigurationApiParams)) *LDAPConfigurationApi_VerifyLdapConfigurationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.VerifyLdapConfigurationApiParams))
	})
	return _c
}

func (_c *LDAPConfigurationApi_VerifyLdapConfigurationWithParams_Call) Return(_a0 admin.VerifyLdapConfigurationApiRequest) *LDAPConfigurationApi_VerifyLdapConfigurationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LDAPConfigurationApi_VerifyLdapConfigurationWithParams_Call) RunAndReturn(run func(context.Context, *admin.VerifyLdapConfigurationApiParams) admin.VerifyLdapConfigurationApiRequest) *LDAPConfigurationApi_VerifyLdapConfigurationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewLDAPConfigurationApi creates a new instance of LDAPConfigurationApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLDAPConfigurationApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *LDAPConfigurationApi {
	mock := &LDAPConfigurationApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
