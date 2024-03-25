// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20231115008/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// CloudProviderAccessApi is an autogenerated mock type for the CloudProviderAccessApi type
type CloudProviderAccessApi struct {
	mock.Mock
}

type CloudProviderAccessApi_Expecter struct {
	mock *mock.Mock
}

func (_m *CloudProviderAccessApi) EXPECT() *CloudProviderAccessApi_Expecter {
	return &CloudProviderAccessApi_Expecter{mock: &_m.Mock}
}

// AuthorizeCloudProviderAccessRole provides a mock function with given fields: ctx, groupId, roleId, cloudProviderAccessRole
func (_m *CloudProviderAccessApi) AuthorizeCloudProviderAccessRole(ctx context.Context, groupId string, roleId string, cloudProviderAccessRole *admin.CloudProviderAccessRole) admin.AuthorizeCloudProviderAccessRoleApiRequest {
	ret := _m.Called(ctx, groupId, roleId, cloudProviderAccessRole)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizeCloudProviderAccessRole")
	}

	var r0 admin.AuthorizeCloudProviderAccessRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.CloudProviderAccessRole) admin.AuthorizeCloudProviderAccessRoleApiRequest); ok {
		r0 = rf(ctx, groupId, roleId, cloudProviderAccessRole)
	} else {
		r0 = ret.Get(0).(admin.AuthorizeCloudProviderAccessRoleApiRequest)
	}

	return r0
}

// CloudProviderAccessApi_AuthorizeCloudProviderAccessRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthorizeCloudProviderAccessRole'
type CloudProviderAccessApi_AuthorizeCloudProviderAccessRole_Call struct {
	*mock.Call
}

// AuthorizeCloudProviderAccessRole is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - roleId string
//   - cloudProviderAccessRole *admin.CloudProviderAccessRole
func (_e *CloudProviderAccessApi_Expecter) AuthorizeCloudProviderAccessRole(ctx interface{}, groupId interface{}, roleId interface{}, cloudProviderAccessRole interface{}) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRole_Call {
	return &CloudProviderAccessApi_AuthorizeCloudProviderAccessRole_Call{Call: _e.mock.On("AuthorizeCloudProviderAccessRole", ctx, groupId, roleId, cloudProviderAccessRole)}
}

func (_c *CloudProviderAccessApi_AuthorizeCloudProviderAccessRole_Call) Run(run func(ctx context.Context, groupId string, roleId string, cloudProviderAccessRole *admin.CloudProviderAccessRole)) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.CloudProviderAccessRole))
	})
	return _c
}

func (_c *CloudProviderAccessApi_AuthorizeCloudProviderAccessRole_Call) Return(_a0 admin.AuthorizeCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloudProviderAccessApi_AuthorizeCloudProviderAccessRole_Call) RunAndReturn(run func(context.Context, string, string, *admin.CloudProviderAccessRole) admin.AuthorizeCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRole_Call {
	_c.Call.Return(run)
	return _c
}

// AuthorizeCloudProviderAccessRoleExecute provides a mock function with given fields: r
func (_m *CloudProviderAccessApi) AuthorizeCloudProviderAccessRoleExecute(r admin.AuthorizeCloudProviderAccessRoleApiRequest) (*admin.CloudProviderAccessRole, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizeCloudProviderAccessRoleExecute")
	}

	var r0 *admin.CloudProviderAccessRole
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.AuthorizeCloudProviderAccessRoleApiRequest) (*admin.CloudProviderAccessRole, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.AuthorizeCloudProviderAccessRoleApiRequest) *admin.CloudProviderAccessRole); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CloudProviderAccessRole)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.AuthorizeCloudProviderAccessRoleApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.AuthorizeCloudProviderAccessRoleApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthorizeCloudProviderAccessRoleExecute'
type CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleExecute_Call struct {
	*mock.Call
}

