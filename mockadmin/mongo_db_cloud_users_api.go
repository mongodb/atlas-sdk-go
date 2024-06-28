// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20240530002/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MongoDBCloudUsersApi is an autogenerated mock type for the MongoDBCloudUsersApi type
type MongoDBCloudUsersApi struct {
	mock.Mock
}

type MongoDBCloudUsersApi_Expecter struct {
	mock *mock.Mock
}

func (_m *MongoDBCloudUsersApi) EXPECT() *MongoDBCloudUsersApi_Expecter {
	return &MongoDBCloudUsersApi_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, cloudAppUser
func (_m *MongoDBCloudUsersApi) CreateUser(ctx context.Context, cloudAppUser *admin.CloudAppUser) admin.CreateUserApiRequest {
	ret := _m.Called(ctx, cloudAppUser)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 admin.CreateUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CloudAppUser) admin.CreateUserApiRequest); ok {
		r0 = rf(ctx, cloudAppUser)
	} else {
		r0 = ret.Get(0).(admin.CreateUserApiRequest)
	}

	return r0
}

// MongoDBCloudUsersApi_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MongoDBCloudUsersApi_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - cloudAppUser *admin.CloudAppUser
func (_e *MongoDBCloudUsersApi_Expecter) CreateUser(ctx interface{}, cloudAppUser interface{}) *MongoDBCloudUsersApi_CreateUser_Call {
	return &MongoDBCloudUsersApi_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, cloudAppUser)}
}

func (_c *MongoDBCloudUsersApi_CreateUser_Call) Run(run func(ctx context.Context, cloudAppUser *admin.CloudAppUser)) *MongoDBCloudUsersApi_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CloudAppUser))
	})
	return _c
}

func (_c *MongoDBCloudUsersApi_CreateUser_Call) Return(_a0 admin.CreateUserApiRequest) *MongoDBCloudUsersApi_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MongoDBCloudUsersApi_CreateUser_Call) RunAndReturn(run func(context.Context, *admin.CloudAppUser) admin.CreateUserApiRequest) *MongoDBCloudUsersApi_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUserExecute provides a mock function with given fields: r
func (_m *MongoDBCloudUsersApi) CreateUserExecute(r admin.CreateUserApiRequest) (*admin.CloudAppUser, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateUserExecute")
	}

	var r0 *admin.CloudAppUser
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateUserApiRequest) (*admin.CloudAppUser, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateUserApiRequest) *admin.CloudAppUser); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CloudAppUser)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateUserApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateUserApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MongoDBCloudUsersApi_CreateUserExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUserExecute'
type MongoDBCloudUsersApi_CreateUserExecute_Call struct {
	*mock.Call
}

// CreateUserExecute is a helper method to define mock.On call
//   - r admin.CreateUserApiRequest
func (_e *MongoDBCloudUsersApi_Expecter) CreateUserExecute(r interface{}) *MongoDBCloudUsersApi_CreateUserExecute_Call {
	return &MongoDBCloudUsersApi_CreateUserExecute_Call{Call: _e.mock.On("CreateUserExecute", r)}
}

func (_c *MongoDBCloudUsersApi_CreateUserExecute_Call) Run(run func(r admin.CreateUserApiRequest)) *MongoDBCloudUsersApi_CreateUserExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateUserApiRequest))
	})
	return _c
}

func (_c *MongoDBCloudUsersApi_CreateUserExecute_Call) Return(_a0 *admin.CloudAppUser, _a1 *http.Response, _a2 error) *MongoDBCloudUsersApi_CreateUserExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MongoDBCloudUsersApi_CreateUserExecute_Call) RunAndReturn(run func(admin.CreateUserApiRequest) (*admin.CloudAppUser, *http.Response, error)) *MongoDBCloudUsersApi_CreateUserExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUserWithParams provides a mock function with given fields: ctx, args
func (_m *MongoDBCloudUsersApi) CreateUserWithParams(ctx context.Context, args *admin.CreateUserApiParams) admin.CreateUserApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateUserWithParams")
	}

	var r0 admin.CreateUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateUserApiParams) admin.CreateUserApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateUserApiRequest)
	}

	return r0
}

// MongoDBCloudUsersApi_CreateUserWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUserWithParams'
type MongoDBCloudUsersApi_CreateUserWithParams_Call struct {
	*mock.Call
}

// CreateUserWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateUserApiParams
func (_e *MongoDBCloudUsersApi_Expecter) CreateUserWithParams(ctx interface{}, args interface{}) *MongoDBCloudUsersApi_CreateUserWithParams_Call {
	return &MongoDBCloudUsersApi_CreateUserWithParams_Call{Call: _e.mock.On("CreateUserWithParams", ctx, args)}
}

func (_c *MongoDBCloudUsersApi_CreateUserWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateUserApiParams)) *MongoDBCloudUsersApi_CreateUserWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateUserApiParams))
	})
	return _c
}

