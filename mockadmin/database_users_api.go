// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20241023002/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// DatabaseUsersApi is an autogenerated mock type for the DatabaseUsersApi type
type DatabaseUsersApi struct {
	mock.Mock
}

type DatabaseUsersApi_Expecter struct {
	mock *mock.Mock
}

func (_m *DatabaseUsersApi) EXPECT() *DatabaseUsersApi_Expecter {
	return &DatabaseUsersApi_Expecter{mock: &_m.Mock}
}

// CreateDatabaseUser provides a mock function with given fields: ctx, groupId, cloudDatabaseUser
func (_m *DatabaseUsersApi) CreateDatabaseUser(ctx context.Context, groupId string, cloudDatabaseUser *admin.CloudDatabaseUser) admin.CreateDatabaseUserApiRequest {
	ret := _m.Called(ctx, groupId, cloudDatabaseUser)

	if len(ret) == 0 {
		panic("no return value specified for CreateDatabaseUser")
	}

	var r0 admin.CreateDatabaseUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.CloudDatabaseUser) admin.CreateDatabaseUserApiRequest); ok {
		r0 = rf(ctx, groupId, cloudDatabaseUser)
	} else {
		r0 = ret.Get(0).(admin.CreateDatabaseUserApiRequest)
	}

	return r0
}

// DatabaseUsersApi_CreateDatabaseUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDatabaseUser'
type DatabaseUsersApi_CreateDatabaseUser_Call struct {
	*mock.Call
}

// CreateDatabaseUser is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - cloudDatabaseUser *admin.CloudDatabaseUser
func (_e *DatabaseUsersApi_Expecter) CreateDatabaseUser(ctx any, groupId any, cloudDatabaseUser any) *DatabaseUsersApi_CreateDatabaseUser_Call {
	return &DatabaseUsersApi_CreateDatabaseUser_Call{Call: _e.mock.On("CreateDatabaseUser", ctx, groupId, cloudDatabaseUser)}
}

func (_c *DatabaseUsersApi_CreateDatabaseUser_Call) Run(run func(ctx context.Context, groupId string, cloudDatabaseUser *admin.CloudDatabaseUser)) *DatabaseUsersApi_CreateDatabaseUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.CloudDatabaseUser))
	})
	return _c
}

func (_c *DatabaseUsersApi_CreateDatabaseUser_Call) Return(_a0 admin.CreateDatabaseUserApiRequest) *DatabaseUsersApi_CreateDatabaseUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DatabaseUsersApi_CreateDatabaseUser_Call) RunAndReturn(run func(context.Context, string, *admin.CloudDatabaseUser) admin.CreateDatabaseUserApiRequest) *DatabaseUsersApi_CreateDatabaseUser_Call {
	_c.Call.Return(run)
	return _c
}

// CreateDatabaseUserExecute provides a mock function with given fields: r
func (_m *DatabaseUsersApi) CreateDatabaseUserExecute(r admin.CreateDatabaseUserApiRequest) (*admin.CloudDatabaseUser, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateDatabaseUserExecute")
	}

	var r0 *admin.CloudDatabaseUser
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateDatabaseUserApiRequest) (*admin.CloudDatabaseUser, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateDatabaseUserApiRequest) *admin.CloudDatabaseUser); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CloudDatabaseUser)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateDatabaseUserApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateDatabaseUserApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DatabaseUsersApi_CreateDatabaseUserExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDatabaseUserExecute'
type DatabaseUsersApi_CreateDatabaseUserExecute_Call struct {
	*mock.Call
}

// CreateDatabaseUserExecute is a helper method to define mock.On call
//   - r admin.CreateDatabaseUserApiRequest
func (_e *DatabaseUsersApi_Expecter) CreateDatabaseUserExecute(r any) *DatabaseUsersApi_CreateDatabaseUserExecute_Call {
	return &DatabaseUsersApi_CreateDatabaseUserExecute_Call{Call: _e.mock.On("CreateDatabaseUserExecute", r)}
}

func (_c *DatabaseUsersApi_CreateDatabaseUserExecute_Call) Run(run func(r admin.CreateDatabaseUserApiRequest)) *DatabaseUsersApi_CreateDatabaseUserExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateDatabaseUserApiRequest))
	})
	return _c
}