// AuthorizeCloudProviderAccessRoleExecute is a helper method to define mock.On call
//   - r admin.AuthorizeCloudProviderAccessRoleApiRequest
func (_e *CloudProviderAccessApi_Expecter) AuthorizeCloudProviderAccessRoleExecute(r interface{}) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleExecute_Call {
	return &CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleExecute_Call{Call: _e.mock.On("AuthorizeCloudProviderAccessRoleExecute", r)}
}

func (_c *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleExecute_Call) Run(run func(r admin.AuthorizeCloudProviderAccessRoleApiRequest)) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.AuthorizeCloudProviderAccessRoleApiRequest))
	})
	return _c
}

func (_c *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleExecute_Call) Return(_a0 *admin.CloudProviderAccessRole, _a1 *http.Response, _a2 error) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleExecute_Call) RunAndReturn(run func(admin.AuthorizeCloudProviderAccessRoleApiRequest) (*admin.CloudProviderAccessRole, *http.Response, error)) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleExecute_Call {
	_c.Call.Return(run)
	return _c
}

// AuthorizeCloudProviderAccessRoleWithParams provides a mock function with given fields: ctx, args
func (_m *CloudProviderAccessApi) AuthorizeCloudProviderAccessRoleWithParams(ctx context.Context, args *admin.AuthorizeCloudProviderAccessRoleApiParams) admin.AuthorizeCloudProviderAccessRoleApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizeCloudProviderAccessRoleWithParams")
	}

	var r0 admin.AuthorizeCloudProviderAccessRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.AuthorizeCloudProviderAccessRoleApiParams) admin.AuthorizeCloudProviderAccessRoleApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.AuthorizeCloudProviderAccessRoleApiRequest)
	}

	return r0
}

// CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthorizeCloudProviderAccessRoleWithParams'
type CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleWithParams_Call struct {
	*mock.Call
}

// AuthorizeCloudProviderAccessRoleWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.AuthorizeCloudProviderAccessRoleApiParams
func (_e *CloudProviderAccessApi_Expecter) AuthorizeCloudProviderAccessRoleWithParams(ctx interface{}, args interface{}) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleWithParams_Call {
	return &CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleWithParams_Call{Call: _e.mock.On("AuthorizeCloudProviderAccessRoleWithParams", ctx, args)}
}

func (_c *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleWithParams_Call) Run(run func(ctx context.Context, args *admin.AuthorizeCloudProviderAccessRoleApiParams)) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.AuthorizeCloudProviderAccessRoleApiParams))
	})
	return _c
}

func (_c *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleWithParams_Call) Return(_a0 admin.AuthorizeCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleWithParams_Call) RunAndReturn(run func(context.Context, *admin.AuthorizeCloudProviderAccessRoleApiParams) admin.AuthorizeCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_AuthorizeCloudProviderAccessRoleWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCloudProviderAccessRole provides a mock function with given fields: ctx, groupId, cloudProviderAccessRole
func (_m *CloudProviderAccessApi) CreateCloudProviderAccessRole(ctx context.Context, groupId string, cloudProviderAccessRole *admin.CloudProviderAccessRole) admin.CreateCloudProviderAccessRoleApiRequest {
	ret := _m.Called(ctx, groupId, cloudProviderAccessRole)

	if len(ret) == 0 {
		panic("no return value specified for CreateCloudProviderAccessRole")
	}

	var r0 admin.CreateCloudProviderAccessRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.CloudProviderAccessRole) admin.CreateCloudProviderAccessRoleApiRequest); ok {
		r0 = rf(ctx, groupId, cloudProviderAccessRole)
	} else {
		r0 = ret.Get(0).(admin.CreateCloudProviderAccessRoleApiRequest)
	}

	return r0
}

// CloudProviderAccessApi_CreateCloudProviderAccessRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCloudProviderAccessRole'
type CloudProviderAccessApi_CreateCloudProviderAccessRole_Call struct {
	*mock.Call
}

