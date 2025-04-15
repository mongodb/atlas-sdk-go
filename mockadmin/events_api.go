// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20231115014/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// EventsApi is an autogenerated mock type for the EventsApi type
type EventsApi struct {
	mock.Mock
}

type EventsApi_Expecter struct {
	mock *mock.Mock
}

func (_m *EventsApi) EXPECT() *EventsApi_Expecter {
	return &EventsApi_Expecter{mock: &_m.Mock}
}

// GetOrganizationEvent provides a mock function with given fields: ctx, orgId, eventId
func (_m *EventsApi) GetOrganizationEvent(ctx context.Context, orgId string, eventId string) admin.GetOrganizationEventApiRequest {
	ret := _m.Called(ctx, orgId, eventId)

	if len(ret) == 0 {
		panic("no return value specified for GetOrganizationEvent")
	}

	var r0 admin.GetOrganizationEventApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetOrganizationEventApiRequest); ok {
		r0 = rf(ctx, orgId, eventId)
	} else {
		r0 = ret.Get(0).(admin.GetOrganizationEventApiRequest)
	}

	return r0
}

// EventsApi_GetOrganizationEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrganizationEvent'
type EventsApi_GetOrganizationEvent_Call struct {
	*mock.Call
}

// GetOrganizationEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - orgId string
//   - eventId string
func (_e *EventsApi_Expecter) GetOrganizationEvent(ctx any, orgId any, eventId any) *EventsApi_GetOrganizationEvent_Call {
	return &EventsApi_GetOrganizationEvent_Call{Call: _e.mock.On("GetOrganizationEvent", ctx, orgId, eventId)}
}

func (_c *EventsApi_GetOrganizationEvent_Call) Run(run func(ctx context.Context, orgId string, eventId string)) *EventsApi_GetOrganizationEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *EventsApi_GetOrganizationEvent_Call) Return(_a0 admin.GetOrganizationEventApiRequest) *EventsApi_GetOrganizationEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsApi_GetOrganizationEvent_Call) RunAndReturn(run func(context.Context, string, string) admin.GetOrganizationEventApiRequest) *EventsApi_GetOrganizationEvent_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrganizationEventExecute provides a mock function with given fields: r
func (_m *EventsApi) GetOrganizationEventExecute(r admin.GetOrganizationEventApiRequest) (*admin.EventViewForOrg, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetOrganizationEventExecute")
	}

	var r0 *admin.EventViewForOrg
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetOrganizationEventApiRequest) (*admin.EventViewForOrg, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetOrganizationEventApiRequest) *admin.EventViewForOrg); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.EventViewForOrg)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetOrganizationEventApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetOrganizationEventApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventsApi_GetOrganizationEventExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrganizationEventExecute'
type EventsApi_GetOrganizationEventExecute_Call struct {
	*mock.Call
}

// GetOrganizationEventExecute is a helper method to define mock.On call
//   - r admin.GetOrganizationEventApiRequest
func (_e *EventsApi_Expecter) GetOrganizationEventExecute(r any) *EventsApi_GetOrganizationEventExecute_Call {
	return &EventsApi_GetOrganizationEventExecute_Call{Call: _e.mock.On("GetOrganizationEventExecute", r)}
}

func (_c *EventsApi_GetOrganizationEventExecute_Call) Run(run func(r admin.GetOrganizationEventApiRequest)) *EventsApi_GetOrganizationEventExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetOrganizationEventApiRequest))
	})
	return _c
}

func (_c *EventsApi_GetOrganizationEventExecute_Call) Return(_a0 *admin.EventViewForOrg, _a1 *http.Response, _a2 error) *EventsApi_GetOrganizationEventExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EventsApi_GetOrganizationEventExecute_Call) RunAndReturn(run func(admin.GetOrganizationEventApiRequest) (*admin.EventViewForOrg, *http.Response, error)) *EventsApi_GetOrganizationEventExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrganizationEventWithParams provides a mock function with given fields: ctx, args
func (_m *EventsApi) GetOrganizationEventWithParams(ctx context.Context, args *admin.GetOrganizationEventApiParams) admin.GetOrganizationEventApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetOrganizationEventWithParams")
	}

	var r0 admin.GetOrganizationEventApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetOrganizationEventApiParams) admin.GetOrganizationEventApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetOrganizationEventApiRequest)
	}

	return r0
}

