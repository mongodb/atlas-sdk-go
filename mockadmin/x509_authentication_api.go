// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "github.com/mongodb/atlas-sdk-go/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// X509AuthenticationApi is an autogenerated mock type for the X509AuthenticationApi type
type X509AuthenticationApi struct {
	mock.Mock
}

type X509AuthenticationApi_Expecter struct {
	mock *mock.Mock
}

func (_m *X509AuthenticationApi) EXPECT() *X509AuthenticationApi_Expecter {
	return &X509AuthenticationApi_Expecter{mock: &_m.Mock}
}

// CreateDatabaseUserCertificate provides a mock function with given fields: ctx, groupId, username, userCert
func (_m *X509AuthenticationApi) CreateDatabaseUserCertificate(ctx context.Context, groupId string, username string, userCert *admin.UserCert) admin.CreateDatabaseUserCertificateApiRequest {
	ret := _m.Called(ctx, groupId, username, userCert)

	if len(ret) == 0 {
		panic("no return value specified for CreateDatabaseUserCertificate")
	}

	var r0 admin.CreateDatabaseUserCertificateApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.UserCert) admin.CreateDatabaseUserCertificateApiRequest); ok {
		r0 = rf(ctx, groupId, username, userCert)
	} else {
		r0 = ret.Get(0).(admin.CreateDatabaseUserCertificateApiRequest)
	}

	return r0
}

// X509AuthenticationApi_CreateDatabaseUserCertificate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDatabaseUserCertificate'
type X509AuthenticationApi_CreateDatabaseUserCertificate_Call struct {
	*mock.Call
}

// CreateDatabaseUserCertificate is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - username string
//   - userCert *admin.UserCert
func (_e *X509AuthenticationApi_Expecter) CreateDatabaseUserCertificate(ctx any, groupId any, username any, userCert any) *X509AuthenticationApi_CreateDatabaseUserCertificate_Call {
	return &X509AuthenticationApi_CreateDatabaseUserCertificate_Call{Call: _e.mock.On("CreateDatabaseUserCertificate", ctx, groupId, username, userCert)}
}

func (_c *X509AuthenticationApi_CreateDatabaseUserCertificate_Call) Run(run func(ctx context.Context, groupId string, username string, userCert *admin.UserCert)) *X509AuthenticationApi_CreateDatabaseUserCertificate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.UserCert))
	})
	return _c
}

func (_c *X509AuthenticationApi_CreateDatabaseUserCertificate_Call) Return(_a0 admin.CreateDatabaseUserCertificateApiRequest) *X509AuthenticationApi_CreateDatabaseUserCertificate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *X509AuthenticationApi_CreateDatabaseUserCertificate_Call) RunAndReturn(run func(context.Context, string, string, *admin.UserCert) admin.CreateDatabaseUserCertificateApiRequest) *X509AuthenticationApi_CreateDatabaseUserCertificate_Call {
	_c.Call.Return(run)
	return _c
}

// CreateDatabaseUserCertificateExecute provides a mock function with given fields: r
func (_m *X509AuthenticationApi) CreateDatabaseUserCertificateExecute(r admin.CreateDatabaseUserCertificateApiRequest) (string, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateDatabaseUserCertificateExecute")
	}

	var r0 string
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateDatabaseUserCertificateApiRequest) (string, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateDatabaseUserCertificateApiRequest) string); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(admin.CreateDatabaseUserCertificateApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateDatabaseUserCertificateApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// X509AuthenticationApi_CreateDatabaseUserCertificateExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDatabaseUserCertificateExecute'
type X509AuthenticationApi_CreateDatabaseUserCertificateExecute_Call struct {
	*mock.Call
}

// CreateDatabaseUserCertificateExecute is a helper method to define mock.On call
//   - r admin.CreateDatabaseUserCertificateApiRequest
func (_e *X509AuthenticationApi_Expecter) CreateDatabaseUserCertificateExecute(r any) *X509AuthenticationApi_CreateDatabaseUserCertificateExecute_Call {
	return &X509AuthenticationApi_CreateDatabaseUserCertificateExecute_Call{Call: _e.mock.On("CreateDatabaseUserCertificateExecute", r)}
}

func (_c *X509AuthenticationApi_CreateDatabaseUserCertificateExecute_Call) Run(run func(r admin.CreateDatabaseUserCertificateApiRequest)) *X509AuthenticationApi_CreateDatabaseUserCertificateExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateDatabaseUserCertificateApiRequest))
	})
	return _c
}

