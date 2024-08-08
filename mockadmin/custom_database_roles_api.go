// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20240805001/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// CustomDatabaseRolesApi is an autogenerated mock type for the CustomDatabaseRolesApi type
type CustomDatabaseRolesApi struct {
	mock.Mock
}

type CustomDatabaseRolesApi_Expecter struct {
	mock *mock.Mock
}

func (_m *CustomDatabaseRolesApi) EXPECT() *CustomDatabaseRolesApi_Expecter {
	return &CustomDatabaseRolesApi_Expecter{mock: &_m.Mock}
}

// CreateCustomDatabaseRole provides a mock function with given fields: ctx, groupId, userCustomDBRole
func (_m *CustomDatabaseRolesApi) CreateCustomDatabaseRole(ctx context.Context, groupId string, userCustomDBRole *admin.UserCustomDBRole) admin.CreateCustomDatabaseRoleApiRequest {
	ret := _m.Called(ctx, groupId, userCustomDBRole)

	if len(ret) == 0 {
		panic("no return value specified for CreateCustomDatabaseRole")
	}

	var r0 admin.CreateCustomDatabaseRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.UserCustomDBRole) admin.CreateCustomDatabaseRoleApiRequest); ok {
		r0 = rf(ctx, groupId, userCustomDBRole)
	} else {
		r0 = ret.Get(0).(admin.CreateCustomDatabaseRoleApiRequest)
	}

	return r0
}

// CustomDatabaseRolesApi_CreateCustomDatabaseRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCustomDatabaseRole'
type CustomDatabaseRolesApi_CreateCustomDatabaseRole_Call struct {
	*mock.Call
}

// CreateCustomDatabaseRole is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - userCustomDBRole *admin.UserCustomDBRole
func (_e *CustomDatabaseRolesApi_Expecter) CreateCustomDatabaseRole(ctx interface{}, groupId interface{}, userCustomDBRole interface{}) *CustomDatabaseRolesApi_CreateCustomDatabaseRole_Call {
	return &CustomDatabaseRolesApi_CreateCustomDatabaseRole_Call{Call: _e.mock.On("CreateCustomDatabaseRole", ctx, groupId, userCustomDBRole)}
}

func (_c *CustomDatabaseRolesApi_CreateCustomDatabaseRole_Call) Run(run func(ctx context.Context, groupId string, userCustomDBRole *admin.UserCustomDBRole)) *CustomDatabaseRolesApi_CreateCustomDatabaseRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.UserCustomDBRole))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_CreateCustomDatabaseRole_Call) Return(_a0 admin.CreateCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_CreateCustomDatabaseRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CustomDatabaseRolesApi_CreateCustomDatabaseRole_Call) RunAndReturn(run func(context.Context, string, *admin.UserCustomDBRole) admin.CreateCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_CreateCustomDatabaseRole_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCustomDatabaseRoleExecute provides a mock function with given fields: r
func (_m *CustomDatabaseRolesApi) CreateCustomDatabaseRoleExecute(r admin.CreateCustomDatabaseRoleApiRequest) (*admin.UserCustomDBRole, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateCustomDatabaseRoleExecute")
	}

	var r0 *admin.UserCustomDBRole
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateCustomDatabaseRoleApiRequest) (*admin.UserCustomDBRole, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateCustomDatabaseRoleApiRequest) *admin.UserCustomDBRole); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.UserCustomDBRole)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateCustomDatabaseRoleApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateCustomDatabaseRoleApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CustomDatabaseRolesApi_CreateCustomDatabaseRoleExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCustomDatabaseRoleExecute'
type CustomDatabaseRolesApi_CreateCustomDatabaseRoleExecute_Call struct {
	*mock.Call
}

// CreateCustomDatabaseRoleExecute is a helper method to define mock.On call
//   - r admin.CreateCustomDatabaseRoleApiRequest
func (_e *CustomDatabaseRolesApi_Expecter) CreateCustomDatabaseRoleExecute(r interface{}) *CustomDatabaseRolesApi_CreateCustomDatabaseRoleExecute_Call {
	return &CustomDatabaseRolesApi_CreateCustomDatabaseRoleExecute_Call{Call: _e.mock.On("CreateCustomDatabaseRoleExecute", r)}
}