// CreateCloudProviderAccessRole is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - cloudProviderAccessRole *admin.CloudProviderAccessRole
func (_e *CloudProviderAccessApi_Expecter) CreateCloudProviderAccessRole(ctx interface{}, groupId interface{}, cloudProviderAccessRole interface{}) *CloudProviderAccessApi_CreateCloudProviderAccessRole_Call {
	return &CloudProviderAccessApi_CreateCloudProviderAccessRole_Call{Call: _e.mock.On("CreateCloudProviderAccessRole", ctx, groupId, cloudProviderAccessRole)}
}

func (_c *CloudProviderAccessApi_CreateCloudProviderAccessRole_Call) Run(run func(ctx context.Context, groupId string, cloudProviderAccessRole *admin.CloudProviderAccessRole)) *CloudProviderAccessApi_CreateCloudProviderAccessRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.CloudProviderAccessRole))
	})
	return _c
}

func (_c *CloudProviderAccessApi_CreateCloudProviderAccessRole_Call) Return(_a0 admin.CreateCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_CreateCloudProviderAccessRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloudProviderAccessApi_CreateCloudProviderAccessRole_Call) RunAndReturn(run func(context.Context, string, *admin.CloudProviderAccessRole) admin.CreateCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_CreateCloudProviderAccessRole_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCloudProviderAccessRoleExecute provides a mock function with given fields: r
func (_m *CloudProviderAccessApi) CreateCloudProviderAccessRoleExecute(r admin.CreateCloudProviderAccessRoleApiRequest) (*admin.CloudProviderAccessRole, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateCloudProviderAccessRoleExecute")
	}

	var r0 *admin.CloudProviderAccessRole
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateCloudProviderAccessRoleApiRequest) (*admin.CloudProviderAccessRole, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateCloudProviderAccessRoleApiRequest) *admin.CloudProviderAccessRole); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CloudProviderAccessRole)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateCloudProviderAccessRoleApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateCloudProviderAccessRoleApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CloudProviderAccessApi_CreateCloudProviderAccessRoleExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCloudProviderAccessRoleExecute'
type CloudProviderAccessApi_CreateCloudProviderAccessRoleExecute_Call struct {
	*mock.Call
}

// CreateCloudProviderAccessRoleExecute is a helper method to define mock.On call
//   - r admin.CreateCloudProviderAccessRoleApiRequest
func (_e *CloudProviderAccessApi_Expecter) CreateCloudProviderAccessRoleExecute(r interface{}) *CloudProviderAccessApi_CreateCloudProviderAccessRoleExecute_Call {
	return &CloudProviderAccessApi_CreateCloudProviderAccessRoleExecute_Call{Call: _e.mock.On("CreateCloudProviderAccessRoleExecute", r)}
}

func (_c *CloudProviderAccessApi_CreateCloudProviderAccessRoleExecute_Call) Run(run func(r admin.CreateCloudProviderAccessRoleApiRequest)) *CloudProviderAccessApi_CreateCloudProviderAccessRoleExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateCloudProviderAccessRoleApiRequest))
	})
	return _c
}

func (_c *CloudProviderAccessApi_CreateCloudProviderAccessRoleExecute_Call) Return(_a0 *admin.CloudProviderAccessRole, _a1 *http.Response, _a2 error) *CloudProviderAccessApi_CreateCloudProviderAccessRoleExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CloudProviderAccessApi_CreateCloudProviderAccessRoleExecute_Call) RunAndReturn(run func(admin.CreateCloudProviderAccessRoleApiRequest) (*admin.CloudProviderAccessRole, *http.Response, error)) *CloudProviderAccessApi_CreateCloudProviderAccessRoleExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCloudProviderAccessRoleWithParams provides a mock function with given fields: ctx, args
func (_m *CloudProviderAccessApi) CreateCloudProviderAccessRoleWithParams(ctx context.Context, args *admin.CreateCloudProviderAccessRoleApiParams) admin.CreateCloudProviderAccessRoleApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateCloudProviderAccessRoleWithParams")
	}

	var r0 admin.CreateCloudProviderAccessRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateCloudProviderAccessRoleApiParams) admin.CreateCloudProviderAccessRoleApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateCloudProviderAccessRoleApiRequest)
	}

	return r0
}

