// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20231115008/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// AuditingApi is an autogenerated mock type for the AuditingApi type
type AuditingApi struct {
	mock.Mock
}

type AuditingApi_Expecter struct {
	mock *mock.Mock
}

func (_m *AuditingApi) EXPECT() *AuditingApi_Expecter {
	return &AuditingApi_Expecter{mock: &_m.Mock}
}

// GetAuditingConfiguration provides a mock function with given fields: ctx, groupId
func (_m *AuditingApi) GetAuditingConfiguration(ctx context.Context, groupId string) admin.GetAuditingConfigurationApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for GetAuditingConfiguration")
	}

	var r0 admin.GetAuditingConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.GetAuditingConfigurationApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.GetAuditingConfigurationApiRequest)
	}

	return r0
}

// AuditingApi_GetAuditingConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuditingConfiguration'
type AuditingApi_GetAuditingConfiguration_Call struct {
	*mock.Call
}

// GetAuditingConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *AuditingApi_Expecter) GetAuditingConfiguration(ctx interface{}, groupId interface{}) *AuditingApi_GetAuditingConfiguration_Call {
	return &AuditingApi_GetAuditingConfiguration_Call{Call: _e.mock.On("GetAuditingConfiguration", ctx, groupId)}
}

func (_c *AuditingApi_GetAuditingConfiguration_Call) Run(run func(ctx context.Context, groupId string)) *AuditingApi_GetAuditingConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuditingApi_GetAuditingConfiguration_Call) Return(_a0 admin.GetAuditingConfigurationApiRequest) *AuditingApi_GetAuditingConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuditingApi_GetAuditingConfiguration_Call) RunAndReturn(run func(context.Context, string) admin.GetAuditingConfigurationApiRequest) *AuditingApi_GetAuditingConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuditingConfigurationExecute provides a mock function with given fields: r
func (_m *AuditingApi) GetAuditingConfigurationExecute(r admin.GetAuditingConfigurationApiRequest) (*admin.AuditLog, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetAuditingConfigurationExecute")
	}

	var r0 *admin.AuditLog
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetAuditingConfigurationApiRequest) (*admin.AuditLog, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetAuditingConfigurationApiRequest) *admin.AuditLog); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.AuditLog)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetAuditingConfigurationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetAuditingConfigurationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AuditingApi_GetAuditingConfigurationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuditingConfigurationExecute'
type AuditingApi_GetAuditingConfigurationExecute_Call struct {
	*mock.Call
}

// GetAuditingConfigurationExecute is a helper method to define mock.On call
//   - r admin.GetAuditingConfigurationApiRequest
func (_e *AuditingApi_Expecter) GetAuditingConfigurationExecute(r interface{}) *AuditingApi_GetAuditingConfigurationExecute_Call {
	return &AuditingApi_GetAuditingConfigurationExecute_Call{Call: _e.mock.On("GetAuditingConfigurationExecute", r)}
}

func (_c *AuditingApi_GetAuditingConfigurationExecute_Call) Run(run func(r admin.GetAuditingConfigurationApiRequest)) *AuditingApi_GetAuditingConfigurationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetAuditingConfigurationApiRequest))
	})
	return _c
}

func (_c *AuditingApi_GetAuditingConfigurationExecute_Call) Return(_a0 *admin.AuditLog, _a1 *http.Response, _a2 error) *AuditingApi_GetAuditingConfigurationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *AuditingApi_GetAuditingConfigurationExecute_Call) RunAndReturn(run func(admin.GetAuditingConfigurationApiRequest) (*admin.AuditLog, *http.Response, error)) *AuditingApi_GetAuditingConfigurationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuditingConfigurationWithParams provides a mock function with given fields: ctx, args
func (_m *AuditingApi) GetAuditingConfigurationWithParams(ctx context.Context, args *admin.GetAuditingConfigurationApiParams) admin.GetAuditingConfigurationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetAuditingConfigurationWithParams")
	}

	var r0 admin.GetAuditingConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetAuditingConfigurationApiParams) admin.GetAuditingConfigurationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetAuditingConfigurationApiRequest)
	}

	return r0
}

// AuditingApi_GetAuditingConfigurationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuditingConfigurationWithParams'
type AuditingApi_GetAuditingConfigurationWithParams_Call struct {
	*mock.Call
}

// GetAuditingConfigurationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetAuditingConfigurationApiParams
func (_e *AuditingApi_Expecter) GetAuditingConfigurationWithParams(ctx interface{}, args interface{}) *AuditingApi_GetAuditingConfigurationWithParams_Call {
	return &AuditingApi_GetAuditingConfigurationWithParams_Call{Call: _e.mock.On("GetAuditingConfigurationWithParams", ctx, args)}
}

func (_c *AuditingApi_GetAuditingConfigurationWithParams_Call) Run(run func(ctx context.Context, args *admin.GetAuditingConfigurationApiParams)) *AuditingApi_GetAuditingConfigurationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetAuditingConfigurationApiParams))
	})
	return _c
}

func (_c *AuditingApi_GetAuditingConfigurationWithParams_Call) Return(_a0 admin.GetAuditingConfigurationApiRequest) *AuditingApi_GetAuditingConfigurationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuditingApi_GetAuditingConfigurationWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetAuditingConfigurationApiParams) admin.GetAuditingConfigurationApiRequest) *AuditingApi_GetAuditingConfigurationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAuditingConfiguration provides a mock function with given fields: ctx, groupId, auditLog
func (_m *AuditingApi) UpdateAuditingConfiguration(ctx context.Context, groupId string, auditLog *admin.AuditLog) admin.UpdateAuditingConfigurationApiRequest {
	ret := _m.Called(ctx, groupId, auditLog)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAuditingConfiguration")
	}

	var r0 admin.UpdateAuditingConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.AuditLog) admin.UpdateAuditingConfigurationApiRequest); ok {
		r0 = rf(ctx, groupId, auditLog)
	} else {
		r0 = ret.Get(0).(admin.UpdateAuditingConfigurationApiRequest)
	}

	return r0
}