func (_c *CustomDatabaseRolesApi_CreateCustomDatabaseRoleExecute_Call) Run(run func(r admin.CreateCustomDatabaseRoleApiRequest)) *CustomDatabaseRolesApi_CreateCustomDatabaseRoleExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateCustomDatabaseRoleApiRequest))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_CreateCustomDatabaseRoleExecute_Call) Return(_a0 *admin.UserCustomDBRole, _a1 *http.Response, _a2 error) *CustomDatabaseRolesApi_CreateCustomDatabaseRoleExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CustomDatabaseRolesApi_CreateCustomDatabaseRoleExecute_Call) RunAndReturn(run func(admin.CreateCustomDatabaseRoleApiRequest) (*admin.UserCustomDBRole, *http.Response, error)) *CustomDatabaseRolesApi_CreateCustomDatabaseRoleExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCustomDatabaseRoleWithParams provides a mock function with given fields: ctx, args
func (_m *CustomDatabaseRolesApi) CreateCustomDatabaseRoleWithParams(ctx context.Context, args *admin.CreateCustomDatabaseRoleApiParams) admin.CreateCustomDatabaseRoleApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateCustomDatabaseRoleWithParams")
	}

	var r0 admin.CreateCustomDatabaseRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateCustomDatabaseRoleApiParams) admin.CreateCustomDatabaseRoleApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateCustomDatabaseRoleApiRequest)
	}

	return r0
}

// CustomDatabaseRolesApi_CreateCustomDatabaseRoleWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCustomDatabaseRoleWithParams'
type CustomDatabaseRolesApi_CreateCustomDatabaseRoleWithParams_Call struct {
	*mock.Call
}

// CreateCustomDatabaseRoleWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateCustomDatabaseRoleApiParams
func (_e *CustomDatabaseRolesApi_Expecter) CreateCustomDatabaseRoleWithParams(ctx interface{}, args interface{}) *CustomDatabaseRolesApi_CreateCustomDatabaseRoleWithParams_Call {
	return &CustomDatabaseRolesApi_CreateCustomDatabaseRoleWithParams_Call{Call: _e.mock.On("CreateCustomDatabaseRoleWithParams", ctx, args)}
}

func (_c *CustomDatabaseRolesApi_CreateCustomDatabaseRoleWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateCustomDatabaseRoleApiParams)) *CustomDatabaseRolesApi_CreateCustomDatabaseRoleWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateCustomDatabaseRoleApiParams))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_CreateCustomDatabaseRoleWithParams_Call) Return(_a0 admin.CreateCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_CreateCustomDatabaseRoleWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CustomDatabaseRolesApi_CreateCustomDatabaseRoleWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateCustomDatabaseRoleApiParams) admin.CreateCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_CreateCustomDatabaseRoleWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCustomDatabaseRole provides a mock function with given fields: ctx, groupId, roleName
func (_m *CustomDatabaseRolesApi) DeleteCustomDatabaseRole(ctx context.Context, groupId string, roleName string) admin.DeleteCustomDatabaseRoleApiRequest {
	ret := _m.Called(ctx, groupId, roleName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCustomDatabaseRole")
	}

	var r0 admin.DeleteCustomDatabaseRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.DeleteCustomDatabaseRoleApiRequest); ok {
		r0 = rf(ctx, groupId, roleName)
	} else {
		r0 = ret.Get(0).(admin.DeleteCustomDatabaseRoleApiRequest)
	}

	return r0
}

// CustomDatabaseRolesApi_DeleteCustomDatabaseRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCustomDatabaseRole'
type CustomDatabaseRolesApi_DeleteCustomDatabaseRole_Call struct {
	*mock.Call
}

// DeleteCustomDatabaseRole is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - roleName string
func (_e *CustomDatabaseRolesApi_Expecter) DeleteCustomDatabaseRole(ctx interface{}, groupId interface{}, roleName interface{}) *CustomDatabaseRolesApi_DeleteCustomDatabaseRole_Call {
	return &CustomDatabaseRolesApi_DeleteCustomDatabaseRole_Call{Call: _e.mock.On("DeleteCustomDatabaseRole", ctx, groupId, roleName)}
}