// CloudProviderAccessApi_CreateCloudProviderAccessRoleWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCloudProviderAccessRoleWithParams'
type CloudProviderAccessApi_CreateCloudProviderAccessRoleWithParams_Call struct {
	*mock.Call
}

// CreateCloudProviderAccessRoleWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateCloudProviderAccessRoleApiParams
func (_e *CloudProviderAccessApi_Expecter) CreateCloudProviderAccessRoleWithParams(ctx interface{}, args interface{}) *CloudProviderAccessApi_CreateCloudProviderAccessRoleWithParams_Call {
	return &CloudProviderAccessApi_CreateCloudProviderAccessRoleWithParams_Call{Call: _e.mock.On("CreateCloudProviderAccessRoleWithParams", ctx, args)}
}

func (_c *CloudProviderAccessApi_CreateCloudProviderAccessRoleWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateCloudProviderAccessRoleApiParams)) *CloudProviderAccessApi_CreateCloudProviderAccessRoleWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateCloudProviderAccessRoleApiParams))
	})
	return _c
}

func (_c *CloudProviderAccessApi_CreateCloudProviderAccessRoleWithParams_Call) Return(_a0 admin.CreateCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_CreateCloudProviderAccessRoleWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloudProviderAccessApi_CreateCloudProviderAccessRoleWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateCloudProviderAccessRoleApiParams) admin.CreateCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_CreateCloudProviderAccessRoleWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// DeauthorizeCloudProviderAccessRole provides a mock function with given fields: ctx, groupId, cloudProvider, roleId
func (_m *CloudProviderAccessApi) DeauthorizeCloudProviderAccessRole(ctx context.Context, groupId string, cloudProvider string, roleId string) admin.DeauthorizeCloudProviderAccessRoleApiRequest {
	ret := _m.Called(ctx, groupId, cloudProvider, roleId)

	if len(ret) == 0 {
		panic("no return value specified for DeauthorizeCloudProviderAccessRole")
	}

	var r0 admin.DeauthorizeCloudProviderAccessRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) admin.DeauthorizeCloudProviderAccessRoleApiRequest); ok {
		r0 = rf(ctx, groupId, cloudProvider, roleId)
	} else {
		r0 = ret.Get(0).(admin.DeauthorizeCloudProviderAccessRoleApiRequest)
	}

	return r0
}

// CloudProviderAccessApi_DeauthorizeCloudProviderAccessRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeauthorizeCloudProviderAccessRole'
type CloudProviderAccessApi_DeauthorizeCloudProviderAccessRole_Call struct {
	*mock.Call
}

// DeauthorizeCloudProviderAccessRole is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - cloudProvider string
//   - roleId string
func (_e *CloudProviderAccessApi_Expecter) DeauthorizeCloudProviderAccessRole(ctx interface{}, groupId interface{}, cloudProvider interface{}, roleId interface{}) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRole_Call {
	return &CloudProviderAccessApi_DeauthorizeCloudProviderAccessRole_Call{Call: _e.mock.On("DeauthorizeCloudProviderAccessRole", ctx, groupId, cloudProvider, roleId)}
}