// EventsApi_GetOrganizationEventWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrganizationEventWithParams'
type EventsApi_GetOrganizationEventWithParams_Call struct {
	*mock.Call
}

// GetOrganizationEventWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetOrganizationEventApiParams
func (_e *EventsApi_Expecter) GetOrganizationEventWithParams(ctx any, args any) *EventsApi_GetOrganizationEventWithParams_Call {
	return &EventsApi_GetOrganizationEventWithParams_Call{Call: _e.mock.On("GetOrganizationEventWithParams", ctx, args)}
}

func (_c *EventsApi_GetOrganizationEventWithParams_Call) Run(run func(ctx context.Context, args *admin.GetOrganizationEventApiParams)) *EventsApi_GetOrganizationEventWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetOrganizationEventApiParams))
	})
	return _c
}

func (_c *EventsApi_GetOrganizationEventWithParams_Call) Return(_a0 admin.GetOrganizationEventApiRequest) *EventsApi_GetOrganizationEventWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsApi_GetOrganizationEventWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetOrganizationEventApiParams) admin.GetOrganizationEventApiRequest) *EventsApi_GetOrganizationEventWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectEvent provides a mock function with given fields: ctx, groupId, eventId
func (_m *EventsApi) GetProjectEvent(ctx context.Context, groupId string, eventId string) admin.GetProjectEventApiRequest {
	ret := _m.Called(ctx, groupId, eventId)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectEvent")
	}

	var r0 admin.GetProjectEventApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetProjectEventApiRequest); ok {
		r0 = rf(ctx, groupId, eventId)
	} else {
		r0 = ret.Get(0).(admin.GetProjectEventApiRequest)
	}

	return r0
}

// EventsApi_GetProjectEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectEvent'
type EventsApi_GetProjectEvent_Call struct {
	*mock.Call
}

// GetProjectEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - eventId string
func (_e *EventsApi_Expecter) GetProjectEvent(ctx any, groupId any, eventId any) *EventsApi_GetProjectEvent_Call {
	return &EventsApi_GetProjectEvent_Call{Call: _e.mock.On("GetProjectEvent", ctx, groupId, eventId)}
}

func (_c *EventsApi_GetProjectEvent_Call) Run(run func(ctx context.Context, groupId string, eventId string)) *EventsApi_GetProjectEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *EventsApi_GetProjectEvent_Call) Return(_a0 admin.GetProjectEventApiRequest) *EventsApi_GetProjectEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsApi_GetProjectEvent_Call) RunAndReturn(run func(context.Context, string, string) admin.GetProjectEventApiRequest) *EventsApi_GetProjectEvent_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectEventExecute provides a mock function with given fields: r
func (_m *EventsApi) GetProjectEventExecute(r admin.GetProjectEventApiRequest) (*admin.EventViewForNdsGroup, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectEventExecute")
	}

	var r0 *admin.EventViewForNdsGroup
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetProjectEventApiRequest) (*admin.EventViewForNdsGroup, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetProjectEventApiRequest) *admin.EventViewForNdsGroup); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.EventViewForNdsGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetProjectEventApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetProjectEventApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventsApi_GetProjectEventExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectEventExecute'
type EventsApi_GetProjectEventExecute_Call struct {
	*mock.Call
}

// GetProjectEventExecute is a helper method to define mock.On call
//   - r admin.GetProjectEventApiRequest
func (_e *EventsApi_Expecter) GetProjectEventExecute(r any) *EventsApi_GetProjectEventExecute_Call {
	return &EventsApi_GetProjectEventExecute_Call{Call: _e.mock.On("GetProjectEventExecute", r)}
}

func (_c *EventsApi_GetProjectEventExecute_Call) Run(run func(r admin.GetProjectEventApiRequest)) *EventsApi_GetProjectEventExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetProjectEventApiRequest))
	})
	return _c
}