func (_c *CustomDatabaseRolesApi_DeleteCustomDatabaseRole_Call) Run(run func(ctx context.Context, groupId string, roleName string)) *CustomDatabaseRolesApi_DeleteCustomDatabaseRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_DeleteCustomDatabaseRole_Call) Return(_a0 admin.DeleteCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_DeleteCustomDatabaseRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CustomDatabaseRolesApi_DeleteCustomDatabaseRole_Call) RunAndReturn(run func(context.Context, string, string) admin.DeleteCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_DeleteCustomDatabaseRole_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCustomDatabaseRoleExecute provides a mock function with given fields: r
func (_m *CustomDatabaseRolesApi) DeleteCustomDatabaseRoleExecute(r admin.DeleteCustomDatabaseRoleApiRequest) (*http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCustomDatabaseRoleExecute")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(admin.DeleteCustomDatabaseRoleApiRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DeleteCustomDatabaseRoleApiRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DeleteCustomDatabaseRoleApiRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CustomDatabaseRolesApi_DeleteCustomDatabaseRoleExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCustomDatabaseRoleExecute'
type CustomDatabaseRolesApi_DeleteCustomDatabaseRoleExecute_Call struct {
	*mock.Call
}

// DeleteCustomDatabaseRoleExecute is a helper method to define mock.On call
//   - r admin.DeleteCustomDatabaseRoleApiRequest
func (_e *CustomDatabaseRolesApi_Expecter) DeleteCustomDatabaseRoleExecute(r interface{}) *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleExecute_Call {
	return &CustomDatabaseRolesApi_DeleteCustomDatabaseRoleExecute_Call{Call: _e.mock.On("DeleteCustomDatabaseRoleExecute", r)}
}

func (_c *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleExecute_Call) Run(run func(r admin.DeleteCustomDatabaseRoleApiRequest)) *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DeleteCustomDatabaseRoleApiRequest))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleExecute_Call) Return(_a0 *http.Response, _a1 error) *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleExecute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleExecute_Call) RunAndReturn(run func(admin.DeleteCustomDatabaseRoleApiRequest) (*http.Response, error)) *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCustomDatabaseRoleWithParams provides a mock function with given fields: ctx, args
func (_m *CustomDatabaseRolesApi) DeleteCustomDatabaseRoleWithParams(ctx context.Context, args *admin.DeleteCustomDatabaseRoleApiParams) admin.DeleteCustomDatabaseRoleApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCustomDatabaseRoleWithParams")
	}

	var r0 admin.DeleteCustomDatabaseRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeleteCustomDatabaseRoleApiParams) admin.DeleteCustomDatabaseRoleApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DeleteCustomDatabaseRoleApiRequest)
	}

	return r0
}

// CustomDatabaseRolesApi_DeleteCustomDatabaseRoleWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCustomDatabaseRoleWithParams'
type CustomDatabaseRolesApi_DeleteCustomDatabaseRoleWithParams_Call struct {
	*mock.Call
}

// DeleteCustomDatabaseRoleWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DeleteCustomDatabaseRoleApiParams
func (_e *CustomDatabaseRolesApi_Expecter) DeleteCustomDatabaseRoleWithParams(ctx interface{}, args interface{}) *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleWithParams_Call {
	return &CustomDatabaseRolesApi_DeleteCustomDatabaseRoleWithParams_Call{Call: _e.mock.On("DeleteCustomDatabaseRoleWithParams", ctx, args)}
}

func (_c *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleWithParams_Call) Run(run func(ctx context.Context, args *admin.DeleteCustomDatabaseRoleApiParams)) *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DeleteCustomDatabaseRoleApiParams))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleWithParams_Call) Return(_a0 admin.DeleteCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleWithParams_Call) RunAndReturn(run func(context.Context, *admin.DeleteCustomDatabaseRoleApiParams) admin.DeleteCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_DeleteCustomDatabaseRoleWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetCustomDatabaseRole provides a mock function with given fields: ctx, groupId, roleName
func (_m *CustomDatabaseRolesApi) GetCustomDatabaseRole(ctx context.Context, groupId string, roleName string) admin.GetCustomDatabaseRoleApiRequest {
	ret := _m.Called(ctx, groupId, roleName)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomDatabaseRole")
	}

	var r0 admin.GetCustomDatabaseRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetCustomDatabaseRoleApiRequest); ok {
		r0 = rf(ctx, groupId, roleName)
	} else {
		r0 = ret.Get(0).(admin.GetCustomDatabaseRoleApiRequest)
	}

	return r0
}

// CustomDatabaseRolesApi_GetCustomDatabaseRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCustomDatabaseRole'
type CustomDatabaseRolesApi_GetCustomDatabaseRole_Call struct {
	*mock.Call
}