func (_c *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRole_Call) Run(run func(ctx context.Context, groupId string, cloudProvider string, roleId string)) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRole_Call) Return(_a0 admin.DeauthorizeCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRole_Call) RunAndReturn(run func(context.Context, string, string, string) admin.DeauthorizeCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRole_Call {
	_c.Call.Return(run)
	return _c
}

// DeauthorizeCloudProviderAccessRoleExecute provides a mock function with given fields: r
func (_m *CloudProviderAccessApi) DeauthorizeCloudProviderAccessRoleExecute(r admin.DeauthorizeCloudProviderAccessRoleApiRequest) (*http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DeauthorizeCloudProviderAccessRoleExecute")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(admin.DeauthorizeCloudProviderAccessRoleApiRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DeauthorizeCloudProviderAccessRoleApiRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DeauthorizeCloudProviderAccessRoleApiRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeauthorizeCloudProviderAccessRoleExecute'
type CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleExecute_Call struct {
	*mock.Call
}

// DeauthorizeCloudProviderAccessRoleExecute is a helper method to define mock.On call
//   - r admin.DeauthorizeCloudProviderAccessRoleApiRequest
func (_e *CloudProviderAccessApi_Expecter) DeauthorizeCloudProviderAccessRoleExecute(r interface{}) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleExecute_Call {
	return &CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleExecute_Call{Call: _e.mock.On("DeauthorizeCloudProviderAccessRoleExecute", r)}
}

func (_c *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleExecute_Call) Run(run func(r admin.DeauthorizeCloudProviderAccessRoleApiRequest)) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DeauthorizeCloudProviderAccessRoleApiRequest))
	})
	return _c
}

func (_c *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleExecute_Call) Return(_a0 *http.Response, _a1 error) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleExecute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleExecute_Call) RunAndReturn(run func(admin.DeauthorizeCloudProviderAccessRoleApiRequest) (*http.Response, error)) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DeauthorizeCloudProviderAccessRoleWithParams provides a mock function with given fields: ctx, args
func (_m *CloudProviderAccessApi) DeauthorizeCloudProviderAccessRoleWithParams(ctx context.Context, args *admin.DeauthorizeCloudProviderAccessRoleApiParams) admin.DeauthorizeCloudProviderAccessRoleApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DeauthorizeCloudProviderAccessRoleWithParams")
	}

	var r0 admin.DeauthorizeCloudProviderAccessRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeauthorizeCloudProviderAccessRoleApiParams) admin.DeauthorizeCloudProviderAccessRoleApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DeauthorizeCloudProviderAccessRoleApiRequest)
	}

	return r0
}

// CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeauthorizeCloudProviderAccessRoleWithParams'
type CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleWithParams_Call struct {
	*mock.Call
}

// DeauthorizeCloudProviderAccessRoleWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DeauthorizeCloudProviderAccessRoleApiParams
func (_e *CloudProviderAccessApi_Expecter) DeauthorizeCloudProviderAccessRoleWithParams(ctx interface{}, args interface{}) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleWithParams_Call {
	return &CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleWithParams_Call{Call: _e.mock.On("DeauthorizeCloudProviderAccessRoleWithParams", ctx, args)}
}

func (_c *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleWithParams_Call) Run(run func(ctx context.Context, args *admin.DeauthorizeCloudProviderAccessRoleApiParams)) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DeauthorizeCloudProviderAccessRoleApiParams))
	})
	return _c
}

func (_c *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleWithParams_Call) Return(_a0 admin.DeauthorizeCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleWithParams_Call) RunAndReturn(run func(context.Context, *admin.DeauthorizeCloudProviderAccessRoleApiParams) admin.DeauthorizeCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_DeauthorizeCloudProviderAccessRoleWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetCloudProviderAccessRole provides a mock function with given fields: ctx, groupId, roleId
func (_m *CloudProviderAccessApi) GetCloudProviderAccessRole(ctx context.Context, groupId string, roleId string) admin.GetCloudProviderAccessRoleApiRequest {
	ret := _m.Called(ctx, groupId, roleId)

	if len(ret) == 0 {
		panic("no return value specified for GetCloudProviderAccessRole")
	}

	var r0 admin.GetCloudProviderAccessRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetCloudProviderAccessRoleApiRequest); ok {
		r0 = rf(ctx, groupId, roleId)
	} else {
		r0 = ret.Get(0).(admin.GetCloudProviderAccessRoleApiRequest)
	}

	return r0
}

// CloudProviderAccessApi_GetCloudProviderAccessRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCloudProviderAccessRole'
type CloudProviderAccessApi_GetCloudProviderAccessRole_Call struct {
	*mock.Call
}