func (_c *EventsApi_GetProjectEventExecute_Call) Return(_a0 *admin.EventViewForNdsGroup, _a1 *http.Response, _a2 error) *EventsApi_GetProjectEventExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EventsApi_GetProjectEventExecute_Call) RunAndReturn(run func(admin.GetProjectEventApiRequest) (*admin.EventViewForNdsGroup, *http.Response, error)) *EventsApi_GetProjectEventExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectEventWithParams provides a mock function with given fields: ctx, args
func (_m *EventsApi) GetProjectEventWithParams(ctx context.Context, args *admin.GetProjectEventApiParams) admin.GetProjectEventApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectEventWithParams")
	}

	var r0 admin.GetProjectEventApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetProjectEventApiParams) admin.GetProjectEventApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetProjectEventApiRequest)
	}

	return r0
}

// EventsApi_GetProjectEventWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectEventWithParams'
type EventsApi_GetProjectEventWithParams_Call struct {
	*mock.Call
}

// GetProjectEventWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetProjectEventApiParams
func (_e *EventsApi_Expecter) GetProjectEventWithParams(ctx any, args any) *EventsApi_GetProjectEventWithParams_Call {
	return &EventsApi_GetProjectEventWithParams_Call{Call: _e.mock.On("GetProjectEventWithParams", ctx, args)}
}

func (_c *EventsApi_GetProjectEventWithParams_Call) Run(run func(ctx context.Context, args *admin.GetProjectEventApiParams)) *EventsApi_GetProjectEventWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetProjectEventApiParams))
	})
	return _c
}

func (_c *EventsApi_GetProjectEventWithParams_Call) Return(_a0 admin.GetProjectEventApiRequest) *EventsApi_GetProjectEventWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsApi_GetProjectEventWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetProjectEventApiParams) admin.GetProjectEventApiRequest) *EventsApi_GetProjectEventWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListEventTypes provides a mock function with given fields: ctx
func (_m *EventsApi) ListEventTypes(ctx context.Context) admin.ListEventTypesApiRequest {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListEventTypes")
	}

	var r0 admin.ListEventTypesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context) admin.ListEventTypesApiRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(admin.ListEventTypesApiRequest)
	}

	return r0
}

// EventsApi_ListEventTypes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListEventTypes'
type EventsApi_ListEventTypes_Call struct {
	*mock.Call
}

// ListEventTypes is a helper method to define mock.On call
//   - ctx context.Context
func (_e *EventsApi_Expecter) ListEventTypes(ctx any) *EventsApi_ListEventTypes_Call {
	return &EventsApi_ListEventTypes_Call{Call: _e.mock.On("ListEventTypes", ctx)}
}

func (_c *EventsApi_ListEventTypes_Call) Run(run func(ctx context.Context)) *EventsApi_ListEventTypes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EventsApi_ListEventTypes_Call) Return(_a0 admin.ListEventTypesApiRequest) *EventsApi_ListEventTypes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsApi_ListEventTypes_Call) RunAndReturn(run func(context.Context) admin.ListEventTypesApiRequest) *EventsApi_ListEventTypes_Call {
	_c.Call.Return(run)
	return _c
}

// ListEventTypesExecute provides a mock function with given fields: r
func (_m *EventsApi) ListEventTypesExecute(r admin.ListEventTypesApiRequest) (*admin.PaginatedEventTypeDetailsResponse, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListEventTypesExecute")
	}

	var r0 *admin.PaginatedEventTypeDetailsResponse
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListEventTypesApiRequest) (*admin.PaginatedEventTypeDetailsResponse, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListEventTypesApiRequest) *admin.PaginatedEventTypeDetailsResponse); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedEventTypeDetailsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListEventTypesApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListEventTypesApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventsApi_ListEventTypesExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListEventTypesExecute'
type EventsApi_ListEventTypesExecute_Call struct {
	*mock.Call
}