func (_c *DatabaseUsersApi_CreateDatabaseUserExecute_Call) Return(_a0 *admin.CloudDatabaseUser, _a1 *http.Response, _a2 error) *DatabaseUsersApi_CreateDatabaseUserExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *DatabaseUsersApi_CreateDatabaseUserExecute_Call) RunAndReturn(run func(admin.CreateDatabaseUserApiRequest) (*admin.CloudDatabaseUser, *http.Response, error)) *DatabaseUsersApi_CreateDatabaseUserExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateDatabaseUserWithParams provides a mock function with given fields: ctx, args
func (_m *DatabaseUsersApi) CreateDatabaseUserWithParams(ctx context.Context, args *admin.CreateDatabaseUserApiParams) admin.CreateDatabaseUserApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateDatabaseUserWithParams")
	}

	var r0 admin.CreateDatabaseUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateDatabaseUserApiParams) admin.CreateDatabaseUserApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateDatabaseUserApiRequest)
	}

	return r0
}

// DatabaseUsersApi_CreateDatabaseUserWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDatabaseUserWithParams'
type DatabaseUsersApi_CreateDatabaseUserWithParams_Call struct {
	*mock.Call
}

// CreateDatabaseUserWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateDatabaseUserApiParams
func (_e *DatabaseUsersApi_Expecter) CreateDatabaseUserWithParams(ctx any, args any) *DatabaseUsersApi_CreateDatabaseUserWithParams_Call {
	return &DatabaseUsersApi_CreateDatabaseUserWithParams_Call{Call: _e.mock.On("CreateDatabaseUserWithParams", ctx, args)}
}

func (_c *DatabaseUsersApi_CreateDatabaseUserWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateDatabaseUserApiParams)) *DatabaseUsersApi_CreateDatabaseUserWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateDatabaseUserApiParams))
	})
	return _c
}

func (_c *DatabaseUsersApi_CreateDatabaseUserWithParams_Call) Return(_a0 admin.CreateDatabaseUserApiRequest) *DatabaseUsersApi_CreateDatabaseUserWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DatabaseUsersApi_CreateDatabaseUserWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateDatabaseUserApiParams) admin.CreateDatabaseUserApiRequest) *DatabaseUsersApi_CreateDatabaseUserWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteDatabaseUser provides a mock function with given fields: ctx, groupId, databaseName, username
func (_m *DatabaseUsersApi) DeleteDatabaseUser(ctx context.Context, groupId string, databaseName string, username string) admin.DeleteDatabaseUserApiRequest {
	ret := _m.Called(ctx, groupId, databaseName, username)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDatabaseUser")
	}

	var r0 admin.DeleteDatabaseUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) admin.DeleteDatabaseUserApiRequest); ok {
		r0 = rf(ctx, groupId, databaseName, username)
	} else {
		r0 = ret.Get(0).(admin.DeleteDatabaseUserApiRequest)
	}

	return r0
}

// DatabaseUsersApi_DeleteDatabaseUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteDatabaseUser'
type DatabaseUsersApi_DeleteDatabaseUser_Call struct {
	*mock.Call
}

// DeleteDatabaseUser is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - databaseName string
//   - username string
func (_e *DatabaseUsersApi_Expecter) DeleteDatabaseUser(ctx any, groupId any, databaseName any, username any) *DatabaseUsersApi_DeleteDatabaseUser_Call {
	return &DatabaseUsersApi_DeleteDatabaseUser_Call{Call: _e.mock.On("DeleteDatabaseUser", ctx, groupId, databaseName, username)}
}

func (_c *DatabaseUsersApi_DeleteDatabaseUser_Call) Run(run func(ctx context.Context, groupId string, databaseName string, username string)) *DatabaseUsersApi_DeleteDatabaseUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *DatabaseUsersApi_DeleteDatabaseUser_Call) Return(_a0 admin.DeleteDatabaseUserApiRequest) *DatabaseUsersApi_DeleteDatabaseUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DatabaseUsersApi_DeleteDatabaseUser_Call) RunAndReturn(run func(context.Context, string, string, string) admin.DeleteDatabaseUserApiRequest) *DatabaseUsersApi_DeleteDatabaseUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteDatabaseUserExecute provides a mock function with given fields: r
func (_m *DatabaseUsersApi) DeleteDatabaseUserExecute(r admin.DeleteDatabaseUserApiRequest) (any, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDatabaseUserExecute")
	}

	var r0 any
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.DeleteDatabaseUserApiRequest) (any, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DeleteDatabaseUserApiRequest) any); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(any)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DeleteDatabaseUserApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.DeleteDatabaseUserApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DatabaseUsersApi_DeleteDatabaseUserExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteDatabaseUserExecute'
type DatabaseUsersApi_DeleteDatabaseUserExecute_Call struct {
	*mock.Call
}

