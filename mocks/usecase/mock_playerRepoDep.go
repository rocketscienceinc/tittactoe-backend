// Code generated by mockery v2.45.0. DO NOT EDIT.

package usecase

import (
	context "context"

	entity "github.com/rocketscienceinc/tictactoe-backend/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockplayerRepoDep is an autogenerated mock type for the playerRepoDep type
type MockplayerRepoDep struct {
	mock.Mock
}

type MockplayerRepoDep_Expecter struct {
	mock *mock.Mock
}

func (_m *MockplayerRepoDep) EXPECT() *MockplayerRepoDep_Expecter {
	return &MockplayerRepoDep_Expecter{mock: &_m.Mock}
}

// CreateOrUpdate provides a mock function with given fields: ctx, player
func (_m *MockplayerRepoDep) CreateOrUpdate(ctx context.Context, player *entity.Player) error {
	ret := _m.Called(ctx, player)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrUpdate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Player) error); ok {
		r0 = rf(ctx, player)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockplayerRepoDep_CreateOrUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrUpdate'
type MockplayerRepoDep_CreateOrUpdate_Call struct {
	*mock.Call
}

// CreateOrUpdate is a helper method to define mock.On call
//   - ctx context.Context
//   - player *entity.Player
func (_e *MockplayerRepoDep_Expecter) CreateOrUpdate(ctx interface{}, player interface{}) *MockplayerRepoDep_CreateOrUpdate_Call {
	return &MockplayerRepoDep_CreateOrUpdate_Call{Call: _e.mock.On("CreateOrUpdate", ctx, player)}
}

func (_c *MockplayerRepoDep_CreateOrUpdate_Call) Run(run func(ctx context.Context, player *entity.Player)) *MockplayerRepoDep_CreateOrUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Player))
	})
	return _c
}

func (_c *MockplayerRepoDep_CreateOrUpdate_Call) Return(_a0 error) *MockplayerRepoDep_CreateOrUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockplayerRepoDep_CreateOrUpdate_Call) RunAndReturn(run func(context.Context, *entity.Player) error) *MockplayerRepoDep_CreateOrUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *MockplayerRepoDep) GetByID(ctx context.Context, id string) (*entity.Player, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *entity.Player
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Player, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Player); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Player)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockplayerRepoDep_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type MockplayerRepoDep_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockplayerRepoDep_Expecter) GetByID(ctx interface{}, id interface{}) *MockplayerRepoDep_GetByID_Call {
	return &MockplayerRepoDep_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *MockplayerRepoDep_GetByID_Call) Run(run func(ctx context.Context, id string)) *MockplayerRepoDep_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockplayerRepoDep_GetByID_Call) Return(_a0 *entity.Player, _a1 error) *MockplayerRepoDep_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockplayerRepoDep_GetByID_Call) RunAndReturn(run func(context.Context, string) (*entity.Player, error)) *MockplayerRepoDep_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockplayerRepoDep creates a new instance of MockplayerRepoDep. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockplayerRepoDep(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockplayerRepoDep {
	mock := &MockplayerRepoDep{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
