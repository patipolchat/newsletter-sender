// Code generated by mockery v2.42.3. DO NOT EDIT.

package repository

import (
	context "context"
	sqlc "energy-response-assignment/db/sqlc"

	mock "github.com/stretchr/testify/mock"
)

// Newsletter is an autogenerated mock type for the Newsletter type
type Newsletter struct {
	mock.Mock
}

type Newsletter_Expecter struct {
	mock *mock.Mock
}

func (_m *Newsletter) EXPECT() *Newsletter_Expecter {
	return &Newsletter_Expecter{mock: &_m.Mock}
}

// BatchUpdateSendingEmailsStatus provides a mock function with given fields: ctx, arg
func (_m *Newsletter) BatchUpdateSendingEmailsStatus(ctx context.Context, arg []sqlc.BatchUpdateSendingEmailsStatusParams) *sqlc.BatchUpdateSendingEmailsStatusBatchResults {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for BatchUpdateSendingEmailsStatus")
	}

	var r0 *sqlc.BatchUpdateSendingEmailsStatusBatchResults
	if rf, ok := ret.Get(0).(func(context.Context, []sqlc.BatchUpdateSendingEmailsStatusParams) *sqlc.BatchUpdateSendingEmailsStatusBatchResults); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlc.BatchUpdateSendingEmailsStatusBatchResults)
		}
	}

	return r0
}

// Newsletter_BatchUpdateSendingEmailsStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchUpdateSendingEmailsStatus'
type Newsletter_BatchUpdateSendingEmailsStatus_Call struct {
	*mock.Call
}

// BatchUpdateSendingEmailsStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - arg []sqlc.BatchUpdateSendingEmailsStatusParams
func (_e *Newsletter_Expecter) BatchUpdateSendingEmailsStatus(ctx interface{}, arg interface{}) *Newsletter_BatchUpdateSendingEmailsStatus_Call {
	return &Newsletter_BatchUpdateSendingEmailsStatus_Call{Call: _e.mock.On("BatchUpdateSendingEmailsStatus", ctx, arg)}
}

func (_c *Newsletter_BatchUpdateSendingEmailsStatus_Call) Run(run func(ctx context.Context, arg []sqlc.BatchUpdateSendingEmailsStatusParams)) *Newsletter_BatchUpdateSendingEmailsStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]sqlc.BatchUpdateSendingEmailsStatusParams))
	})
	return _c
}

func (_c *Newsletter_BatchUpdateSendingEmailsStatus_Call) Return(_a0 *sqlc.BatchUpdateSendingEmailsStatusBatchResults) *Newsletter_BatchUpdateSendingEmailsStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Newsletter_BatchUpdateSendingEmailsStatus_Call) RunAndReturn(run func(context.Context, []sqlc.BatchUpdateSendingEmailsStatusParams) *sqlc.BatchUpdateSendingEmailsStatusBatchResults) *Newsletter_BatchUpdateSendingEmailsStatus_Call {
	_c.Call.Return(run)
	return _c
}

// CountActiveSubscriber provides a mock function with given fields: ctx
func (_m *Newsletter) CountActiveSubscriber(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CountActiveSubscriber")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Newsletter_CountActiveSubscriber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountActiveSubscriber'
type Newsletter_CountActiveSubscriber_Call struct {
	*mock.Call
}

// CountActiveSubscriber is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Newsletter_Expecter) CountActiveSubscriber(ctx interface{}) *Newsletter_CountActiveSubscriber_Call {
	return &Newsletter_CountActiveSubscriber_Call{Call: _e.mock.On("CountActiveSubscriber", ctx)}
}

func (_c *Newsletter_CountActiveSubscriber_Call) Run(run func(ctx context.Context)) *Newsletter_CountActiveSubscriber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Newsletter_CountActiveSubscriber_Call) Return(_a0 int64, _a1 error) *Newsletter_CountActiveSubscriber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Newsletter_CountActiveSubscriber_Call) RunAndReturn(run func(context.Context) (int64, error)) *Newsletter_CountActiveSubscriber_Call {
	_c.Call.Return(run)
	return _c
}