// DeleteDatabaseUserExecute is a helper method to define mock.On call
//   - r admin.DeleteDatabaseUserApiRequest
func (_e *DatabaseUsersApi_Expecter) DeleteDatabaseUserExecute(r any) *DatabaseUsersApi_DeleteDatabaseUserExecute_Call {
	return &DatabaseUsersApi_DeleteDatabaseUserExecute_Call{Call: _e.mock.On("DeleteDatabaseUserExecute", r)}
}

func (_c *DatabaseUsersApi_DeleteDatabaseUserExecute_Call) Run(run func(r admin.DeleteDatabaseUserApiRequest)) *DatabaseUsersApi_DeleteDatabaseUserExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DeleteDatabaseUserApiRequest))
	})
	return _c
}

func (_c *DatabaseUsersApi_DeleteDatabaseUserExecute_Call) Return(_a0 any, _a1 *http.Response, _a2 error) *DatabaseUsersApi_DeleteDatabaseUserExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *DatabaseUsersApi_DeleteDatabaseUserExecute_Call) RunAndReturn(run func(admin.DeleteDatabaseUserApiRequest) (any, *http.Response, error)) *DatabaseUsersApi_DeleteDatabaseUserExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteDatabaseUserWithParams provides a mock function with given fields: ctx, args
func (_m *DatabaseUsersApi) DeleteDatabaseUserWithParams(ctx context.Context, args *admin.DeleteDatabaseUserApiParams) admin.DeleteDatabaseUserApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDatabaseUserWithParams")
	}

	var r0 admin.DeleteDatabaseUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeleteDatabaseUserApiParams) admin.DeleteDatabaseUserApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DeleteDatabaseUserApiRequest)
	}

	return r0
}

// DatabaseUsersApi_DeleteDatabaseUserWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteDatabaseUserWithParams'
type DatabaseUsersApi_DeleteDatabaseUserWithParams_Call struct {
	*mock.Call
}

// DeleteDatabaseUserWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DeleteDatabaseUserApiParams
func (_e *DatabaseUsersApi_Expecter) DeleteDatabaseUserWithParams(ctx any, args any) *DatabaseUsersApi_DeleteDatabaseUserWithParams_Call {
	return &DatabaseUsersApi_DeleteDatabaseUserWithParams_Call{Call: _e.mock.On("DeleteDatabaseUserWithParams", ctx, args)}
}

func (_c *DatabaseUsersApi_DeleteDatabaseUserWithParams_Call) Run(run func(ctx context.Context, args *admin.DeleteDatabaseUserApiParams)) *DatabaseUsersApi_DeleteDatabaseUserWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DeleteDatabaseUserApiParams))
	})
	return _c
}

func (_c *DatabaseUsersApi_DeleteDatabaseUserWithParams_Call) Return(_a0 admin.DeleteDatabaseUserApiRequest) *DatabaseUsersApi_DeleteDatabaseUserWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DatabaseUsersApi_DeleteDatabaseUserWithParams_Call) RunAndReturn(run func(context.Context, *admin.DeleteDatabaseUserApiParams) admin.DeleteDatabaseUserApiRequest) *DatabaseUsersApi_DeleteDatabaseUserWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetDatabaseUser provides a mock function with given fields: ctx, groupId, databaseName, username
func (_m *DatabaseUsersApi) GetDatabaseUser(ctx context.Context, groupId string, databaseName string, username string) admin.GetDatabaseUserApiRequest {
	ret := _m.Called(ctx, groupId, databaseName, username)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabaseUser")
	}

	var r0 admin.GetDatabaseUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) admin.GetDatabaseUserApiRequest); ok {
		r0 = rf(ctx, groupId, databaseName, username)
	} else {
		r0 = ret.Get(0).(admin.GetDatabaseUserApiRequest)
	}

	return r0
}