// ListEventTypesExecute is a helper method to define mock.On call
//   - r admin.ListEventTypesApiRequest
func (_e *EventsApi_Expecter) ListEventTypesExecute(r any) *EventsApi_ListEventTypesExecute_Call {
	return &EventsApi_ListEventTypesExecute_Call{Call: _e.mock.On("ListEventTypesExecute", r)}
}

func (_c *EventsApi_ListEventTypesExecute_Call) Run(run func(r admin.ListEventTypesApiRequest)) *EventsApi_ListEventTypesExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListEventTypesApiRequest))
	})
	return _c
}

func (_c *EventsApi_ListEventTypesExecute_Call) Return(_a0 *admin.PaginatedEventTypeDetailsResponse, _a1 *http.Response, _a2 error) *EventsApi_ListEventTypesExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EventsApi_ListEventTypesExecute_Call) RunAndReturn(run func(admin.ListEventTypesApiRequest) (*admin.PaginatedEventTypeDetailsResponse, *http.Response, error)) *EventsApi_ListEventTypesExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListEventTypesWithParams provides a mock function with given fields: ctx, args
func (_m *EventsApi) ListEventTypesWithParams(ctx context.Context, args *admin.ListEventTypesApiParams) admin.ListEventTypesApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListEventTypesWithParams")
	}

	var r0 admin.ListEventTypesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListEventTypesApiParams) admin.ListEventTypesApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListEventTypesApiRequest)
	}

	return r0
}

// EventsApi_ListEventTypesWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListEventTypesWithParams'
type EventsApi_ListEventTypesWithParams_Call struct {
	*mock.Call
}

// ListEventTypesWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListEventTypesApiParams
func (_e *EventsApi_Expecter) ListEventTypesWithParams(ctx any, args any) *EventsApi_ListEventTypesWithParams_Call {
	return &EventsApi_ListEventTypesWithParams_Call{Call: _e.mock.On("ListEventTypesWithParams", ctx, args)}
}

func (_c *EventsApi_ListEventTypesWithParams_Call) Run(run func(ctx context.Context, args *admin.ListEventTypesApiParams)) *EventsApi_ListEventTypesWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListEventTypesApiParams))
	})
	return _c
}

func (_c *EventsApi_ListEventTypesWithParams_Call) Return(_a0 admin.ListEventTypesApiRequest) *EventsApi_ListEventTypesWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsApi_ListEventTypesWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListEventTypesApiParams) admin.ListEventTypesApiRequest) *EventsApi_ListEventTypesWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListOrganizationEvents provides a mock function with given fields: ctx, orgId
func (_m *EventsApi) ListOrganizationEvents(ctx context.Context, orgId string) admin.ListOrganizationEventsApiRequest {
	ret := _m.Called(ctx, orgId)

	if len(ret) == 0 {
		panic("no return value specified for ListOrganizationEvents")
	}

	var r0 admin.ListOrganizationEventsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ListOrganizationEventsApiRequest); ok {
		r0 = rf(ctx, orgId)
	} else {
		r0 = ret.Get(0).(admin.ListOrganizationEventsApiRequest)
	}

	return r0
}

// EventsApi_ListOrganizationEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOrganizationEvents'
type EventsApi_ListOrganizationEvents_Call struct {
	*mock.Call
}

// ListOrganizationEvents is a helper method to define mock.On call
//   - ctx context.Context
//   - orgId string
func (_e *EventsApi_Expecter) ListOrganizationEvents(ctx any, orgId any) *EventsApi_ListOrganizationEvents_Call {
	return &EventsApi_ListOrganizationEvents_Call{Call: _e.mock.On("ListOrganizationEvents", ctx, orgId)}
}

func (_c *EventsApi_ListOrganizationEvents_Call) Run(run func(ctx context.Context, orgId string)) *EventsApi_ListOrganizationEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *EventsApi_ListOrganizationEvents_Call) Return(_a0 admin.ListOrganizationEventsApiRequest) *EventsApi_ListOrganizationEvents_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsApi_ListOrganizationEvents_Call) RunAndReturn(run func(context.Context, string) admin.ListOrganizationEventsApiRequest) *EventsApi_ListOrganizationEvents_Call {
	_c.Call.Return(run)
	return _c
}