func (_c *X509AuthenticationApi_CreateDatabaseUserCertificateExecute_Call) Return(_a0 string, _a1 *http.Response, _a2 error) *X509AuthenticationApi_CreateDatabaseUserCertificateExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *X509AuthenticationApi_CreateDatabaseUserCertificateExecute_Call) RunAndReturn(run func(admin.CreateDatabaseUserCertificateApiRequest) (string, *http.Response, error)) *X509AuthenticationApi_CreateDatabaseUserCertificateExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateDatabaseUserCertificateWithParams provides a mock function with given fields: ctx, args
func (_m *X509AuthenticationApi) CreateDatabaseUserCertificateWithParams(ctx context.Context, args *admin.CreateDatabaseUserCertificateApiParams) admin.CreateDatabaseUserCertificateApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateDatabaseUserCertificateWithParams")
	}

	var r0 admin.CreateDatabaseUserCertificateApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateDatabaseUserCertificateApiParams) admin.CreateDatabaseUserCertificateApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateDatabaseUserCertificateApiRequest)
	}

	return r0
}

// X509AuthenticationApi_CreateDatabaseUserCertificateWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDatabaseUserCertificateWithParams'
type X509AuthenticationApi_CreateDatabaseUserCertificateWithParams_Call struct {
	*mock.Call
}

// CreateDatabaseUserCertificateWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateDatabaseUserCertificateApiParams
func (_e *X509AuthenticationApi_Expecter) CreateDatabaseUserCertificateWithParams(ctx any, args any) *X509AuthenticationApi_CreateDatabaseUserCertificateWithParams_Call {
	return &X509AuthenticationApi_CreateDatabaseUserCertificateWithParams_Call{Call: _e.mock.On("CreateDatabaseUserCertificateWithParams", ctx, args)}
}

func (_c *X509AuthenticationApi_CreateDatabaseUserCertificateWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateDatabaseUserCertificateApiParams)) *X509AuthenticationApi_CreateDatabaseUserCertificateWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateDatabaseUserCertificateApiParams))
	})
	return _c
}

func (_c *X509AuthenticationApi_CreateDatabaseUserCertificateWithParams_Call) Return(_a0 admin.CreateDatabaseUserCertificateApiRequest) *X509AuthenticationApi_CreateDatabaseUserCertificateWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *X509AuthenticationApi_CreateDatabaseUserCertificateWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateDatabaseUserCertificateApiParams) admin.CreateDatabaseUserCertificateApiRequest) *X509AuthenticationApi_CreateDatabaseUserCertificateWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// DisableCustomerManagedX509 provides a mock function with given fields: ctx, groupId
func (_m *X509AuthenticationApi) DisableCustomerManagedX509(ctx context.Context, groupId string) admin.DisableCustomerManagedX509ApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for DisableCustomerManagedX509")
	}

	var r0 admin.DisableCustomerManagedX509ApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.DisableCustomerManagedX509ApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.DisableCustomerManagedX509ApiRequest)
	}

	return r0
}

// X509AuthenticationApi_DisableCustomerManagedX509_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisableCustomerManagedX509'
type X509AuthenticationApi_DisableCustomerManagedX509_Call struct {
	*mock.Call
}