// DatabaseUsersApi_GetDatabaseUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDatabaseUser'
type DatabaseUsersApi_GetDatabaseUser_Call struct {
	*mock.Call
}

// GetDatabaseUser is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - databaseName string
//   - username string
func (_e *DatabaseUsersApi_Expecter) GetDatabaseUser(ctx any, groupId any, databaseName any, username any) *DatabaseUsersApi_GetDatabaseUser_Call {
	return &DatabaseUsersApi_GetDatabaseUser_Call{Call: _e.mock.On("GetDatabaseUser", ctx, groupId, databaseName, username)}
}

func (_c *DatabaseUsersApi_GetDatabaseUser_Call) Run(run func(ctx context.Context, groupId string, databaseName string, username string)) *DatabaseUsersApi_GetDatabaseUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *DatabaseUsersApi_GetDatabaseUser_Call) Return(_a0 admin.GetDatabaseUserApiRequest) *DatabaseUsersApi_GetDatabaseUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DatabaseUsersApi_GetDatabaseUser_Call) RunAndReturn(run func(context.Context, string, string, string) admin.GetDatabaseUserApiRequest) *DatabaseUsersApi_GetDatabaseUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetDatabaseUserExecute provides a mock function with given fields: r
func (_m *DatabaseUsersApi) GetDatabaseUserExecute(r admin.GetDatabaseUserApiRequest) (*admin.CloudDatabaseUser, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabaseUserExecute")
	}

	var r0 *admin.CloudDatabaseUser
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetDatabaseUserApiRequest) (*admin.CloudDatabaseUser, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetDatabaseUserApiRequest) *admin.CloudDatabaseUser); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CloudDatabaseUser)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetDatabaseUserApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetDatabaseUserApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DatabaseUsersApi_GetDatabaseUserExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDatabaseUserExecute'
type DatabaseUsersApi_GetDatabaseUserExecute_Call struct {
	*mock.Call
}

// GetDatabaseUserExecute is a helper method to define mock.On call
//   - r admin.GetDatabaseUserApiRequest
func (_e *DatabaseUsersApi_Expecter) GetDatabaseUserExecute(r any) *DatabaseUsersApi_GetDatabaseUserExecute_Call {
	return &DatabaseUsersApi_GetDatabaseUserExecute_Call{Call: _e.mock.On("GetDatabaseUserExecute", r)}
}

func (_c *DatabaseUsersApi_GetDatabaseUserExecute_Call) Run(run func(r admin.GetDatabaseUserApiRequest)) *DatabaseUsersApi_GetDatabaseUserExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetDatabaseUserApiRequest))
	})
	return _c
}

func (_c *DatabaseUsersApi_GetDatabaseUserExecute_Call) Return(_a0 *admin.CloudDatabaseUser, _a1 *http.Response, _a2 error) *DatabaseUsersApi_GetDatabaseUserExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *DatabaseUsersApi_GetDatabaseUserExecute_Call) RunAndReturn(run func(admin.GetDatabaseUserApiRequest) (*admin.CloudDatabaseUser, *http.Response, error)) *DatabaseUsersApi_GetDatabaseUserExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetDatabaseUserWithParams provides a mock function with given fields: ctx, args
func (_m *DatabaseUsersApi) GetDatabaseUserWithParams(ctx context.Context, args *admin.GetDatabaseUserApiParams) admin.GetDatabaseUserApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabaseUserWithParams")
	}

	var r0 admin.GetDatabaseUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetDatabaseUserApiParams) admin.GetDatabaseUserApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetDatabaseUserApiRequest)
	}

	return r0
}

// DatabaseUsersApi_GetDatabaseUserWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDatabaseUserWithParams'
type DatabaseUsersApi_GetDatabaseUserWithParams_Call struct {
	*mock.Call
}

// GetDatabaseUserWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetDatabaseUserApiParams
func (_e *DatabaseUsersApi_Expecter) GetDatabaseUserWithParams(ctx any, args any) *DatabaseUsersApi_GetDatabaseUserWithParams_Call {
	return &DatabaseUsersApi_GetDatabaseUserWithParams_Call{Call: _e.mock.On("GetDatabaseUserWithParams", ctx, args)}
}

