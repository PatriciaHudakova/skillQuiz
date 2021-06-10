// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package db is a generated GoMock package.
package db

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIDatabase is a mock of IDatabase interface.
type MockIDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockIDatabaseMockRecorder
}

// MockIDatabaseMockRecorder is the mock recorder for MockIDatabase.
type MockIDatabaseMockRecorder struct {
	mock *MockIDatabase
}

// NewMockIDatabase creates a new mock instance.
func NewMockIDatabase(ctrl *gomock.Controller) *MockIDatabase {
	mock := &MockIDatabase{ctrl: ctrl}
	mock.recorder = &MockIDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDatabase) EXPECT() *MockIDatabaseMockRecorder {
	return m.recorder
}

// GetAllRows mocks base method.
func (m *MockIDatabase) GetAllRows() (*sql.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRows")
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRows indicates an expected call of GetAllRows.
func (mr *MockIDatabaseMockRecorder) GetAllRows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRows", reflect.TypeOf((*MockIDatabase)(nil).GetAllRows))
}

// GetOverallAverageFromDB mocks base method.
func (m *MockIDatabase) GetOverallAverageFromDB() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOverallAverageFromDB")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOverallAverageFromDB indicates an expected call of GetOverallAverageFromDB.
func (mr *MockIDatabaseMockRecorder) GetOverallAverageFromDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOverallAverageFromDB", reflect.TypeOf((*MockIDatabase)(nil).GetOverallAverageFromDB))
}

// MakeCurrentRatingTheAverage mocks base method.
func (m *MockIDatabase) MakeCurrentRatingTheAverage(currentRating string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeCurrentRatingTheAverage", currentRating)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeCurrentRatingTheAverage indicates an expected call of MakeCurrentRatingTheAverage.
func (mr *MockIDatabaseMockRecorder) MakeCurrentRatingTheAverage(currentRating interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeCurrentRatingTheAverage", reflect.TypeOf((*MockIDatabase)(nil).MakeCurrentRatingTheAverage), currentRating)
}

// UpdateAverage mocks base method.
func (m *MockIDatabase) UpdateAverage(newAverage int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAverage", newAverage)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAverage indicates an expected call of UpdateAverage.
func (mr *MockIDatabaseMockRecorder) UpdateAverage(newAverage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAverage", reflect.TypeOf((*MockIDatabase)(nil).UpdateAverage), newAverage)
}