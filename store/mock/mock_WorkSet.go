// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package store

import (
	"trainer-helper/model"

	mock "github.com/stretchr/testify/mock"
)

// NewMockWorkSet creates a new instance of MockWorkSet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWorkSet(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWorkSet {
	mock := &MockWorkSet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockWorkSet is an autogenerated mock type for the WorkSet type
type MockWorkSet struct {
	mock.Mock
}

type MockWorkSet_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWorkSet) EXPECT() *MockWorkSet_Expecter {
	return &MockWorkSet_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function for the type MockWorkSet
func (_mock *MockWorkSet) Delete(modelId int) error {
	ret := _mock.Called(modelId)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(int) error); ok {
		r0 = returnFunc(modelId)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockWorkSet_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockWorkSet_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - modelId
func (_e *MockWorkSet_Expecter) Delete(modelId interface{}) *MockWorkSet_Delete_Call {
	return &MockWorkSet_Delete_Call{Call: _e.mock.On("Delete", modelId)}
}

func (_c *MockWorkSet_Delete_Call) Run(run func(modelId int)) *MockWorkSet_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockWorkSet_Delete_Call) Return(err error) *MockWorkSet_Delete_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockWorkSet_Delete_Call) RunAndReturn(run func(modelId int) error) *MockWorkSet_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteMany provides a mock function for the type MockWorkSet
func (_mock *MockWorkSet) DeleteMany(modelIds []int) error {
	ret := _mock.Called(modelIds)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMany")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func([]int) error); ok {
		r0 = returnFunc(modelIds)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockWorkSet_DeleteMany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteMany'
type MockWorkSet_DeleteMany_Call struct {
	*mock.Call
}

// DeleteMany is a helper method to define mock.On call
//   - modelIds
func (_e *MockWorkSet_Expecter) DeleteMany(modelIds interface{}) *MockWorkSet_DeleteMany_Call {
	return &MockWorkSet_DeleteMany_Call{Call: _e.mock.On("DeleteMany", modelIds)}
}

func (_c *MockWorkSet_DeleteMany_Call) Run(run func(modelIds []int)) *MockWorkSet_DeleteMany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]int))
	})
	return _c
}

func (_c *MockWorkSet_DeleteMany_Call) Return(err error) *MockWorkSet_DeleteMany_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockWorkSet_DeleteMany_Call) RunAndReturn(run func(modelIds []int) error) *MockWorkSet_DeleteMany_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function for the type MockWorkSet
func (_mock *MockWorkSet) Get() ([]model.WorkSet, error) {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []model.WorkSet
	var r1 error
	if returnFunc, ok := ret.Get(0).(func() ([]model.WorkSet, error)); ok {
		return returnFunc()
	}
	if returnFunc, ok := ret.Get(0).(func() []model.WorkSet); ok {
		r0 = returnFunc()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.WorkSet)
		}
	}
	if returnFunc, ok := ret.Get(1).(func() error); ok {
		r1 = returnFunc()
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockWorkSet_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockWorkSet_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *MockWorkSet_Expecter) Get() *MockWorkSet_Get_Call {
	return &MockWorkSet_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *MockWorkSet_Get_Call) Run(run func()) *MockWorkSet_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWorkSet_Get_Call) Return(workSets []model.WorkSet, err error) *MockWorkSet_Get_Call {
	_c.Call.Return(workSets, err)
	return _c
}

func (_c *MockWorkSet_Get_Call) RunAndReturn(run func() ([]model.WorkSet, error)) *MockWorkSet_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function for the type MockWorkSet
func (_mock *MockWorkSet) GetById(modelId int) (model.WorkSet, error) {
	ret := _mock.Called(modelId)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 model.WorkSet
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(int) (model.WorkSet, error)); ok {
		return returnFunc(modelId)
	}
	if returnFunc, ok := ret.Get(0).(func(int) model.WorkSet); ok {
		r0 = returnFunc(modelId)
	} else {
		r0 = ret.Get(0).(model.WorkSet)
	}
	if returnFunc, ok := ret.Get(1).(func(int) error); ok {
		r1 = returnFunc(modelId)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockWorkSet_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockWorkSet_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - modelId
func (_e *MockWorkSet_Expecter) GetById(modelId interface{}) *MockWorkSet_GetById_Call {
	return &MockWorkSet_GetById_Call{Call: _e.mock.On("GetById", modelId)}
}

func (_c *MockWorkSet_GetById_Call) Run(run func(modelId int)) *MockWorkSet_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockWorkSet_GetById_Call) Return(model1 model.WorkSet, err error) *MockWorkSet_GetById_Call {
	_c.Call.Return(model1, err)
	return _c
}

func (_c *MockWorkSet_GetById_Call) RunAndReturn(run func(modelId int) (model.WorkSet, error)) *MockWorkSet_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// Insert provides a mock function for the type MockWorkSet
func (_mock *MockWorkSet) Insert(model1 *model.WorkSet) error {
	ret := _mock.Called(model1)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(*model.WorkSet) error); ok {
		r0 = returnFunc(model1)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockWorkSet_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type MockWorkSet_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - model1
func (_e *MockWorkSet_Expecter) Insert(model1 interface{}) *MockWorkSet_Insert_Call {
	return &MockWorkSet_Insert_Call{Call: _e.mock.On("Insert", model1)}
}

func (_c *MockWorkSet_Insert_Call) Run(run func(model1 *model.WorkSet)) *MockWorkSet_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.WorkSet))
	})
	return _c
}