func (_c *MongoDBCloudUsersApi_CreateUserWithParams_Call) Return(_a0 admin.CreateUserApiRequest) *MongoDBCloudUsersApi_CreateUserWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MongoDBCloudUsersApi_CreateUserWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateUserApiParams) admin.CreateUserApiRequest) *MongoDBCloudUsersApi_CreateUserWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetUser provides a mock function with given fields: ctx, userId
func (_m *MongoDBCloudUsersApi) GetUser(ctx context.Context, userId string) admin.GetUserApiRequest {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 admin.GetUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.GetUserApiRequest); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(admin.GetUserApiRequest)
	}

	return r0
}

// MongoDBCloudUsersApi_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type MongoDBCloudUsersApi_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - ctx context.Context
//   - userId string
func (_e *MongoDBCloudUsersApi_Expecter) GetUser(ctx interface{}, userId interface{}) *MongoDBCloudUsersApi_GetUser_Call {
	return &MongoDBCloudUsersApi_GetUser_Call{Call: _e.mock.On("GetUser", ctx, userId)}
}

func (_c *MongoDBCloudUsersApi_GetUser_Call) Run(run func(ctx context.Context, userId string)) *MongoDBCloudUsersApi_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUser_Call) Return(_a0 admin.GetUserApiRequest) *MongoDBCloudUsersApi_GetUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUser_Call) RunAndReturn(run func(context.Context, string) admin.GetUserApiRequest) *MongoDBCloudUsersApi_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByUsername provides a mock function with given fields: ctx, userName
func (_m *MongoDBCloudUsersApi) GetUserByUsername(ctx context.Context, userName string) admin.GetUserByUsernameApiRequest {
	ret := _m.Called(ctx, userName)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsername")
	}

	var r0 admin.GetUserByUsernameApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.GetUserByUsernameApiRequest); ok {
		r0 = rf(ctx, userName)
	} else {
		r0 = ret.Get(0).(admin.GetUserByUsernameApiRequest)
	}

	return r0
}

// MongoDBCloudUsersApi_GetUserByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByUsername'
type MongoDBCloudUsersApi_GetUserByUsername_Call struct {
	*mock.Call
}

// GetUserByUsername is a helper method to define mock.On call
//   - ctx context.Context
//   - userName string
func (_e *MongoDBCloudUsersApi_Expecter) GetUserByUsername(ctx interface{}, userName interface{}) *MongoDBCloudUsersApi_GetUserByUsername_Call {
	return &MongoDBCloudUsersApi_GetUserByUsername_Call{Call: _e.mock.On("GetUserByUsername", ctx, userName)}
}

func (_c *MongoDBCloudUsersApi_GetUserByUsername_Call) Run(run func(ctx context.Context, userName string)) *MongoDBCloudUsersApi_GetUserByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUserByUsername_Call) Return(_a0 admin.GetUserByUsernameApiRequest) *MongoDBCloudUsersApi_GetUserByUsername_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUserByUsername_Call) RunAndReturn(run func(context.Context, string) admin.GetUserByUsernameApiRequest) *MongoDBCloudUsersApi_GetUserByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByUsernameExecute provides a mock function with given fields: r
func (_m *MongoDBCloudUsersApi) GetUserByUsernameExecute(r admin.GetUserByUsernameApiRequest) (*admin.CloudAppUser, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsernameExecute")
	}

	var r0 *admin.CloudAppUser
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetUserByUsernameApiRequest) (*admin.CloudAppUser, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetUserByUsernameApiRequest) *admin.CloudAppUser); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CloudAppUser)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetUserByUsernameApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetUserByUsernameApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MongoDBCloudUsersApi_GetUserByUsernameExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByUsernameExecute'
type MongoDBCloudUsersApi_GetUserByUsernameExecute_Call struct {
	*mock.Call
}

// GetUserByUsernameExecute is a helper method to define mock.On call
//   - r admin.GetUserByUsernameApiRequest
func (_e *MongoDBCloudUsersApi_Expecter) GetUserByUsernameExecute(r interface{}) *MongoDBCloudUsersApi_GetUserByUsernameExecute_Call {
	return &MongoDBCloudUsersApi_GetUserByUsernameExecute_Call{Call: _e.mock.On("GetUserByUsernameExecute", r)}
}

func (_c *MongoDBCloudUsersApi_GetUserByUsernameExecute_Call) Run(run func(r admin.GetUserByUsernameApiRequest)) *MongoDBCloudUsersApi_GetUserByUsernameExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetUserByUsernameApiRequest))
	})
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUserByUsernameExecute_Call) Return(_a0 *admin.CloudAppUser, _a1 *http.Response, _a2 error) *MongoDBCloudUsersApi_GetUserByUsernameExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUserByUsernameExecute_Call) RunAndReturn(run func(admin.GetUserByUsernameApiRequest) (*admin.CloudAppUser, *http.Response, error)) *MongoDBCloudUsersApi_GetUserByUsernameExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByUsernameWithParams provides a mock function with given fields: ctx, args
func (_m *MongoDBCloudUsersApi) GetUserByUsernameWithParams(ctx context.Context, args *admin.GetUserByUsernameApiParams) admin.GetUserByUsernameApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsernameWithParams")
	}

	var r0 admin.GetUserByUsernameApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetUserByUsernameApiParams) admin.GetUserByUsernameApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetUserByUsernameApiRequest)
	}

	return r0
}