// GetCustomDatabaseRole is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - roleName string
func (_e *CustomDatabaseRolesApi_Expecter) GetCustomDatabaseRole(ctx interface{}, groupId interface{}, roleName interface{}) *CustomDatabaseRolesApi_GetCustomDatabaseRole_Call {
	return &CustomDatabaseRolesApi_GetCustomDatabaseRole_Call{Call: _e.mock.On("GetCustomDatabaseRole", ctx, groupId, roleName)}
}

func (_c *CustomDatabaseRolesApi_GetCustomDatabaseRole_Call) Run(run func(ctx context.Context, groupId string, roleName string)) *CustomDatabaseRolesApi_GetCustomDatabaseRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_GetCustomDatabaseRole_Call) Return(_a0 admin.GetCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_GetCustomDatabaseRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CustomDatabaseRolesApi_GetCustomDatabaseRole_Call) RunAndReturn(run func(context.Context, string, string) admin.GetCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_GetCustomDatabaseRole_Call {
	_c.Call.Return(run)
	return _c
}

// GetCustomDatabaseRoleExecute provides a mock function with given fields: r
func (_m *CustomDatabaseRolesApi) GetCustomDatabaseRoleExecute(r admin.GetCustomDatabaseRoleApiRequest) (*admin.UserCustomDBRole, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomDatabaseRoleExecute")
	}

	var r0 *admin.UserCustomDBRole
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetCustomDatabaseRoleApiRequest) (*admin.UserCustomDBRole, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetCustomDatabaseRoleApiRequest) *admin.UserCustomDBRole); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.UserCustomDBRole)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetCustomDatabaseRoleApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetCustomDatabaseRoleApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CustomDatabaseRolesApi_GetCustomDatabaseRoleExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCustomDatabaseRoleExecute'
type CustomDatabaseRolesApi_GetCustomDatabaseRoleExecute_Call struct {
	*mock.Call
}

// GetCustomDatabaseRoleExecute is a helper method to define mock.On call
//   - r admin.GetCustomDatabaseRoleApiRequest
func (_e *CustomDatabaseRolesApi_Expecter) GetCustomDatabaseRoleExecute(r interface{}) *CustomDatabaseRolesApi_GetCustomDatabaseRoleExecute_Call {
	return &CustomDatabaseRolesApi_GetCustomDatabaseRoleExecute_Call{Call: _e.mock.On("GetCustomDatabaseRoleExecute", r)}
}

func (_c *CustomDatabaseRolesApi_GetCustomDatabaseRoleExecute_Call) Run(run func(r admin.GetCustomDatabaseRoleApiRequest)) *CustomDatabaseRolesApi_GetCustomDatabaseRoleExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetCustomDatabaseRoleApiRequest))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_GetCustomDatabaseRoleExecute_Call) Return(_a0 *admin.UserCustomDBRole, _a1 *http.Response, _a2 error) *CustomDatabaseRolesApi_GetCustomDatabaseRoleExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CustomDatabaseRolesApi_GetCustomDatabaseRoleExecute_Call) RunAndReturn(run func(admin.GetCustomDatabaseRoleApiRequest) (*admin.UserCustomDBRole, *http.Response, error)) *CustomDatabaseRolesApi_GetCustomDatabaseRoleExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetCustomDatabaseRoleWithParams provides a mock function with given fields: ctx, args
func (_m *CustomDatabaseRolesApi) GetCustomDatabaseRoleWithParams(ctx context.Context, args *admin.GetCustomDatabaseRoleApiParams) admin.GetCustomDatabaseRoleApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomDatabaseRoleWithParams")
	}

	var r0 admin.GetCustomDatabaseRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetCustomDatabaseRoleApiParams) admin.GetCustomDatabaseRoleApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetCustomDatabaseRoleApiRequest)
	}

	return r0
}

// CustomDatabaseRolesApi_GetCustomDatabaseRoleWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCustomDatabaseRoleWithParams'
type CustomDatabaseRolesApi_GetCustomDatabaseRoleWithParams_Call struct {
	*mock.Call
}

// GetCustomDatabaseRoleWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetCustomDatabaseRoleApiParams
func (_e *CustomDatabaseRolesApi_Expecter) GetCustomDatabaseRoleWithParams(ctx interface{}, args interface{}) *CustomDatabaseRolesApi_GetCustomDatabaseRoleWithParams_Call {
	return &CustomDatabaseRolesApi_GetCustomDatabaseRoleWithParams_Call{Call: _e.mock.On("GetCustomDatabaseRoleWithParams", ctx, args)}
}