func (_c *DatabaseUsersApi_GetDatabaseUserWithParams_Call) Run(run func(ctx context.Context, args *admin.GetDatabaseUserApiParams)) *DatabaseUsersApi_GetDatabaseUserWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetDatabaseUserApiParams))
	})
	return _c
}

func (_c *DatabaseUsersApi_GetDatabaseUserWithParams_Call) Return(_a0 admin.GetDatabaseUserApiRequest) *DatabaseUsersApi_GetDatabaseUserWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DatabaseUsersApi_GetDatabaseUserWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetDatabaseUserApiParams) admin.GetDatabaseUserApiRequest) *DatabaseUsersApi_GetDatabaseUserWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListDatabaseUsers provides a mock function with given fields: ctx, groupId
func (_m *DatabaseUsersApi) ListDatabaseUsers(ctx context.Context, groupId string) admin.ListDatabaseUsersApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for ListDatabaseUsers")
	}

	var r0 admin.ListDatabaseUsersApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ListDatabaseUsersApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.ListDatabaseUsersApiRequest)
	}

	return r0
}

// DatabaseUsersApi_ListDatabaseUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDatabaseUsers'
type DatabaseUsersApi_ListDatabaseUsers_Call struct {
	*mock.Call
}

// ListDatabaseUsers is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *DatabaseUsersApi_Expecter) ListDatabaseUsers(ctx any, groupId any) *DatabaseUsersApi_ListDatabaseUsers_Call {
	return &DatabaseUsersApi_ListDatabaseUsers_Call{Call: _e.mock.On("ListDatabaseUsers", ctx, groupId)}
}

func (_c *DatabaseUsersApi_ListDatabaseUsers_Call) Run(run func(ctx context.Context, groupId string)) *DatabaseUsersApi_ListDatabaseUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DatabaseUsersApi_ListDatabaseUsers_Call) Return(_a0 admin.ListDatabaseUsersApiRequest) *DatabaseUsersApi_ListDatabaseUsers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DatabaseUsersApi_ListDatabaseUsers_Call) RunAndReturn(run func(context.Context, string) admin.ListDatabaseUsersApiRequest) *DatabaseUsersApi_ListDatabaseUsers_Call {
	_c.Call.Return(run)
	return _c
}

// ListDatabaseUsersExecute provides a mock function with given fields: r
func (_m *DatabaseUsersApi) ListDatabaseUsersExecute(r admin.ListDatabaseUsersApiRequest) (*admin.PaginatedApiAtlasDatabaseUser, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListDatabaseUsersExecute")
	}

	var r0 *admin.PaginatedApiAtlasDatabaseUser
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListDatabaseUsersApiRequest) (*admin.PaginatedApiAtlasDatabaseUser, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListDatabaseUsersApiRequest) *admin.PaginatedApiAtlasDatabaseUser); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedApiAtlasDatabaseUser)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListDatabaseUsersApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListDatabaseUsersApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DatabaseUsersApi_ListDatabaseUsersExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDatabaseUsersExecute'
type DatabaseUsersApi_ListDatabaseUsersExecute_Call struct {
	*mock.Call
}

// ListDatabaseUsersExecute is a helper method to define mock.On call
//   - r admin.ListDatabaseUsersApiRequest
func (_e *DatabaseUsersApi_Expecter) ListDatabaseUsersExecute(r any) *DatabaseUsersApi_ListDatabaseUsersExecute_Call {
	return &DatabaseUsersApi_ListDatabaseUsersExecute_Call{Call: _e.mock.On("ListDatabaseUsersExecute", r)}
}

func (_c *DatabaseUsersApi_ListDatabaseUsersExecute_Call) Run(run func(r admin.ListDatabaseUsersApiRequest)) *DatabaseUsersApi_ListDatabaseUsersExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListDatabaseUsersApiRequest))
	})
	return _c
}