// DisableCustomerManagedX509 is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *X509AuthenticationApi_Expecter) DisableCustomerManagedX509(ctx any, groupId any) *X509AuthenticationApi_DisableCustomerManagedX509_Call {
	return &X509AuthenticationApi_DisableCustomerManagedX509_Call{Call: _e.mock.On("DisableCustomerManagedX509", ctx, groupId)}
}

func (_c *X509AuthenticationApi_DisableCustomerManagedX509_Call) Run(run func(ctx context.Context, groupId string)) *X509AuthenticationApi_DisableCustomerManagedX509_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *X509AuthenticationApi_DisableCustomerManagedX509_Call) Return(_a0 admin.DisableCustomerManagedX509ApiRequest) *X509AuthenticationApi_DisableCustomerManagedX509_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *X509AuthenticationApi_DisableCustomerManagedX509_Call) RunAndReturn(run func(context.Context, string) admin.DisableCustomerManagedX509ApiRequest) *X509AuthenticationApi_DisableCustomerManagedX509_Call {
	_c.Call.Return(run)
	return _c
}

// DisableCustomerManagedX509Execute provides a mock function with given fields: r
func (_m *X509AuthenticationApi) DisableCustomerManagedX509Execute(r admin.DisableCustomerManagedX509ApiRequest) (*admin.UserSecurity, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DisableCustomerManagedX509Execute")
	}

	var r0 *admin.UserSecurity
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.DisableCustomerManagedX509ApiRequest) (*admin.UserSecurity, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DisableCustomerManagedX509ApiRequest) *admin.UserSecurity); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.UserSecurity)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DisableCustomerManagedX509ApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.DisableCustomerManagedX509ApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// X509AuthenticationApi_DisableCustomerManagedX509Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisableCustomerManagedX509Execute'
type X509AuthenticationApi_DisableCustomerManagedX509Execute_Call struct {
	*mock.Call
}

// DisableCustomerManagedX509Execute is a helper method to define mock.On call
//   - r admin.DisableCustomerManagedX509ApiRequest
func (_e *X509AuthenticationApi_Expecter) DisableCustomerManagedX509Execute(r any) *X509AuthenticationApi_DisableCustomerManagedX509Execute_Call {
	return &X509AuthenticationApi_DisableCustomerManagedX509Execute_Call{Call: _e.mock.On("DisableCustomerManagedX509Execute", r)}
}

func (_c *X509AuthenticationApi_DisableCustomerManagedX509Execute_Call) Run(run func(r admin.DisableCustomerManagedX509ApiRequest)) *X509AuthenticationApi_DisableCustomerManagedX509Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DisableCustomerManagedX509ApiRequest))
	})
	return _c
}

func (_c *X509AuthenticationApi_DisableCustomerManagedX509Execute_Call) Return(_a0 *admin.UserSecurity, _a1 *http.Response, _a2 error) *X509AuthenticationApi_DisableCustomerManagedX509Execute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *X509AuthenticationApi_DisableCustomerManagedX509Execute_Call) RunAndReturn(run func(admin.DisableCustomerManagedX509ApiRequest) (*admin.UserSecurity, *http.Response, error)) *X509AuthenticationApi_DisableCustomerManagedX509Execute_Call {
	_c.Call.Return(run)
	return _c
}

// DisableCustomerManagedX509WithParams provides a mock function with given fields: ctx, args
func (_m *X509AuthenticationApi) DisableCustomerManagedX509WithParams(ctx context.Context, args *admin.DisableCustomerManagedX509ApiParams) admin.DisableCustomerManagedX509ApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DisableCustomerManagedX509WithParams")
	}

	var r0 admin.DisableCustomerManagedX509ApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DisableCustomerManagedX509ApiParams) admin.DisableCustomerManagedX509ApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DisableCustomerManagedX509ApiRequest)
	}

	return r0
}

// X509AuthenticationApi_DisableCustomerManagedX509WithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisableCustomerManagedX509WithParams'
type X509AuthenticationApi_DisableCustomerManagedX509WithParams_Call struct {
	*mock.Call
}