// CreateNewsLetter provides a mock function with given fields: ctx, arg
func (_m *Newsletter) CreateNewsLetter(ctx context.Context, arg sqlc.CreateNewsLetterParams) (sqlc.Newsletter, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateNewsLetter")
	}

	var r0 sqlc.Newsletter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sqlc.CreateNewsLetterParams) (sqlc.Newsletter, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sqlc.CreateNewsLetterParams) sqlc.Newsletter); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(sqlc.Newsletter)
	}

	if rf, ok := ret.Get(1).(func(context.Context, sqlc.CreateNewsLetterParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Newsletter_CreateNewsLetter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateNewsLetter'
type Newsletter_CreateNewsLetter_Call struct {
	*mock.Call
}

// CreateNewsLetter is a helper method to define mock.On call
//   - ctx context.Context
//   - arg sqlc.CreateNewsLetterParams
func (_e *Newsletter_Expecter) CreateNewsLetter(ctx interface{}, arg interface{}) *Newsletter_CreateNewsLetter_Call {
	return &Newsletter_CreateNewsLetter_Call{Call: _e.mock.On("CreateNewsLetter", ctx, arg)}
}

func (_c *Newsletter_CreateNewsLetter_Call) Run(run func(ctx context.Context, arg sqlc.CreateNewsLetterParams)) *Newsletter_CreateNewsLetter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sqlc.CreateNewsLetterParams))
	})
	return _c
}

func (_c *Newsletter_CreateNewsLetter_Call) Return(_a0 sqlc.Newsletter, _a1 error) *Newsletter_CreateNewsLetter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Newsletter_CreateNewsLetter_Call) RunAndReturn(run func(context.Context, sqlc.CreateNewsLetterParams) (sqlc.Newsletter, error)) *Newsletter_CreateNewsLetter_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSendingEmails provides a mock function with given fields: ctx, arg
func (_m *Newsletter) CreateSendingEmails(ctx context.Context, arg []sqlc.CreateSendingEmailsParams) *sqlc.CreateSendingEmailsBatchResults {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateSendingEmails")
	}

	var r0 *sqlc.CreateSendingEmailsBatchResults
	if rf, ok := ret.Get(0).(func(context.Context, []sqlc.CreateSendingEmailsParams) *sqlc.CreateSendingEmailsBatchResults); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlc.CreateSendingEmailsBatchResults)
		}
	}

	return r0
}

// Newsletter_CreateSendingEmails_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSendingEmails'
type Newsletter_CreateSendingEmails_Call struct {
	*mock.Call
}

// CreateSendingEmails is a helper method to define mock.On call
//   - ctx context.Context
//   - arg []sqlc.CreateSendingEmailsParams
func (_e *Newsletter_Expecter) CreateSendingEmails(ctx interface{}, arg interface{}) *Newsletter_CreateSendingEmails_Call {
	return &Newsletter_CreateSendingEmails_Call{Call: _e.mock.On("CreateSendingEmails", ctx, arg)}
}

func (_c *Newsletter_CreateSendingEmails_Call) Run(run func(ctx context.Context, arg []sqlc.CreateSendingEmailsParams)) *Newsletter_CreateSendingEmails_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]sqlc.CreateSendingEmailsParams))
	})
	return _c
}

func (_c *Newsletter_CreateSendingEmails_Call) Return(_a0 *sqlc.CreateSendingEmailsBatchResults) *Newsletter_CreateSendingEmails_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Newsletter_CreateSendingEmails_Call) RunAndReturn(run func(context.Context, []sqlc.CreateSendingEmailsParams) *sqlc.CreateSendingEmailsBatchResults) *Newsletter_CreateSendingEmails_Call {
	_c.Call.Return(run)
	return _c
}