func (_c *DatabaseUsersApi_ListDatabaseUsersExecute_Call) Return(_a0 *admin.PaginatedApiAtlasDatabaseUser, _a1 *http.Response, _a2 error) *DatabaseUsersApi_ListDatabaseUsersExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *DatabaseUsersApi_ListDatabaseUsersExecute_Call) RunAndReturn(run func(admin.ListDatabaseUsersApiRequest) (*admin.PaginatedApiAtlasDatabaseUser, *http.Response, error)) *DatabaseUsersApi_ListDatabaseUsersExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListDatabaseUsersWithParams provides a mock function with given fields: ctx, args
func (_m *DatabaseUsersApi) ListDatabaseUsersWithParams(ctx context.Context, args *admin.ListDatabaseUsersApiParams) admin.ListDatabaseUsersApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListDatabaseUsersWithParams")
	}

	var r0 admin.ListDatabaseUsersApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListDatabaseUsersApiParams) admin.ListDatabaseUsersApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListDatabaseUsersApiRequest)
	}

	return r0
}

// DatabaseUsersApi_ListDatabaseUsersWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDatabaseUsersWithParams'
type DatabaseUsersApi_ListDatabaseUsersWithParams_Call struct {
	*mock.Call
}

// ListDatabaseUsersWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListDatabaseUsersApiParams
func (_e *DatabaseUsersApi_Expecter) ListDatabaseUsersWithParams(ctx any, args any) *DatabaseUsersApi_ListDatabaseUsersWithParams_Call {
	return &DatabaseUsersApi_ListDatabaseUsersWithParams_Call{Call: _e.mock.On("ListDatabaseUsersWithParams", ctx, args)}
}

func (_c *DatabaseUsersApi_ListDatabaseUsersWithParams_Call) Run(run func(ctx context.Context, args *admin.ListDatabaseUsersApiParams)) *DatabaseUsersApi_ListDatabaseUsersWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListDatabaseUsersApiParams))
	})
	return _c
}

func (_c *DatabaseUsersApi_ListDatabaseUsersWithParams_Call) Return(_a0 admin.ListDatabaseUsersApiRequest) *DatabaseUsersApi_ListDatabaseUsersWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DatabaseUsersApi_ListDatabaseUsersWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListDatabaseUsersApiParams) admin.ListDatabaseUsersApiRequest) *DatabaseUsersApi_ListDatabaseUsersWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDatabaseUser provides a mock function with given fields: ctx, groupId, databaseName, username, cloudDatabaseUser
func (_m *DatabaseUsersApi) UpdateDatabaseUser(ctx context.Context, groupId string, databaseName string, username string, cloudDatabaseUser *admin.CloudDatabaseUser) admin.UpdateDatabaseUserApiRequest {
	ret := _m.Called(ctx, groupId, databaseName, username, cloudDatabaseUser)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDatabaseUser")
	}

	var r0 admin.UpdateDatabaseUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *admin.CloudDatabaseUser) admin.UpdateDatabaseUserApiRequest); ok {
		r0 = rf(ctx, groupId, databaseName, username, cloudDatabaseUser)
	} else {
		r0 = ret.Get(0).(admin.UpdateDatabaseUserApiRequest)
	}

	return r0
}

// DatabaseUsersApi_UpdateDatabaseUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDatabaseUser'
type DatabaseUsersApi_UpdateDatabaseUser_Call struct {
	*mock.Call
}

// UpdateDatabaseUser is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - databaseName string
//   - username string
//   - cloudDatabaseUser *admin.CloudDatabaseUser
func (_e *DatabaseUsersApi_Expecter) UpdateDatabaseUser(ctx any, groupId any, databaseName any, username any, cloudDatabaseUser any) *DatabaseUsersApi_UpdateDatabaseUser_Call {
	return &DatabaseUsersApi_UpdateDatabaseUser_Call{Call: _e.mock.On("UpdateDatabaseUser", ctx, groupId, databaseName, username, cloudDatabaseUser)}
}

func (_c *DatabaseUsersApi_UpdateDatabaseUser_Call) Run(run func(ctx context.Context, groupId string, databaseName string, username string, cloudDatabaseUser *admin.CloudDatabaseUser)) *DatabaseUsersApi_UpdateDatabaseUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(*admin.CloudDatabaseUser))
	})
	return _c
}