// DisableCustomerManagedX509WithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DisableCustomerManagedX509ApiParams
func (_e *X509AuthenticationApi_Expecter) DisableCustomerManagedX509WithParams(ctx any, args any) *X509AuthenticationApi_DisableCustomerManagedX509WithParams_Call {
	return &X509AuthenticationApi_DisableCustomerManagedX509WithParams_Call{Call: _e.mock.On("DisableCustomerManagedX509WithParams", ctx, args)}
}

func (_c *X509AuthenticationApi_DisableCustomerManagedX509WithParams_Call) Run(run func(ctx context.Context, args *admin.DisableCustomerManagedX509ApiParams)) *X509AuthenticationApi_DisableCustomerManagedX509WithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DisableCustomerManagedX509ApiParams))
	})
	return _c
}

func (_c *X509AuthenticationApi_DisableCustomerManagedX509WithParams_Call) Return(_a0 admin.DisableCustomerManagedX509ApiRequest) *X509AuthenticationApi_DisableCustomerManagedX509WithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *X509AuthenticationApi_DisableCustomerManagedX509WithParams_Call) RunAndReturn(run func(context.Context, *admin.DisableCustomerManagedX509ApiParams) admin.DisableCustomerManagedX509ApiRequest) *X509AuthenticationApi_DisableCustomerManagedX509WithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListDatabaseUserCertificates provides a mock function with given fields: ctx, groupId, username
func (_m *X509AuthenticationApi) ListDatabaseUserCertificates(ctx context.Context, groupId string, username string) admin.ListDatabaseUserCertificatesApiRequest {
	ret := _m.Called(ctx, groupId, username)

	if len(ret) == 0 {
		panic("no return value specified for ListDatabaseUserCertificates")
	}

	var r0 admin.ListDatabaseUserCertificatesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.ListDatabaseUserCertificatesApiRequest); ok {
		r0 = rf(ctx, groupId, username)
	} else {
		r0 = ret.Get(0).(admin.ListDatabaseUserCertificatesApiRequest)
	}

	return r0
}

// X509AuthenticationApi_ListDatabaseUserCertificates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDatabaseUserCertificates'
type X509AuthenticationApi_ListDatabaseUserCertificates_Call struct {
	*mock.Call
}

// ListDatabaseUserCertificates is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - username string
func (_e *X509AuthenticationApi_Expecter) ListDatabaseUserCertificates(ctx any, groupId any, username any) *X509AuthenticationApi_ListDatabaseUserCertificates_Call {
	return &X509AuthenticationApi_ListDatabaseUserCertificates_Call{Call: _e.mock.On("ListDatabaseUserCertificates", ctx, groupId, username)}
}

func (_c *X509AuthenticationApi_ListDatabaseUserCertificates_Call) Run(run func(ctx context.Context, groupId string, username string)) *X509AuthenticationApi_ListDatabaseUserCertificates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *X509AuthenticationApi_ListDatabaseUserCertificates_Call) Return(_a0 admin.ListDatabaseUserCertificatesApiRequest) *X509AuthenticationApi_ListDatabaseUserCertificates_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *X509AuthenticationApi_ListDatabaseUserCertificates_Call) RunAndReturn(run func(context.Context, string, string) admin.ListDatabaseUserCertificatesApiRequest) *X509AuthenticationApi_ListDatabaseUserCertificates_Call {
	_c.Call.Return(run)
	return _c
}

// ListDatabaseUserCertificatesExecute provides a mock function with given fields: r
func (_m *X509AuthenticationApi) ListDatabaseUserCertificatesExecute(r admin.ListDatabaseUserCertificatesApiRequest) (*admin.PaginatedUserCert, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListDatabaseUserCertificatesExecute")
	}

	var r0 *admin.PaginatedUserCert
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListDatabaseUserCertificatesApiRequest) (*admin.PaginatedUserCert, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListDatabaseUserCertificatesApiRequest) *admin.PaginatedUserCert); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedUserCert)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListDatabaseUserCertificatesApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListDatabaseUserCertificatesApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// X509AuthenticationApi_ListDatabaseUserCertificatesExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDatabaseUserCertificatesExecute'
type X509AuthenticationApi_ListDatabaseUserCertificatesExecute_Call struct {
	*mock.Call
}

