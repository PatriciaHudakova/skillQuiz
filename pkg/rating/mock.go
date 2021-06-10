// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package rating is a generated GoMock package.
package rating

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRating is a mock of Rating interface.
type MockRating struct {
	ctrl     *gomock.Controller
	recorder *MockRatingMockRecorder
}

// MockRatingMockRecorder is the mock recorder for MockRating.
type MockRatingMockRecorder struct {
	mock *MockRating
}

// NewMockRating creates a new mock instance.
func NewMockRating(ctrl *gomock.Controller) *MockRating {
	mock := &MockRating{ctrl: ctrl}
	mock.recorder = &MockRatingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRating) EXPECT() *MockRatingMockRecorder {
	return m.recorder
}

// CalculateAverageRating mocks base method.
func (m *MockRating) CalculateAverageRating(db *sql.DB, currentRating string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateAverageRating", db, currentRating)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateAverageRating indicates an expected call of CalculateAverageRating.
func (mr *MockRatingMockRecorder) CalculateAverageRating(db, currentRating interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateAverageRating", reflect.TypeOf((*MockRating)(nil).CalculateAverageRating), db, currentRating)
}

// CalculateImmediateRating mocks base method.
func (m *MockRating) CalculateImmediateRating(answers []string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateImmediateRating", answers)
	ret0, _ := ret[0].(string)
	return ret0
}

// CalculateImmediateRating indicates an expected call of CalculateImmediateRating.
func (mr *MockRatingMockRecorder) CalculateImmediateRating(answers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateImmediateRating", reflect.TypeOf((*MockRating)(nil).CalculateImmediateRating), answers)
}

// PrintRatings mocks base method.
func (m *MockRating) PrintRatings(db *sql.DB, answers []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PrintRatings", db, answers)
}

// PrintRatings indicates an expected call of PrintRatings.
func (mr *MockRatingMockRecorder) PrintRatings(db, answers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintRatings", reflect.TypeOf((*MockRating)(nil).PrintRatings), db, answers)
}
