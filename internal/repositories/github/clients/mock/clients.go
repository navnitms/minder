// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stacklok/minder/internal/repositories/github/clients (interfaces: GitHubRepoClient)
//
// Generated by this command:
//
//	mockgen -package mockghclients -destination internal/repositories/github/clients/mock/clients.go github.com/stacklok/minder/internal/repositories/github/clients GitHubRepoClient
//

// Package mockghclients is a generated GoMock package.
package mockghclients

import (
	context "context"
	reflect "reflect"

	github "github.com/google/go-github/v56/github"
	gomock "go.uber.org/mock/gomock"
)

// MockGitHubRepoClient is a mock of GitHubRepoClient interface.
type MockGitHubRepoClient struct {
	ctrl     *gomock.Controller
	recorder *MockGitHubRepoClientMockRecorder
}

// MockGitHubRepoClientMockRecorder is the mock recorder for MockGitHubRepoClient.
type MockGitHubRepoClientMockRecorder struct {
	mock *MockGitHubRepoClient
}

// NewMockGitHubRepoClient creates a new mock instance.
func NewMockGitHubRepoClient(ctrl *gomock.Controller) *MockGitHubRepoClient {
	mock := &MockGitHubRepoClient{ctrl: ctrl}
	mock.recorder = &MockGitHubRepoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitHubRepoClient) EXPECT() *MockGitHubRepoClientMockRecorder {
	return m.recorder
}

// CreateHook mocks base method.
func (m *MockGitHubRepoClient) CreateHook(arg0 context.Context, arg1, arg2 string, arg3 *github.Hook) (*github.Hook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHook", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Hook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHook indicates an expected call of CreateHook.
func (mr *MockGitHubRepoClientMockRecorder) CreateHook(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHook", reflect.TypeOf((*MockGitHubRepoClient)(nil).CreateHook), arg0, arg1, arg2, arg3)
}

// DeleteHook mocks base method.
func (m *MockGitHubRepoClient) DeleteHook(arg0 context.Context, arg1, arg2 string, arg3 int64) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHook", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteHook indicates an expected call of DeleteHook.
func (mr *MockGitHubRepoClientMockRecorder) DeleteHook(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHook", reflect.TypeOf((*MockGitHubRepoClient)(nil).DeleteHook), arg0, arg1, arg2, arg3)
}

// GetRepository mocks base method.
func (m *MockGitHubRepoClient) GetRepository(arg0 context.Context, arg1, arg2 string) (*github.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepository", arg0, arg1, arg2)
	ret0, _ := ret[0].(*github.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepository indicates an expected call of GetRepository.
func (mr *MockGitHubRepoClientMockRecorder) GetRepository(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepository", reflect.TypeOf((*MockGitHubRepoClient)(nil).GetRepository), arg0, arg1, arg2)
}

// ListHooks mocks base method.
func (m *MockGitHubRepoClient) ListHooks(arg0 context.Context, arg1, arg2 string) ([]*github.Hook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHooks", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.Hook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHooks indicates an expected call of ListHooks.
func (mr *MockGitHubRepoClientMockRecorder) ListHooks(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHooks", reflect.TypeOf((*MockGitHubRepoClient)(nil).ListHooks), arg0, arg1, arg2)
}