// ListDatabaseUserCertificatesExecute is a helper method to define mock.On call
//   - r admin.ListDatabaseUserCertificatesApiRequest
func (_e *X509AuthenticationApi_Expecter) ListDatabaseUserCertificatesExecute(r any) *X509AuthenticationApi_ListDatabaseUserCertificatesExecute_Call {
	return &X509AuthenticationApi_ListDatabaseUserCertificatesExecute_Call{Call: _e.mock.On("ListDatabaseUserCertificatesExecute", r)}
}

func (_c *X509AuthenticationApi_ListDatabaseUserCertificatesExecute_Call) Run(run func(r admin.ListDatabaseUserCertificatesApiRequest)) *X509AuthenticationApi_ListDatabaseUserCertificatesExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListDatabaseUserCertificatesApiRequest))
	})
	return _c
}

func (_c *X509AuthenticationApi_ListDatabaseUserCertificatesExecute_Call) Return(_a0 *admin.PaginatedUserCert, _a1 *http.Response, _a2 error) *X509AuthenticationApi_ListDatabaseUserCertificatesExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *X509AuthenticationApi_ListDatabaseUserCertificatesExecute_Call) RunAndReturn(run func(admin.ListDatabaseUserCertificatesApiRequest) (*admin.PaginatedUserCert, *http.Response, error)) *X509AuthenticationApi_ListDatabaseUserCertificatesExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListDatabaseUserCertificatesWithParams provides a mock function with given fields: ctx, args
func (_m *X509AuthenticationApi) ListDatabaseUserCertificatesWithParams(ctx context.Context, args *admin.ListDatabaseUserCertificatesApiParams) admin.ListDatabaseUserCertificatesApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListDatabaseUserCertificatesWithParams")
	}

	var r0 admin.ListDatabaseUserCertificatesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListDatabaseUserCertificatesApiParams) admin.ListDatabaseUserCertificatesApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListDatabaseUserCertificatesApiRequest)
	}

	return r0
}

// X509AuthenticationApi_ListDatabaseUserCertificatesWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDatabaseUserCertificatesWithParams'
type X509AuthenticationApi_ListDatabaseUserCertificatesWithParams_Call struct {
	*mock.Call
}

// ListDatabaseUserCertificatesWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListDatabaseUserCertificatesApiParams
func (_e *X509AuthenticationApi_Expecter) ListDatabaseUserCertificatesWithParams(ctx any, args any) *X509AuthenticationApi_ListDatabaseUserCertificatesWithParams_Call {
	return &X509AuthenticationApi_ListDatabaseUserCertificatesWithParams_Call{Call: _e.mock.On("ListDatabaseUserCertificatesWithParams", ctx, args)}
}

func (_c *X509AuthenticationApi_ListDatabaseUserCertificatesWithParams_Call) Run(run func(ctx context.Context, args *admin.ListDatabaseUserCertificatesApiParams)) *X509AuthenticationApi_ListDatabaseUserCertificatesWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListDatabaseUserCertificatesApiParams))
	})
	return _c
}

func (_c *X509AuthenticationApi_ListDatabaseUserCertificatesWithParams_Call) Return(_a0 admin.ListDatabaseUserCertificatesApiRequest) *X509AuthenticationApi_ListDatabaseUserCertificatesWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *X509AuthenticationApi_ListDatabaseUserCertificatesWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListDatabaseUserCertificatesApiParams) admin.ListDatabaseUserCertificatesApiRequest) *X509AuthenticationApi_ListDatabaseUserCertificatesWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewX509AuthenticationApi creates a new instance of X509AuthenticationApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewX509AuthenticationApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *X509AuthenticationApi {
	mock := &X509AuthenticationApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