// ListOrganizationEventsExecute provides a mock function with given fields: r
func (_m *EventsApi) ListOrganizationEventsExecute(r admin.ListOrganizationEventsApiRequest) (*admin.OrgPaginatedEvent, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListOrganizationEventsExecute")
	}

	var r0 *admin.OrgPaginatedEvent
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListOrganizationEventsApiRequest) (*admin.OrgPaginatedEvent, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListOrganizationEventsApiRequest) *admin.OrgPaginatedEvent); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.OrgPaginatedEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListOrganizationEventsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListOrganizationEventsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventsApi_ListOrganizationEventsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOrganizationEventsExecute'
type EventsApi_ListOrganizationEventsExecute_Call struct {
	*mock.Call
}

// ListOrganizationEventsExecute is a helper method to define mock.On call
//   - r admin.ListOrganizationEventsApiRequest
func (_e *EventsApi_Expecter) ListOrganizationEventsExecute(r any) *EventsApi_ListOrganizationEventsExecute_Call {
	return &EventsApi_ListOrganizationEventsExecute_Call{Call: _e.mock.On("ListOrganizationEventsExecute", r)}
}

func (_c *EventsApi_ListOrganizationEventsExecute_Call) Run(run func(r admin.ListOrganizationEventsApiRequest)) *EventsApi_ListOrganizationEventsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListOrganizationEventsApiRequest))
	})
	return _c
}

func (_c *EventsApi_ListOrganizationEventsExecute_Call) Return(_a0 *admin.OrgPaginatedEvent, _a1 *http.Response, _a2 error) *EventsApi_ListOrganizationEventsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EventsApi_ListOrganizationEventsExecute_Call) RunAndReturn(run func(admin.ListOrganizationEventsApiRequest) (*admin.OrgPaginatedEvent, *http.Response, error)) *EventsApi_ListOrganizationEventsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListOrganizationEventsWithParams provides a mock function with given fields: ctx, args
func (_m *EventsApi) ListOrganizationEventsWithParams(ctx context.Context, args *admin.ListOrganizationEventsApiParams) admin.ListOrganizationEventsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListOrganizationEventsWithParams")
	}

	var r0 admin.ListOrganizationEventsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListOrganizationEventsApiParams) admin.ListOrganizationEventsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListOrganizationEventsApiRequest)
	}

	return r0
}

// EventsApi_ListOrganizationEventsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOrganizationEventsWithParams'
type EventsApi_ListOrganizationEventsWithParams_Call struct {
	*mock.Call
}

// ListOrganizationEventsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListOrganizationEventsApiParams
func (_e *EventsApi_Expecter) ListOrganizationEventsWithParams(ctx any, args any) *EventsApi_ListOrganizationEventsWithParams_Call {
	return &EventsApi_ListOrganizationEventsWithParams_Call{Call: _e.mock.On("ListOrganizationEventsWithParams", ctx, args)}
}

func (_c *EventsApi_ListOrganizationEventsWithParams_Call) Run(run func(ctx context.Context, args *admin.ListOrganizationEventsApiParams)) *EventsApi_ListOrganizationEventsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListOrganizationEventsApiParams))
	})
	return _c
}

func (_c *EventsApi_ListOrganizationEventsWithParams_Call) Return(_a0 admin.ListOrganizationEventsApiRequest) *EventsApi_ListOrganizationEventsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsApi_ListOrganizationEventsWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListOrganizationEventsApiParams) admin.ListOrganizationEventsApiRequest) *EventsApi_ListOrganizationEventsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectEvents provides a mock function with given fields: ctx, groupId
func (_m *EventsApi) ListProjectEvents(ctx context.Context, groupId string) admin.ListProjectEventsApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectEvents")
	}

	var r0 admin.ListProjectEventsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ListProjectEventsApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.ListProjectEventsApiRequest)
	}

	return r0
}

// EventsApi_ListProjectEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectEvents'
type EventsApi_ListProjectEvents_Call struct {
	*mock.Call
}