func (_c *CustomDatabaseRolesApi_GetCustomDatabaseRoleWithParams_Call) Run(run func(ctx context.Context, args *admin.GetCustomDatabaseRoleApiParams)) *CustomDatabaseRolesApi_GetCustomDatabaseRoleWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetCustomDatabaseRoleApiParams))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_GetCustomDatabaseRoleWithParams_Call) Return(_a0 admin.GetCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_GetCustomDatabaseRoleWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CustomDatabaseRolesApi_GetCustomDatabaseRoleWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetCustomDatabaseRoleApiParams) admin.GetCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_GetCustomDatabaseRoleWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListCustomDatabaseRoles provides a mock function with given fields: ctx, groupId
func (_m *CustomDatabaseRolesApi) ListCustomDatabaseRoles(ctx context.Context, groupId string) admin.ListCustomDatabaseRolesApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for ListCustomDatabaseRoles")
	}

	var r0 admin.ListCustomDatabaseRolesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ListCustomDatabaseRolesApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.ListCustomDatabaseRolesApiRequest)
	}

	return r0
}

// CustomDatabaseRolesApi_ListCustomDatabaseRoles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCustomDatabaseRoles'
type CustomDatabaseRolesApi_ListCustomDatabaseRoles_Call struct {
	*mock.Call
}

// ListCustomDatabaseRoles is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *CustomDatabaseRolesApi_Expecter) ListCustomDatabaseRoles(ctx interface{}, groupId interface{}) *CustomDatabaseRolesApi_ListCustomDatabaseRoles_Call {
	return &CustomDatabaseRolesApi_ListCustomDatabaseRoles_Call{Call: _e.mock.On("ListCustomDatabaseRoles", ctx, groupId)}
}

func (_c *CustomDatabaseRolesApi_ListCustomDatabaseRoles_Call) Run(run func(ctx context.Context, groupId string)) *CustomDatabaseRolesApi_ListCustomDatabaseRoles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_ListCustomDatabaseRoles_Call) Return(_a0 admin.ListCustomDatabaseRolesApiRequest) *CustomDatabaseRolesApi_ListCustomDatabaseRoles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CustomDatabaseRolesApi_ListCustomDatabaseRoles_Call) RunAndReturn(run func(context.Context, string) admin.ListCustomDatabaseRolesApiRequest) *CustomDatabaseRolesApi_ListCustomDatabaseRoles_Call {
	_c.Call.Return(run)
	return _c
}

// ListCustomDatabaseRolesExecute provides a mock function with given fields: r
func (_m *CustomDatabaseRolesApi) ListCustomDatabaseRolesExecute(r admin.ListCustomDatabaseRolesApiRequest) ([]admin.UserCustomDBRole, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListCustomDatabaseRolesExecute")
	}

	var r0 []admin.UserCustomDBRole
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListCustomDatabaseRolesApiRequest) ([]admin.UserCustomDBRole, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListCustomDatabaseRolesApiRequest) []admin.UserCustomDBRole); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]admin.UserCustomDBRole)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListCustomDatabaseRolesApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListCustomDatabaseRolesApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CustomDatabaseRolesApi_ListCustomDatabaseRolesExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCustomDatabaseRolesExecute'
type CustomDatabaseRolesApi_ListCustomDatabaseRolesExecute_Call struct {
	*mock.Call
}

// ListCustomDatabaseRolesExecute is a helper method to define mock.On call
//   - r admin.ListCustomDatabaseRolesApiRequest
func (_e *CustomDatabaseRolesApi_Expecter) ListCustomDatabaseRolesExecute(r interface{}) *CustomDatabaseRolesApi_ListCustomDatabaseRolesExecute_Call {
	return &CustomDatabaseRolesApi_ListCustomDatabaseRolesExecute_Call{Call: _e.mock.On("ListCustomDatabaseRolesExecute", r)}
}