func (_c *DatabaseUsersApi_UpdateDatabaseUser_Call) Return(_a0 admin.UpdateDatabaseUserApiRequest) *DatabaseUsersApi_UpdateDatabaseUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DatabaseUsersApi_UpdateDatabaseUser_Call) RunAndReturn(run func(context.Context, string, string, string, *admin.CloudDatabaseUser) admin.UpdateDatabaseUserApiRequest) *DatabaseUsersApi_UpdateDatabaseUser_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDatabaseUserExecute provides a mock function with given fields: r
func (_m *DatabaseUsersApi) UpdateDatabaseUserExecute(r admin.UpdateDatabaseUserApiRequest) (*admin.CloudDatabaseUser, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDatabaseUserExecute")
	}

	var r0 *admin.CloudDatabaseUser
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.UpdateDatabaseUserApiRequest) (*admin.CloudDatabaseUser, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.UpdateDatabaseUserApiRequest) *admin.CloudDatabaseUser); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CloudDatabaseUser)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.UpdateDatabaseUserApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.UpdateDatabaseUserApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DatabaseUsersApi_UpdateDatabaseUserExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDatabaseUserExecute'
type DatabaseUsersApi_UpdateDatabaseUserExecute_Call struct {
	*mock.Call
}

// UpdateDatabaseUserExecute is a helper method to define mock.On call
//   - r admin.UpdateDatabaseUserApiRequest
func (_e *DatabaseUsersApi_Expecter) UpdateDatabaseUserExecute(r any) *DatabaseUsersApi_UpdateDatabaseUserExecute_Call {
	return &DatabaseUsersApi_UpdateDatabaseUserExecute_Call{Call: _e.mock.On("UpdateDatabaseUserExecute", r)}
}

func (_c *DatabaseUsersApi_UpdateDatabaseUserExecute_Call) Run(run func(r admin.UpdateDatabaseUserApiRequest)) *DatabaseUsersApi_UpdateDatabaseUserExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.UpdateDatabaseUserApiRequest))
	})
	return _c
}

func (_c *DatabaseUsersApi_UpdateDatabaseUserExecute_Call) Return(_a0 *admin.CloudDatabaseUser, _a1 *http.Response, _a2 error) *DatabaseUsersApi_UpdateDatabaseUserExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *DatabaseUsersApi_UpdateDatabaseUserExecute_Call) RunAndReturn(run func(admin.UpdateDatabaseUserApiRequest) (*admin.CloudDatabaseUser, *http.Response, error)) *DatabaseUsersApi_UpdateDatabaseUserExecute_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDatabaseUserWithParams provides a mock function with given fields: ctx, args
func (_m *DatabaseUsersApi) UpdateDatabaseUserWithParams(ctx context.Context, args *admin.UpdateDatabaseUserApiParams) admin.UpdateDatabaseUserApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDatabaseUserWithParams")
	}

	var r0 admin.UpdateDatabaseUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.UpdateDatabaseUserApiParams) admin.UpdateDatabaseUserApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.UpdateDatabaseUserApiRequest)
	}

	return r0
}

// DatabaseUsersApi_UpdateDatabaseUserWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDatabaseUserWithParams'
type DatabaseUsersApi_UpdateDatabaseUserWithParams_Call struct {
	*mock.Call
}

// UpdateDatabaseUserWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.UpdateDatabaseUserApiParams
func (_e *DatabaseUsersApi_Expecter) UpdateDatabaseUserWithParams(ctx any, args any) *DatabaseUsersApi_UpdateDatabaseUserWithParams_Call {
	return &DatabaseUsersApi_UpdateDatabaseUserWithParams_Call{Call: _e.mock.On("UpdateDatabaseUserWithParams", ctx, args)}
}

func (_c *DatabaseUsersApi_UpdateDatabaseUserWithParams_Call) Run(run func(ctx context.Context, args *admin.UpdateDatabaseUserApiParams)) *DatabaseUsersApi_UpdateDatabaseUserWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.UpdateDatabaseUserApiParams))
	})
	return _c
}

func (_c *DatabaseUsersApi_UpdateDatabaseUserWithParams_Call) Return(_a0 admin.UpdateDatabaseUserApiRequest) *DatabaseUsersApi_UpdateDatabaseUserWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DatabaseUsersApi_UpdateDatabaseUserWithParams_Call) RunAndReturn(run func(context.Context, *admin.UpdateDatabaseUserApiParams) admin.UpdateDatabaseUserApiRequest) *DatabaseUsersApi_UpdateDatabaseUserWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewDatabaseUsersApi creates a new instance of DatabaseUsersApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatabaseUsersApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *DatabaseUsersApi {
	mock := &DatabaseUsersApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