// GetCloudProviderAccessRole is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - roleId string
func (_e *CloudProviderAccessApi_Expecter) GetCloudProviderAccessRole(ctx interface{}, groupId interface{}, roleId interface{}) *CloudProviderAccessApi_GetCloudProviderAccessRole_Call {
	return &CloudProviderAccessApi_GetCloudProviderAccessRole_Call{Call: _e.mock.On("GetCloudProviderAccessRole", ctx, groupId, roleId)}
}

func (_c *CloudProviderAccessApi_GetCloudProviderAccessRole_Call) Run(run func(ctx context.Context, groupId string, roleId string)) *CloudProviderAccessApi_GetCloudProviderAccessRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *CloudProviderAccessApi_GetCloudProviderAccessRole_Call) Return(_a0 admin.GetCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_GetCloudProviderAccessRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloudProviderAccessApi_GetCloudProviderAccessRole_Call) RunAndReturn(run func(context.Context, string, string) admin.GetCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_GetCloudProviderAccessRole_Call {
	_c.Call.Return(run)
	return _c
}

// GetCloudProviderAccessRoleExecute provides a mock function with given fields: r
func (_m *CloudProviderAccessApi) GetCloudProviderAccessRoleExecute(r admin.GetCloudProviderAccessRoleApiRequest) (*admin.CloudProviderAccessRole, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetCloudProviderAccessRoleExecute")
	}

	var r0 *admin.CloudProviderAccessRole
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetCloudProviderAccessRoleApiRequest) (*admin.CloudProviderAccessRole, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetCloudProviderAccessRoleApiRequest) *admin.CloudProviderAccessRole); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CloudProviderAccessRole)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetCloudProviderAccessRoleApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetCloudProviderAccessRoleApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CloudProviderAccessApi_GetCloudProviderAccessRoleExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCloudProviderAccessRoleExecute'
type CloudProviderAccessApi_GetCloudProviderAccessRoleExecute_Call struct {
	*mock.Call
}

// GetCloudProviderAccessRoleExecute is a helper method to define mock.On call
//   - r admin.GetCloudProviderAccessRoleApiRequest
func (_e *CloudProviderAccessApi_Expecter) GetCloudProviderAccessRoleExecute(r interface{}) *CloudProviderAccessApi_GetCloudProviderAccessRoleExecute_Call {
	return &CloudProviderAccessApi_GetCloudProviderAccessRoleExecute_Call{Call: _e.mock.On("GetCloudProviderAccessRoleExecute", r)}
}

func (_c *CloudProviderAccessApi_GetCloudProviderAccessRoleExecute_Call) Run(run func(r admin.GetCloudProviderAccessRoleApiRequest)) *CloudProviderAccessApi_GetCloudProviderAccessRoleExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetCloudProviderAccessRoleApiRequest))
	})
	return _c
}

func (_c *CloudProviderAccessApi_GetCloudProviderAccessRoleExecute_Call) Return(_a0 *admin.CloudProviderAccessRole, _a1 *http.Response, _a2 error) *CloudProviderAccessApi_GetCloudProviderAccessRoleExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CloudProviderAccessApi_GetCloudProviderAccessRoleExecute_Call) RunAndReturn(run func(admin.GetCloudProviderAccessRoleApiRequest) (*admin.CloudProviderAccessRole, *http.Response, error)) *CloudProviderAccessApi_GetCloudProviderAccessRoleExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetCloudProviderAccessRoleWithParams provides a mock function with given fields: ctx, args
func (_m *CloudProviderAccessApi) GetCloudProviderAccessRoleWithParams(ctx context.Context, args *admin.GetCloudProviderAccessRoleApiParams) admin.GetCloudProviderAccessRoleApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetCloudProviderAccessRoleWithParams")
	}

	var r0 admin.GetCloudProviderAccessRoleApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetCloudProviderAccessRoleApiParams) admin.GetCloudProviderAccessRoleApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetCloudProviderAccessRoleApiRequest)
	}

	return r0
}

// CloudProviderAccessApi_GetCloudProviderAccessRoleWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCloudProviderAccessRoleWithParams'
type CloudProviderAccessApi_GetCloudProviderAccessRoleWithParams_Call struct {
	*mock.Call
}

// GetCloudProviderAccessRoleWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetCloudProviderAccessRoleApiParams
func (_e *CloudProviderAccessApi_Expecter) GetCloudProviderAccessRoleWithParams(ctx interface{}, args interface{}) *CloudProviderAccessApi_GetCloudProviderAccessRoleWithParams_Call {
	return &CloudProviderAccessApi_GetCloudProviderAccessRoleWithParams_Call{Call: _e.mock.On("GetCloudProviderAccessRoleWithParams", ctx, args)}
}

func (_c *CloudProviderAccessApi_GetCloudProviderAccessRoleWithParams_Call) Run(run func(ctx context.Context, args *admin.GetCloudProviderAccessRoleApiParams)) *CloudProviderAccessApi_GetCloudProviderAccessRoleWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetCloudProviderAccessRoleApiParams))
	})
	return _c
}

func (_c *CloudProviderAccessApi_GetCloudProviderAccessRoleWithParams_Call) Return(_a0 admin.GetCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_GetCloudProviderAccessRoleWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloudProviderAccessApi_GetCloudProviderAccessRoleWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetCloudProviderAccessRoleApiParams) admin.GetCloudProviderAccessRoleApiRequest) *CloudProviderAccessApi_GetCloudProviderAccessRoleWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListCloudProviderAccessRoles provides a mock function with given fields: ctx, groupId
func (_m *CloudProviderAccessApi) ListCloudProviderAccessRoles(ctx context.Context, groupId string) admin.ListCloudProviderAccessRolesApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for ListCloudProviderAccessRoles")
	}

	var r0 admin.ListCloudProviderAccessRolesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ListCloudProviderAccessRolesApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.ListCloudProviderAccessRolesApiRequest)
	}

	return r0
}

// CloudProviderAccessApi_ListCloudProviderAccessRoles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCloudProviderAccessRoles'
type CloudProviderAccessApi_ListCloudProviderAccessRoles_Call struct {
	*mock.Call
}

// ListCloudProviderAccessRoles is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *CloudProviderAccessApi_Expecter) ListCloudProviderAccessRoles(ctx interface{}, groupId interface{}) *CloudProviderAccessApi_ListCloudProviderAccessRoles_Call {
	return &CloudProviderAccessApi_ListCloudProviderAccessRoles_Call{Call: _e.mock.On("ListCloudProviderAccessRoles", ctx, groupId)}
}

func (_c *CloudProviderAccessApi_ListCloudProviderAccessRoles_Call) Run(run func(ctx context.Context, groupId string)) *CloudProviderAccessApi_ListCloudProviderAccessRoles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CloudProviderAccessApi_ListCloudProviderAccessRoles_Call) Return(_a0 admin.ListCloudProviderAccessRolesApiRequest) *CloudProviderAccessApi_ListCloudProviderAccessRoles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloudProviderAccessApi_ListCloudProviderAccessRoles_Call) RunAndReturn(run func(context.Context, string) admin.ListCloudProviderAccessRolesApiRequest) *CloudProviderAccessApi_ListCloudProviderAccessRoles_Call {
	_c.Call.Return(run)
	return _c
}

// ListCloudProviderAccessRolesExecute provides a mock function with given fields: r
func (_m *CloudProviderAccessApi) ListCloudProviderAccessRolesExecute(r admin.ListCloudProviderAccessRolesApiRequest) (*admin.CloudProviderAccessRoles, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListCloudProviderAccessRolesExecute")
	}

	var r0 *admin.CloudProviderAccessRoles
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListCloudProviderAccessRolesApiRequest) (*admin.CloudProviderAccessRoles, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListCloudProviderAccessRolesApiRequest) *admin.CloudProviderAccessRoles); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CloudProviderAccessRoles)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListCloudProviderAccessRolesApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListCloudProviderAccessRolesApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CloudProviderAccessApi_ListCloudProviderAccessRolesExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCloudProviderAccessRolesExecute'
type CloudProviderAccessApi_ListCloudProviderAccessRolesExecute_Call struct {
	*mock.Call
}