func (_c *CustomDatabaseRolesApi_ListCustomDatabaseRolesExecute_Call) Run(run func(r admin.ListCustomDatabaseRolesApiRequest)) *CustomDatabaseRolesApi_ListCustomDatabaseRolesExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListCustomDatabaseRolesApiRequest))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_ListCustomDatabaseRolesExecute_Call) Return(_a0 []admin.UserCustomDBRole, _a1 *http.Response, _a2 error) *CustomDatabaseRolesApi_ListCustomDatabaseRolesExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CustomDatabaseRolesApi_ListCustomDatabaseRolesExecute_Call) RunAndReturn(run func(admin.ListCustomDatabaseRolesApiRequest) ([]admin.UserCustomDBRole, *http.Response, error)) *CustomDatabaseRolesApi_ListCustomDatabaseRolesExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListCustomDatabaseRolesWithParams provides a mock function with given fields: ctx, args
func (_m *CustomDatabaseRolesApi) ListCustomDatabaseRolesWithParams(ctx context.Context, args *admin.ListCustomDatabaseRolesApiParams) admin.ListCustomDatabaseRolesApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListCustomDatabaseRolesWithParams")
	}

	var r0 admin.ListCustomDatabaseRolesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListCustomDatabaseRolesApiParams) admin.ListCustomDatabaseRolesApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListCustomDatabaseRolesApiRequest)
	}

	return r0
}

// CustomDatabaseRolesApi_ListCustomDatabaseRolesWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCustomDatabaseRolesWithParams'
type CustomDatabaseRolesApi_ListCustomDatabaseRolesWithParams_Call struct {
	*mock.Call
}

// ListCustomDatabaseRolesWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListCustomDatabaseRolesApiParams
func (_e *CustomDatabaseRolesApi_Expecter) ListCustomDatabaseRolesWithParams(ctx interface{}, args interface{}) *CustomDatabaseRolesApi_ListCustomDatabaseRolesWithParams_Call {
	return &CustomDatabaseRolesApi_ListCustomDatabaseRolesWithParams_Call{Call: _e.mock.On("ListCustomDatabaseRolesWithParams", ctx, args)}
}

func (_c *CustomDatabaseRolesApi_ListCustomDatabaseRolesWithParams_Call) Run(run func(ctx context.Context, args *admin.ListCustomDatabaseRolesApiParams)) *CustomDatabaseRolesApi_ListCustomDatabaseRolesWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListCustomDatabaseRolesApiParams))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_ListCustomDatabaseRolesWithParams_Call) Return(_a0 admin.ListCustomDatabaseRolesApiRequest) *CustomDatabaseRolesApi_ListCustomDatabaseRolesWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CustomDatabaseRolesApi_ListCustomDatabaseRolesWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListCustomDatabaseRolesApiParams) admin.ListCustomDatabaseRolesApiRequest) *CustomDatabaseRolesApi_ListCustomDatabaseRolesWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCustomDatabaseRole provides a mock function with given fields: ctx, groupId, roleName, updateCustomDBRole
func (_m *CustomDatabaseRolesApi) UpdateCustomDatabaseRole(ctx context.Context, groupId string, roleName string, updateCustomDBRole *admin.UpdateCustomDBRole) admin.UpdateCustomDatabaseRoleApiRequest {
	ret := _m.Called(ctx, groupId, roleName, updateCustomDBRole)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCustomDatabaseRole")
	}

	var r0 admin.UpdateCustomDatabaseRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.UpdateCustomDBRole) admin.UpdateCustomDatabaseRoleApiRequest); ok {
		r0 = rf(ctx, groupId, roleName, updateCustomDBRole)
	} else {
		r0 = ret.Get(0).(admin.UpdateCustomDatabaseRoleApiRequest)
	}

	return r0
}

// CustomDatabaseRolesApi_UpdateCustomDatabaseRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCustomDatabaseRole'
type CustomDatabaseRolesApi_UpdateCustomDatabaseRole_Call struct {
	*mock.Call
}

// UpdateCustomDatabaseRole is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - roleName string
//   - updateCustomDBRole *admin.UpdateCustomDBRole
func (_e *CustomDatabaseRolesApi_Expecter) UpdateCustomDatabaseRole(ctx interface{}, groupId interface{}, roleName interface{}, updateCustomDBRole interface{}) *CustomDatabaseRolesApi_UpdateCustomDatabaseRole_Call {
	return &CustomDatabaseRolesApi_UpdateCustomDatabaseRole_Call{Call: _e.mock.On("UpdateCustomDatabaseRole", ctx, groupId, roleName, updateCustomDBRole)}
}