func (_c *MockWorkSet_Insert_Call) Return(err error) *MockWorkSet_Insert_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockWorkSet_Insert_Call) RunAndReturn(run func(model1 *model.WorkSet) error) *MockWorkSet_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// InsertMany provides a mock function for the type MockWorkSet
func (_mock *MockWorkSet) InsertMany(models *[]model.WorkSet) error {
	ret := _mock.Called(models)

	if len(ret) == 0 {
		panic("no return value specified for InsertMany")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(*[]model.WorkSet) error); ok {
		r0 = returnFunc(models)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockWorkSet_InsertMany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertMany'
type MockWorkSet_InsertMany_Call struct {
	*mock.Call
}

// InsertMany is a helper method to define mock.On call
//   - models
func (_e *MockWorkSet_Expecter) InsertMany(models interface{}) *MockWorkSet_InsertMany_Call {
	return &MockWorkSet_InsertMany_Call{Call: _e.mock.On("InsertMany", models)}
}

func (_c *MockWorkSet_InsertMany_Call) Run(run func(models *[]model.WorkSet)) *MockWorkSet_InsertMany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*[]model.WorkSet))
	})
	return _c
}

func (_c *MockWorkSet_InsertMany_Call) Return(err error) *MockWorkSet_InsertMany_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockWorkSet_InsertMany_Call) RunAndReturn(run func(models *[]model.WorkSet) error) *MockWorkSet_InsertMany_Call {
	_c.Call.Return(run)
	return _c
}

// UndeleteMany provides a mock function for the type MockWorkSet
func (_mock *MockWorkSet) UndeleteMany(modelIds []int) error {
	ret := _mock.Called(modelIds)

	if len(ret) == 0 {
		panic("no return value specified for UndeleteMany")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func([]int) error); ok {
		r0 = returnFunc(modelIds)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockWorkSet_UndeleteMany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UndeleteMany'
type MockWorkSet_UndeleteMany_Call struct {
	*mock.Call
}

// UndeleteMany is a helper method to define mock.On call
//   - modelIds
func (_e *MockWorkSet_Expecter) UndeleteMany(modelIds interface{}) *MockWorkSet_UndeleteMany_Call {
	return &MockWorkSet_UndeleteMany_Call{Call: _e.mock.On("UndeleteMany", modelIds)}
}

func (_c *MockWorkSet_UndeleteMany_Call) Run(run func(modelIds []int)) *MockWorkSet_UndeleteMany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]int))
	})
	return _c
}

func (_c *MockWorkSet_UndeleteMany_Call) Return(err error) *MockWorkSet_UndeleteMany_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockWorkSet_UndeleteMany_Call) RunAndReturn(run func(modelIds []int) error) *MockWorkSet_UndeleteMany_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function for the type MockWorkSet
func (_mock *MockWorkSet) Update(model1 *model.WorkSet) error {
	ret := _mock.Called(model1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(*model.WorkSet) error); ok {
		r0 = returnFunc(model1)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockWorkSet_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockWorkSet_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - model1
func (_e *MockWorkSet_Expecter) Update(model1 interface{}) *MockWorkSet_Update_Call {
	return &MockWorkSet_Update_Call{Call: _e.mock.On("Update", model1)}
}

func (_c *MockWorkSet_Update_Call) Run(run func(model1 *model.WorkSet)) *MockWorkSet_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.WorkSet))
	})
	return _c
}

func (_c *MockWorkSet_Update_Call) Return(err error) *MockWorkSet_Update_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockWorkSet_Update_Call) RunAndReturn(run func(model1 *model.WorkSet) error) *MockWorkSet_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateMany provides a mock function for the type MockWorkSet
func (_mock *MockWorkSet) UpdateMany(models []model.WorkSet) error {
	ret := _mock.Called(models)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMany")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func([]model.WorkSet) error); ok {
		r0 = returnFunc(models)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockWorkSet_UpdateMany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMany'
type MockWorkSet_UpdateMany_Call struct {
	*mock.Call
}

// UpdateMany is a helper method to define mock.On call
//   - models
func (_e *MockWorkSet_Expecter) UpdateMany(models interface{}) *MockWorkSet_UpdateMany_Call {
	return &MockWorkSet_UpdateMany_Call{Call: _e.mock.On("UpdateMany", models)}
}

func (_c *MockWorkSet_UpdateMany_Call) Run(run func(models []model.WorkSet)) *MockWorkSet_UpdateMany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]model.WorkSet))
	})
	return _c
}

func (_c *MockWorkSet_UpdateMany_Call) Return(err error) *MockWorkSet_UpdateMany_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockWorkSet_UpdateMany_Call) RunAndReturn(run func(models []model.WorkSet) error) *MockWorkSet_UpdateMany_Call {
	_c.Call.Return(run)
	return _c
}