// AuditingApi_UpdateAuditingConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAuditingConfiguration'
type AuditingApi_UpdateAuditingConfiguration_Call struct {
	*mock.Call
}

// UpdateAuditingConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - auditLog *admin.AuditLog
func (_e *AuditingApi_Expecter) UpdateAuditingConfiguration(ctx interface{}, groupId interface{}, auditLog interface{}) *AuditingApi_UpdateAuditingConfiguration_Call {
	return &AuditingApi_UpdateAuditingConfiguration_Call{Call: _e.mock.On("UpdateAuditingConfiguration", ctx, groupId, auditLog)}
}

func (_c *AuditingApi_UpdateAuditingConfiguration_Call) Run(run func(ctx context.Context, groupId string, auditLog *admin.AuditLog)) *AuditingApi_UpdateAuditingConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.AuditLog))
	})
	return _c
}

func (_c *AuditingApi_UpdateAuditingConfiguration_Call) Return(_a0 admin.UpdateAuditingConfigurationApiRequest) *AuditingApi_UpdateAuditingConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuditingApi_UpdateAuditingConfiguration_Call) RunAndReturn(run func(context.Context, string, *admin.AuditLog) admin.UpdateAuditingConfigurationApiRequest) *AuditingApi_UpdateAuditingConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAuditingConfigurationExecute provides a mock function with given fields: r
func (_m *AuditingApi) UpdateAuditingConfigurationExecute(r admin.UpdateAuditingConfigurationApiRequest) (*admin.AuditLog, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAuditingConfigurationExecute")
	}

	var r0 *admin.AuditLog
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.UpdateAuditingConfigurationApiRequest) (*admin.AuditLog, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.UpdateAuditingConfigurationApiRequest) *admin.AuditLog); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.AuditLog)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.UpdateAuditingConfigurationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.UpdateAuditingConfigurationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AuditingApi_UpdateAuditingConfigurationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAuditingConfigurationExecute'
type AuditingApi_UpdateAuditingConfigurationExecute_Call struct {
	*mock.Call
}

// UpdateAuditingConfigurationExecute is a helper method to define mock.On call
//   - r admin.UpdateAuditingConfigurationApiRequest
func (_e *AuditingApi_Expecter) UpdateAuditingConfigurationExecute(r interface{}) *AuditingApi_UpdateAuditingConfigurationExecute_Call {
	return &AuditingApi_UpdateAuditingConfigurationExecute_Call{Call: _e.mock.On("UpdateAuditingConfigurationExecute", r)}
}

func (_c *AuditingApi_UpdateAuditingConfigurationExecute_Call) Run(run func(r admin.UpdateAuditingConfigurationApiRequest)) *AuditingApi_UpdateAuditingConfigurationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.UpdateAuditingConfigurationApiRequest))
	})
	return _c
}

func (_c *AuditingApi_UpdateAuditingConfigurationExecute_Call) Return(_a0 *admin.AuditLog, _a1 *http.Response, _a2 error) *AuditingApi_UpdateAuditingConfigurationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *AuditingApi_UpdateAuditingConfigurationExecute_Call) RunAndReturn(run func(admin.UpdateAuditingConfigurationApiRequest) (*admin.AuditLog, *http.Response, error)) *AuditingApi_UpdateAuditingConfigurationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAuditingConfigurationWithParams provides a mock function with given fields: ctx, args
func (_m *AuditingApi) UpdateAuditingConfigurationWithParams(ctx context.Context, args *admin.UpdateAuditingConfigurationApiParams) admin.UpdateAuditingConfigurationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAuditingConfigurationWithParams")
	}

	var r0 admin.UpdateAuditingConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.UpdateAuditingConfigurationApiParams) admin.UpdateAuditingConfigurationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.UpdateAuditingConfigurationApiRequest)
	}

	return r0
}

// AuditingApi_UpdateAuditingConfigurationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAuditingConfigurationWithParams'
type AuditingApi_UpdateAuditingConfigurationWithParams_Call struct {
	*mock.Call
}

// UpdateAuditingConfigurationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.UpdateAuditingConfigurationApiParams
func (_e *AuditingApi_Expecter) UpdateAuditingConfigurationWithParams(ctx interface{}, args interface{}) *AuditingApi_UpdateAuditingConfigurationWithParams_Call {
	return &AuditingApi_UpdateAuditingConfigurationWithParams_Call{Call: _e.mock.On("UpdateAuditingConfigurationWithParams", ctx, args)}
}

func (_c *AuditingApi_UpdateAuditingConfigurationWithParams_Call) Run(run func(ctx context.Context, args *admin.UpdateAuditingConfigurationApiParams)) *AuditingApi_UpdateAuditingConfigurationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.UpdateAuditingConfigurationApiParams))
	})
	return _c
}

func (_c *AuditingApi_UpdateAuditingConfigurationWithParams_Call) Return(_a0 admin.UpdateAuditingConfigurationApiRequest) *AuditingApi_UpdateAuditingConfigurationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuditingApi_UpdateAuditingConfigurationWithParams_Call) RunAndReturn(run func(context.Context, *admin.UpdateAuditingConfigurationApiParams) admin.UpdateAuditingConfigurationApiRequest) *AuditingApi_UpdateAuditingConfigurationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuditingApi creates a new instance of AuditingApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuditingApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuditingApi {
	mock := &AuditingApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
