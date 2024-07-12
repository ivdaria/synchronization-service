// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go
//
// Generated by this command:
//
//	mockgen -source=usecase.go -destination=./mocks/mock.go -typed
//

// Package mock_updateclientalgorithms is a generated GoMock package.
package mock_updateclientalgorithms

import (
	context "context"
	reflect "reflect"
	entity "synchronizationService/internal/entity"

	gomock "go.uber.org/mock/gomock"
)

// MockalgorithmStatusesRepo is a mock of algorithmStatusesRepo interface.
type MockalgorithmStatusesRepo struct {
	ctrl     *gomock.Controller
	recorder *MockalgorithmStatusesRepoMockRecorder
}

// MockalgorithmStatusesRepoMockRecorder is the mock recorder for MockalgorithmStatusesRepo.
type MockalgorithmStatusesRepoMockRecorder struct {
	mock *MockalgorithmStatusesRepo
}

// NewMockalgorithmStatusesRepo creates a new mock instance.
func NewMockalgorithmStatusesRepo(ctrl *gomock.Controller) *MockalgorithmStatusesRepo {
	mock := &MockalgorithmStatusesRepo{ctrl: ctrl}
	mock.recorder = &MockalgorithmStatusesRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockalgorithmStatusesRepo) EXPECT() *MockalgorithmStatusesRepoMockRecorder {
	return m.recorder
}

// UpdateAlgorithmStatus mocks base method.
func (m *MockalgorithmStatusesRepo) UpdateAlgorithmStatus(ctx context.Context, algorithm *entity.AlgorithmStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlgorithmStatus", ctx, algorithm)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAlgorithmStatus indicates an expected call of UpdateAlgorithmStatus.
func (mr *MockalgorithmStatusesRepoMockRecorder) UpdateAlgorithmStatus(ctx, algorithm any) *MockalgorithmStatusesRepoUpdateAlgorithmStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlgorithmStatus", reflect.TypeOf((*MockalgorithmStatusesRepo)(nil).UpdateAlgorithmStatus), ctx, algorithm)
	return &MockalgorithmStatusesRepoUpdateAlgorithmStatusCall{Call: call}
}

// MockalgorithmStatusesRepoUpdateAlgorithmStatusCall wrap *gomock.Call
type MockalgorithmStatusesRepoUpdateAlgorithmStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockalgorithmStatusesRepoUpdateAlgorithmStatusCall) Return(arg0 error) *MockalgorithmStatusesRepoUpdateAlgorithmStatusCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockalgorithmStatusesRepoUpdateAlgorithmStatusCall) Do(f func(context.Context, *entity.AlgorithmStatus) error) *MockalgorithmStatusesRepoUpdateAlgorithmStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockalgorithmStatusesRepoUpdateAlgorithmStatusCall) DoAndReturn(f func(context.Context, *entity.AlgorithmStatus) error) *MockalgorithmStatusesRepoUpdateAlgorithmStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}