// ListProjectEvents is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *EventsApi_Expecter) ListProjectEvents(ctx any, groupId any) *EventsApi_ListProjectEvents_Call {
	return &EventsApi_ListProjectEvents_Call{Call: _e.mock.On("ListProjectEvents", ctx, groupId)}
}

func (_c *EventsApi_ListProjectEvents_Call) Run(run func(ctx context.Context, groupId string)) *EventsApi_ListProjectEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *EventsApi_ListProjectEvents_Call) Return(_a0 admin.ListProjectEventsApiRequest) *EventsApi_ListProjectEvents_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsApi_ListProjectEvents_Call) RunAndReturn(run func(context.Context, string) admin.ListProjectEventsApiRequest) *EventsApi_ListProjectEvents_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectEventsExecute provides a mock function with given fields: r
func (_m *EventsApi) ListProjectEventsExecute(r admin.ListProjectEventsApiRequest) (*admin.GroupPaginatedEvent, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectEventsExecute")
	}

	var r0 *admin.GroupPaginatedEvent
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListProjectEventsApiRequest) (*admin.GroupPaginatedEvent, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListProjectEventsApiRequest) *admin.GroupPaginatedEvent); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GroupPaginatedEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListProjectEventsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListProjectEventsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventsApi_ListProjectEventsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectEventsExecute'
type EventsApi_ListProjectEventsExecute_Call struct {
	*mock.Call
}

// ListProjectEventsExecute is a helper method to define mock.On call
//   - r admin.ListProjectEventsApiRequest
func (_e *EventsApi_Expecter) ListProjectEventsExecute(r any) *EventsApi_ListProjectEventsExecute_Call {
	return &EventsApi_ListProjectEventsExecute_Call{Call: _e.mock.On("ListProjectEventsExecute", r)}
}

func (_c *EventsApi_ListProjectEventsExecute_Call) Run(run func(r admin.ListProjectEventsApiRequest)) *EventsApi_ListProjectEventsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListProjectEventsApiRequest))
	})
	return _c
}

func (_c *EventsApi_ListProjectEventsExecute_Call) Return(_a0 *admin.GroupPaginatedEvent, _a1 *http.Response, _a2 error) *EventsApi_ListProjectEventsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EventsApi_ListProjectEventsExecute_Call) RunAndReturn(run func(admin.ListProjectEventsApiRequest) (*admin.GroupPaginatedEvent, *http.Response, error)) *EventsApi_ListProjectEventsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectEventsWithParams provides a mock function with given fields: ctx, args
func (_m *EventsApi) ListProjectEventsWithParams(ctx context.Context, args *admin.ListProjectEventsApiParams) admin.ListProjectEventsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectEventsWithParams")
	}

	var r0 admin.ListProjectEventsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListProjectEventsApiParams) admin.ListProjectEventsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListProjectEventsApiRequest)
	}

	return r0
}

// EventsApi_ListProjectEventsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectEventsWithParams'
type EventsApi_ListProjectEventsWithParams_Call struct {
	*mock.Call
}

// ListProjectEventsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListProjectEventsApiParams
func (_e *EventsApi_Expecter) ListProjectEventsWithParams(ctx any, args any) *EventsApi_ListProjectEventsWithParams_Call {
	return &EventsApi_ListProjectEventsWithParams_Call{Call: _e.mock.On("ListProjectEventsWithParams", ctx, args)}
}

func (_c *EventsApi_ListProjectEventsWithParams_Call) Run(run func(ctx context.Context, args *admin.ListProjectEventsApiParams)) *EventsApi_ListProjectEventsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListProjectEventsApiParams))
	})
	return _c
}

func (_c *EventsApi_ListProjectEventsWithParams_Call) Return(_a0 admin.ListProjectEventsApiRequest) *EventsApi_ListProjectEventsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsApi_ListProjectEventsWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListProjectEventsApiParams) admin.ListProjectEventsApiRequest) *EventsApi_ListProjectEventsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewEventsApi creates a new instance of EventsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventsApi {
	mock := &EventsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