// ListCloudProviderAccessRolesExecute is a helper method to define mock.On call
//   - r admin.ListCloudProviderAccessRolesApiRequest
func (_e *CloudProviderAccessApi_Expecter) ListCloudProviderAccessRolesExecute(r interface{}) *CloudProviderAccessApi_ListCloudProviderAccessRolesExecute_Call {
	return &CloudProviderAccessApi_ListCloudProviderAccessRolesExecute_Call{Call: _e.mock.On("ListCloudProviderAccessRolesExecute", r)}
}

func (_c *CloudProviderAccessApi_ListCloudProviderAccessRolesExecute_Call) Run(run func(r admin.ListCloudProviderAccessRolesApiRequest)) *CloudProviderAccessApi_ListCloudProviderAccessRolesExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListCloudProviderAccessRolesApiRequest))
	})
	return _c
}

func (_c *CloudProviderAccessApi_ListCloudProviderAccessRolesExecute_Call) Return(_a0 *admin.CloudProviderAccessRoles, _a1 *http.Response, _a2 error) *CloudProviderAccessApi_ListCloudProviderAccessRolesExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CloudProviderAccessApi_ListCloudProviderAccessRolesExecute_Call) RunAndReturn(run func(admin.ListCloudProviderAccessRolesApiRequest) (*admin.CloudProviderAccessRoles, *http.Response, error)) *CloudProviderAccessApi_ListCloudProviderAccessRolesExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListCloudProviderAccessRolesWithParams provides a mock function with given fields: ctx, args
func (_m *CloudProviderAccessApi) ListCloudProviderAccessRolesWithParams(ctx context.Context, args *admin.ListCloudProviderAccessRolesApiParams) admin.ListCloudProviderAccessRolesApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListCloudProviderAccessRolesWithParams")
	}

	var r0 admin.ListCloudProviderAccessRolesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListCloudProviderAccessRolesApiParams) admin.ListCloudProviderAccessRolesApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListCloudProviderAccessRolesApiRequest)
	}

	return r0
}

// CloudProviderAccessApi_ListCloudProviderAccessRolesWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCloudProviderAccessRolesWithParams'
type CloudProviderAccessApi_ListCloudProviderAccessRolesWithParams_Call struct {
	*mock.Call
}

// ListCloudProviderAccessRolesWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListCloudProviderAccessRolesApiParams
func (_e *CloudProviderAccessApi_Expecter) ListCloudProviderAccessRolesWithParams(ctx interface{}, args interface{}) *CloudProviderAccessApi_ListCloudProviderAccessRolesWithParams_Call {
	return &CloudProviderAccessApi_ListCloudProviderAccessRolesWithParams_Call{Call: _e.mock.On("ListCloudProviderAccessRolesWithParams", ctx, args)}
}

func (_c *CloudProviderAccessApi_ListCloudProviderAccessRolesWithParams_Call) Run(run func(ctx context.Context, args *admin.ListCloudProviderAccessRolesApiParams)) *CloudProviderAccessApi_ListCloudProviderAccessRolesWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListCloudProviderAccessRolesApiParams))
	})
	return _c
}

func (_c *CloudProviderAccessApi_ListCloudProviderAccessRolesWithParams_Call) Return(_a0 admin.ListCloudProviderAccessRolesApiRequest) *CloudProviderAccessApi_ListCloudProviderAccessRolesWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloudProviderAccessApi_ListCloudProviderAccessRolesWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListCloudProviderAccessRolesApiParams) admin.ListCloudProviderAccessRolesApiRequest) *CloudProviderAccessApi_ListCloudProviderAccessRolesWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewCloudProviderAccessApi creates a new instance of CloudProviderAccessApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCloudProviderAccessApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *CloudProviderAccessApi {
	mock := &CloudProviderAccessApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}