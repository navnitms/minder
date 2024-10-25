// Code generated by MockGen. DO NOT EDIT.
// Source: ./selectors.go
//
// Generated by this command:
//
//	mockgen -package mock_selectors -destination=./mock/selectors.go -source=./selectors.go
//

// Package mock_selectors is a generated GoMock package.
package mock_selectors

import (
	reflect "reflect"

	proto "github.com/mindersec/minder/internal/proto"
	v1 "github.com/mindersec/minder/pkg/api/protobuf/go/minder/v1"
	selectors "github.com/mindersec/minder/pkg/engine/selectors"
	models "github.com/mindersec/minder/pkg/profiles/models"
	gomock "go.uber.org/mock/gomock"
)

// MockSelectionBuilder is a mock of SelectionBuilder interface.
type MockSelectionBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockSelectionBuilderMockRecorder
	isgomock struct{}
}

// MockSelectionBuilderMockRecorder is the mock recorder for MockSelectionBuilder.
type MockSelectionBuilderMockRecorder struct {
	mock *MockSelectionBuilder
}

// NewMockSelectionBuilder creates a new mock instance.
func NewMockSelectionBuilder(ctrl *gomock.Controller) *MockSelectionBuilder {
	mock := &MockSelectionBuilder{ctrl: ctrl}
	mock.recorder = &MockSelectionBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSelectionBuilder) EXPECT() *MockSelectionBuilderMockRecorder {
	return m.recorder
}

// NewSelectionFromProfile mocks base method.
func (m *MockSelectionBuilder) NewSelectionFromProfile(arg0 v1.Entity, arg1 []models.ProfileSelector) (selectors.Selection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSelectionFromProfile", arg0, arg1)
	ret0, _ := ret[0].(selectors.Selection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSelectionFromProfile indicates an expected call of NewSelectionFromProfile.
func (mr *MockSelectionBuilderMockRecorder) NewSelectionFromProfile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSelectionFromProfile", reflect.TypeOf((*MockSelectionBuilder)(nil).NewSelectionFromProfile), arg0, arg1)
}

// MockSelectionChecker is a mock of SelectionChecker interface.
type MockSelectionChecker struct {
	ctrl     *gomock.Controller
	recorder *MockSelectionCheckerMockRecorder
	isgomock struct{}
}

// MockSelectionCheckerMockRecorder is the mock recorder for MockSelectionChecker.
type MockSelectionCheckerMockRecorder struct {
	mock *MockSelectionChecker
}

// NewMockSelectionChecker creates a new mock instance.
func NewMockSelectionChecker(ctrl *gomock.Controller) *MockSelectionChecker {
	mock := &MockSelectionChecker{ctrl: ctrl}
	mock.recorder = &MockSelectionCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSelectionChecker) EXPECT() *MockSelectionCheckerMockRecorder {
	return m.recorder
}

// CheckSelector mocks base method.
func (m *MockSelectionChecker) CheckSelector(arg0 *v1.Profile_Selector) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSelector", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckSelector indicates an expected call of CheckSelector.
func (mr *MockSelectionCheckerMockRecorder) CheckSelector(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSelector", reflect.TypeOf((*MockSelectionChecker)(nil).CheckSelector), arg0)
}

// MockSelection is a mock of Selection interface.
type MockSelection struct {
	ctrl     *gomock.Controller
	recorder *MockSelectionMockRecorder
	isgomock struct{}
}

// MockSelectionMockRecorder is the mock recorder for MockSelection.
type MockSelectionMockRecorder struct {
	mock *MockSelection
}

// NewMockSelection creates a new mock instance.
func NewMockSelection(ctrl *gomock.Controller) *MockSelection {
	mock := &MockSelection{ctrl: ctrl}
	mock.recorder = &MockSelectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSelection) EXPECT() *MockSelectionMockRecorder {
	return m.recorder
}

// Select mocks base method.
func (m *MockSelection) Select(arg0 *proto.SelectorEntity, arg1 ...selectors.SelectOption) (bool, string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Select indicates an expected call of Select.
func (mr *MockSelectionMockRecorder) Select(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockSelection)(nil).Select), varargs...)
}