func (_c *CustomDatabaseRolesApi_UpdateCustomDatabaseRole_Call) Run(run func(ctx context.Context, groupId string, roleName string, updateCustomDBRole *admin.UpdateCustomDBRole)) *CustomDatabaseRolesApi_UpdateCustomDatabaseRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.UpdateCustomDBRole))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_UpdateCustomDatabaseRole_Call) Return(_a0 admin.UpdateCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_UpdateCustomDatabaseRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CustomDatabaseRolesApi_UpdateCustomDatabaseRole_Call) RunAndReturn(run func(context.Context, string, string, *admin.UpdateCustomDBRole) admin.UpdateCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_UpdateCustomDatabaseRole_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCustomDatabaseRoleExecute provides a mock function with given fields: r
func (_m *CustomDatabaseRolesApi) UpdateCustomDatabaseRoleExecute(r admin.UpdateCustomDatabaseRoleApiRequest) (*admin.UserCustomDBRole, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCustomDatabaseRoleExecute")
	}

	var r0 *admin.UserCustomDBRole
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.UpdateCustomDatabaseRoleApiRequest) (*admin.UserCustomDBRole, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.UpdateCustomDatabaseRoleApiRequest) *admin.UserCustomDBRole); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.UserCustomDBRole)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.UpdateCustomDatabaseRoleApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.UpdateCustomDatabaseRoleApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CustomDatabaseRolesApi_UpdateCustomDatabaseRoleExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCustomDatabaseRoleExecute'
type CustomDatabaseRolesApi_UpdateCustomDatabaseRoleExecute_Call struct {
	*mock.Call
}

// UpdateCustomDatabaseRoleExecute is a helper method to define mock.On call
//   - r admin.UpdateCustomDatabaseRoleApiRequest
func (_e *CustomDatabaseRolesApi_Expecter) UpdateCustomDatabaseRoleExecute(r interface{}) *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleExecute_Call {
	return &CustomDatabaseRolesApi_UpdateCustomDatabaseRoleExecute_Call{Call: _e.mock.On("UpdateCustomDatabaseRoleExecute", r)}
}

func (_c *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleExecute_Call) Run(run func(r admin.UpdateCustomDatabaseRoleApiRequest)) *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.UpdateCustomDatabaseRoleApiRequest))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleExecute_Call) Return(_a0 *admin.UserCustomDBRole, _a1 *http.Response, _a2 error) *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleExecute_Call) RunAndReturn(run func(admin.UpdateCustomDatabaseRoleApiRequest) (*admin.UserCustomDBRole, *http.Response, error)) *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleExecute_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCustomDatabaseRoleWithParams provides a mock function with given fields: ctx, args
func (_m *CustomDatabaseRolesApi) UpdateCustomDatabaseRoleWithParams(ctx context.Context, args *admin.UpdateCustomDatabaseRoleApiParams) admin.UpdateCustomDatabaseRoleApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCustomDatabaseRoleWithParams")
	}

	var r0 admin.UpdateCustomDatabaseRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.UpdateCustomDatabaseRoleApiParams) admin.UpdateCustomDatabaseRoleApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.UpdateCustomDatabaseRoleApiRequest)
	}

	return r0
}

// CustomDatabaseRolesApi_UpdateCustomDatabaseRoleWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCustomDatabaseRoleWithParams'
type CustomDatabaseRolesApi_UpdateCustomDatabaseRoleWithParams_Call struct {
	*mock.Call
}

// UpdateCustomDatabaseRoleWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.UpdateCustomDatabaseRoleApiParams
func (_e *CustomDatabaseRolesApi_Expecter) UpdateCustomDatabaseRoleWithParams(ctx interface{}, args interface{}) *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleWithParams_Call {
	return &CustomDatabaseRolesApi_UpdateCustomDatabaseRoleWithParams_Call{Call: _e.mock.On("UpdateCustomDatabaseRoleWithParams", ctx, args)}
}

func (_c *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleWithParams_Call) Run(run func(ctx context.Context, args *admin.UpdateCustomDatabaseRoleApiParams)) *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.UpdateCustomDatabaseRoleApiParams))
	})
	return _c
}

func (_c *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleWithParams_Call) Return(_a0 admin.UpdateCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleWithParams_Call) RunAndReturn(run func(context.Context, *admin.UpdateCustomDatabaseRoleApiParams) admin.UpdateCustomDatabaseRoleApiRequest) *CustomDatabaseRolesApi_UpdateCustomDatabaseRoleWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewCustomDatabaseRolesApi creates a new instance of CustomDatabaseRolesApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCustomDatabaseRolesApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *CustomDatabaseRolesApi {
	mock := &CustomDatabaseRolesApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
