// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20240805001/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// EncryptionAtRestUsingCustomerKeyManagementApi is an autogenerated mock type for the EncryptionAtRestUsingCustomerKeyManagementApi type
type EncryptionAtRestUsingCustomerKeyManagementApi struct {
	mock.Mock
}

type EncryptionAtRestUsingCustomerKeyManagementApi_Expecter struct {
	mock *mock.Mock
}

func (_m *EncryptionAtRestUsingCustomerKeyManagementApi) EXPECT() *EncryptionAtRestUsingCustomerKeyManagementApi_Expecter {
	return &EncryptionAtRestUsingCustomerKeyManagementApi_Expecter{mock: &_m.Mock}
}

// GetEncryptionAtRest provides a mock function with given fields: ctx, groupId
func (_m *EncryptionAtRestUsingCustomerKeyManagementApi) GetEncryptionAtRest(ctx context.Context, groupId string) admin.GetEncryptionAtRestApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for GetEncryptionAtRest")
	}

	var r0 admin.GetEncryptionAtRestApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.GetEncryptionAtRestApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.GetEncryptionAtRestApiRequest)
	}

	return r0
}

// EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEncryptionAtRest'
type EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRest_Call struct {
	*mock.Call
}

// GetEncryptionAtRest is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *EncryptionAtRestUsingCustomerKeyManagementApi_Expecter) GetEncryptionAtRest(ctx interface{}, groupId interface{}) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRest_Call {
	return &EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRest_Call{Call: _e.mock.On("GetEncryptionAtRest", ctx, groupId)}
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRest_Call) Run(run func(ctx context.Context, groupId string)) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRest_Call) Return(_a0 admin.GetEncryptionAtRestApiRequest) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRest_Call) RunAndReturn(run func(context.Context, string) admin.GetEncryptionAtRestApiRequest) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRest_Call {
	_c.Call.Return(run)
	return _c
}

// GetEncryptionAtRestExecute provides a mock function with given fields: r
func (_m *EncryptionAtRestUsingCustomerKeyManagementApi) GetEncryptionAtRestExecute(r admin.GetEncryptionAtRestApiRequest) (*admin.EncryptionAtRest, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetEncryptionAtRestExecute")
	}

	var r0 *admin.EncryptionAtRest
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetEncryptionAtRestApiRequest) (*admin.EncryptionAtRest, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetEncryptionAtRestApiRequest) *admin.EncryptionAtRest); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.EncryptionAtRest)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetEncryptionAtRestApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetEncryptionAtRestApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEncryptionAtRestExecute'
type EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestExecute_Call struct {
	*mock.Call
}

// GetEncryptionAtRestExecute is a helper method to define mock.On call
//   - r admin.GetEncryptionAtRestApiRequest
func (_e *EncryptionAtRestUsingCustomerKeyManagementApi_Expecter) GetEncryptionAtRestExecute(r interface{}) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestExecute_Call {
	return &EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestExecute_Call{Call: _e.mock.On("GetEncryptionAtRestExecute", r)}
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestExecute_Call) Run(run func(r admin.GetEncryptionAtRestApiRequest)) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetEncryptionAtRestApiRequest))
	})
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestExecute_Call) Return(_a0 *admin.EncryptionAtRest, _a1 *http.Response, _a2 error) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestExecute_Call) RunAndReturn(run func(admin.GetEncryptionAtRestApiRequest) (*admin.EncryptionAtRest, *http.Response, error)) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetEncryptionAtRestWithParams provides a mock function with given fields: ctx, args
func (_m *EncryptionAtRestUsingCustomerKeyManagementApi) GetEncryptionAtRestWithParams(ctx context.Context, args *admin.GetEncryptionAtRestApiParams) admin.GetEncryptionAtRestApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetEncryptionAtRestWithParams")
	}

	var r0 admin.GetEncryptionAtRestApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetEncryptionAtRestApiParams) admin.GetEncryptionAtRestApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetEncryptionAtRestApiRequest)
	}

	return r0
}

// EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEncryptionAtRestWithParams'
type EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestWithParams_Call struct {
	*mock.Call
}

// GetEncryptionAtRestWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetEncryptionAtRestApiParams
func (_e *EncryptionAtRestUsingCustomerKeyManagementApi_Expecter) GetEncryptionAtRestWithParams(ctx interface{}, args interface{}) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestWithParams_Call {
	return &EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestWithParams_Call{Call: _e.mock.On("GetEncryptionAtRestWithParams", ctx, args)}
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestWithParams_Call) Run(run func(ctx context.Context, args *admin.GetEncryptionAtRestApiParams)) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetEncryptionAtRestApiParams))
	})
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestWithParams_Call) Return(_a0 admin.GetEncryptionAtRestApiRequest) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetEncryptionAtRestApiParams) admin.GetEncryptionAtRestApiRequest) *EncryptionAtRestUsingCustomerKeyManagementApi_GetEncryptionAtRestWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateEncryptionAtRest provides a mock function with given fields: ctx, groupId, encryptionAtRest
func (_m *EncryptionAtRestUsingCustomerKeyManagementApi) UpdateEncryptionAtRest(ctx context.Context, groupId string, encryptionAtRest *admin.EncryptionAtRest) admin.UpdateEncryptionAtRestApiRequest {
	ret := _m.Called(ctx, groupId, encryptionAtRest)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEncryptionAtRest")
	}

	var r0 admin.UpdateEncryptionAtRestApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.EncryptionAtRest) admin.UpdateEncryptionAtRestApiRequest); ok {
		r0 = rf(ctx, groupId, encryptionAtRest)
	} else {
		r0 = ret.Get(0).(admin.UpdateEncryptionAtRestApiRequest)
	}

	return r0
}

// EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEncryptionAtRest'
type EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRest_Call struct {
	*mock.Call
}

// UpdateEncryptionAtRest is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - encryptionAtRest *admin.EncryptionAtRest
func (_e *EncryptionAtRestUsingCustomerKeyManagementApi_Expecter) UpdateEncryptionAtRest(ctx interface{}, groupId interface{}, encryptionAtRest interface{}) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRest_Call {
	return &EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRest_Call{Call: _e.mock.On("UpdateEncryptionAtRest", ctx, groupId, encryptionAtRest)}
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRest_Call) Run(run func(ctx context.Context, groupId string, encryptionAtRest *admin.EncryptionAtRest)) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.EncryptionAtRest))
	})
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRest_Call) Return(_a0 admin.UpdateEncryptionAtRestApiRequest) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRest_Call) RunAndReturn(run func(context.Context, string, *admin.EncryptionAtRest) admin.UpdateEncryptionAtRestApiRequest) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRest_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateEncryptionAtRestExecute provides a mock function with given fields: r
func (_m *EncryptionAtRestUsingCustomerKeyManagementApi) UpdateEncryptionAtRestExecute(r admin.UpdateEncryptionAtRestApiRequest) (*admin.EncryptionAtRest, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEncryptionAtRestExecute")
	}

	var r0 *admin.EncryptionAtRest
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.UpdateEncryptionAtRestApiRequest) (*admin.EncryptionAtRest, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.UpdateEncryptionAtRestApiRequest) *admin.EncryptionAtRest); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.EncryptionAtRest)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.UpdateEncryptionAtRestApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.UpdateEncryptionAtRestApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEncryptionAtRestExecute'
type EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestExecute_Call struct {
	*mock.Call
}

// UpdateEncryptionAtRestExecute is a helper method to define mock.On call
//   - r admin.UpdateEncryptionAtRestApiRequest
func (_e *EncryptionAtRestUsingCustomerKeyManagementApi_Expecter) UpdateEncryptionAtRestExecute(r interface{}) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestExecute_Call {
	return &EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestExecute_Call{Call: _e.mock.On("UpdateEncryptionAtRestExecute", r)}
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestExecute_Call) Run(run func(r admin.UpdateEncryptionAtRestApiRequest)) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.UpdateEncryptionAtRestApiRequest))
	})
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestExecute_Call) Return(_a0 *admin.EncryptionAtRest, _a1 *http.Response, _a2 error) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestExecute_Call) RunAndReturn(run func(admin.UpdateEncryptionAtRestApiRequest) (*admin.EncryptionAtRest, *http.Response, error)) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestExecute_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateEncryptionAtRestWithParams provides a mock function with given fields: ctx, args
func (_m *EncryptionAtRestUsingCustomerKeyManagementApi) UpdateEncryptionAtRestWithParams(ctx context.Context, args *admin.UpdateEncryptionAtRestApiParams) admin.UpdateEncryptionAtRestApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEncryptionAtRestWithParams")
	}

	var r0 admin.UpdateEncryptionAtRestApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.UpdateEncryptionAtRestApiParams) admin.UpdateEncryptionAtRestApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.UpdateEncryptionAtRestApiRequest)
	}

	return r0
}

// EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEncryptionAtRestWithParams'
type EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestWithParams_Call struct {
	*mock.Call
}

// UpdateEncryptionAtRestWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.UpdateEncryptionAtRestApiParams
func (_e *EncryptionAtRestUsingCustomerKeyManagementApi_Expecter) UpdateEncryptionAtRestWithParams(ctx interface{}, args interface{}) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestWithParams_Call {
	return &EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestWithParams_Call{Call: _e.mock.On("UpdateEncryptionAtRestWithParams", ctx, args)}
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestWithParams_Call) Run(run func(ctx context.Context, args *admin.UpdateEncryptionAtRestApiParams)) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.UpdateEncryptionAtRestApiParams))
	})
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestWithParams_Call) Return(_a0 admin.UpdateEncryptionAtRestApiRequest) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestWithParams_Call) RunAndReturn(run func(context.Context, *admin.UpdateEncryptionAtRestApiParams) admin.UpdateEncryptionAtRestApiRequest) *EncryptionAtRestUsingCustomerKeyManagementApi_UpdateEncryptionAtRestWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewEncryptionAtRestUsingCustomerKeyManagementApi creates a new instance of EncryptionAtRestUsingCustomerKeyManagementApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEncryptionAtRestUsingCustomerKeyManagementApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *EncryptionAtRestUsingCustomerKeyManagementApi {
	mock := &EncryptionAtRestUsingCustomerKeyManagementApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