// GetActiveSubscribersWithPaginate provides a mock function with given fields: ctx, arg
func (_m *Newsletter) GetActiveSubscribersWithPaginate(ctx context.Context, arg sqlc.GetActiveSubscribersWithPaginateParams) ([]sqlc.Subscriber, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetActiveSubscribersWithPaginate")
	}

	var r0 []sqlc.Subscriber
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sqlc.GetActiveSubscribersWithPaginateParams) ([]sqlc.Subscriber, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sqlc.GetActiveSubscribersWithPaginateParams) []sqlc.Subscriber); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sqlc.Subscriber)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sqlc.GetActiveSubscribersWithPaginateParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Newsletter_GetActiveSubscribersWithPaginate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActiveSubscribersWithPaginate'
type Newsletter_GetActiveSubscribersWithPaginate_Call struct {
	*mock.Call
}

// GetActiveSubscribersWithPaginate is a helper method to define mock.On call
//   - ctx context.Context
//   - arg sqlc.GetActiveSubscribersWithPaginateParams
func (_e *Newsletter_Expecter) GetActiveSubscribersWithPaginate(ctx interface{}, arg interface{}) *Newsletter_GetActiveSubscribersWithPaginate_Call {
	return &Newsletter_GetActiveSubscribersWithPaginate_Call{Call: _e.mock.On("GetActiveSubscribersWithPaginate", ctx, arg)}
}

func (_c *Newsletter_GetActiveSubscribersWithPaginate_Call) Run(run func(ctx context.Context, arg sqlc.GetActiveSubscribersWithPaginateParams)) *Newsletter_GetActiveSubscribersWithPaginate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sqlc.GetActiveSubscribersWithPaginateParams))
	})
	return _c
}

func (_c *Newsletter_GetActiveSubscribersWithPaginate_Call) Return(_a0 []sqlc.Subscriber, _a1 error) *Newsletter_GetActiveSubscribersWithPaginate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Newsletter_GetActiveSubscribersWithPaginate_Call) RunAndReturn(run func(context.Context, sqlc.GetActiveSubscribersWithPaginateParams) ([]sqlc.Subscriber, error)) *Newsletter_GetActiveSubscribersWithPaginate_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSendingEmailsStatus provides a mock function with given fields: ctx, arg
func (_m *Newsletter) UpdateSendingEmailsStatus(ctx context.Context, arg sqlc.UpdateSendingEmailsStatusParams) (sqlc.SendingEmail, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSendingEmailsStatus")
	}

	var r0 sqlc.SendingEmail
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sqlc.UpdateSendingEmailsStatusParams) (sqlc.SendingEmail, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sqlc.UpdateSendingEmailsStatusParams) sqlc.SendingEmail); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(sqlc.SendingEmail)
	}

	if rf, ok := ret.Get(1).(func(context.Context, sqlc.UpdateSendingEmailsStatusParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Newsletter_UpdateSendingEmailsStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSendingEmailsStatus'
type Newsletter_UpdateSendingEmailsStatus_Call struct {
	*mock.Call
}

// UpdateSendingEmailsStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - arg sqlc.UpdateSendingEmailsStatusParams
func (_e *Newsletter_Expecter) UpdateSendingEmailsStatus(ctx interface{}, arg interface{}) *Newsletter_UpdateSendingEmailsStatus_Call {
	return &Newsletter_UpdateSendingEmailsStatus_Call{Call: _e.mock.On("UpdateSendingEmailsStatus", ctx, arg)}
}

func (_c *Newsletter_UpdateSendingEmailsStatus_Call) Run(run func(ctx context.Context, arg sqlc.UpdateSendingEmailsStatusParams)) *Newsletter_UpdateSendingEmailsStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sqlc.UpdateSendingEmailsStatusParams))
	})
	return _c
}

func (_c *Newsletter_UpdateSendingEmailsStatus_Call) Return(_a0 sqlc.SendingEmail, _a1 error) *Newsletter_UpdateSendingEmailsStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Newsletter_UpdateSendingEmailsStatus_Call) RunAndReturn(run func(context.Context, sqlc.UpdateSendingEmailsStatusParams) (sqlc.SendingEmail, error)) *Newsletter_UpdateSendingEmailsStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewNewsletter creates a new instance of Newsletter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNewsletter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Newsletter {
	mock := &Newsletter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