// MongoDBCloudUsersApi_GetUserByUsernameWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByUsernameWithParams'
type MongoDBCloudUsersApi_GetUserByUsernameWithParams_Call struct {
	*mock.Call
}

// GetUserByUsernameWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetUserByUsernameApiParams
func (_e *MongoDBCloudUsersApi_Expecter) GetUserByUsernameWithParams(ctx interface{}, args interface{}) *MongoDBCloudUsersApi_GetUserByUsernameWithParams_Call {
	return &MongoDBCloudUsersApi_GetUserByUsernameWithParams_Call{Call: _e.mock.On("GetUserByUsernameWithParams", ctx, args)}
}

func (_c *MongoDBCloudUsersApi_GetUserByUsernameWithParams_Call) Run(run func(ctx context.Context, args *admin.GetUserByUsernameApiParams)) *MongoDBCloudUsersApi_GetUserByUsernameWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetUserByUsernameApiParams))
	})
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUserByUsernameWithParams_Call) Return(_a0 admin.GetUserByUsernameApiRequest) *MongoDBCloudUsersApi_GetUserByUsernameWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUserByUsernameWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetUserByUsernameApiParams) admin.GetUserByUsernameApiRequest) *MongoDBCloudUsersApi_GetUserByUsernameWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserExecute provides a mock function with given fields: r
func (_m *MongoDBCloudUsersApi) GetUserExecute(r admin.GetUserApiRequest) (*admin.CloudAppUser, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetUserExecute")
	}

	var r0 *admin.CloudAppUser
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetUserApiRequest) (*admin.CloudAppUser, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetUserApiRequest) *admin.CloudAppUser); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CloudAppUser)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetUserApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetUserApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MongoDBCloudUsersApi_GetUserExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserExecute'
type MongoDBCloudUsersApi_GetUserExecute_Call struct {
	*mock.Call
}

// GetUserExecute is a helper method to define mock.On call
//   - r admin.GetUserApiRequest
func (_e *MongoDBCloudUsersApi_Expecter) GetUserExecute(r interface{}) *MongoDBCloudUsersApi_GetUserExecute_Call {
	return &MongoDBCloudUsersApi_GetUserExecute_Call{Call: _e.mock.On("GetUserExecute", r)}
}

func (_c *MongoDBCloudUsersApi_GetUserExecute_Call) Run(run func(r admin.GetUserApiRequest)) *MongoDBCloudUsersApi_GetUserExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetUserApiRequest))
	})
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUserExecute_Call) Return(_a0 *admin.CloudAppUser, _a1 *http.Response, _a2 error) *MongoDBCloudUsersApi_GetUserExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUserExecute_Call) RunAndReturn(run func(admin.GetUserApiRequest) (*admin.CloudAppUser, *http.Response, error)) *MongoDBCloudUsersApi_GetUserExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserWithParams provides a mock function with given fields: ctx, args
func (_m *MongoDBCloudUsersApi) GetUserWithParams(ctx context.Context, args *admin.GetUserApiParams) admin.GetUserApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetUserWithParams")
	}

	var r0 admin.GetUserApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetUserApiParams) admin.GetUserApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetUserApiRequest)
	}

	return r0
}

// MongoDBCloudUsersApi_GetUserWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserWithParams'
type MongoDBCloudUsersApi_GetUserWithParams_Call struct {
	*mock.Call
}

// GetUserWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetUserApiParams
func (_e *MongoDBCloudUsersApi_Expecter) GetUserWithParams(ctx interface{}, args interface{}) *MongoDBCloudUsersApi_GetUserWithParams_Call {
	return &MongoDBCloudUsersApi_GetUserWithParams_Call{Call: _e.mock.On("GetUserWithParams", ctx, args)}
}

func (_c *MongoDBCloudUsersApi_GetUserWithParams_Call) Run(run func(ctx context.Context, args *admin.GetUserApiParams)) *MongoDBCloudUsersApi_GetUserWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetUserApiParams))
	})
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUserWithParams_Call) Return(_a0 admin.GetUserApiRequest) *MongoDBCloudUsersApi_GetUserWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MongoDBCloudUsersApi_GetUserWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetUserApiParams) admin.GetUserApiRequest) *MongoDBCloudUsersApi_GetUserWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewMongoDBCloudUsersApi creates a new instance of MongoDBCloudUsersApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMongoDBCloudUsersApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *MongoDBCloudUsersApi {
	mock := &MongoDBCloudUsersApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
