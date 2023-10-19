// Code generated by MockGen. DO NOT EDIT.
// Source: service/post_service.go
//
// Generated by this command:
//
//	mockgen -package=mocks -destination=_mocks/post_service.go -source=service/post_service.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	gomock "go.uber.org/mock/gomock"
)

// MockPostService is a mock of PostService interface.
type MockPostService struct {
	ctrl     *gomock.Controller
	recorder *MockPostServiceMockRecorder
}

// MockPostServiceMockRecorder is the mock recorder for MockPostService.
type MockPostServiceMockRecorder struct {
	mock *MockPostService
}

// NewMockPostService creates a new mock instance.
func NewMockPostService(ctrl *gomock.Controller) *MockPostService {
	mock := &MockPostService{ctrl: ctrl}
	mock.recorder = &MockPostServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostService) EXPECT() *MockPostServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPostService) Create(post *models.Post) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", post)
}

// Create indicates an expected call of Create.
func (mr *MockPostServiceMockRecorder) Create(post any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostService)(nil).Create), post)
}

// Delete mocks base method.
func (m *MockPostService) Delete(post *models.Post) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", post)
}

// Delete indicates an expected call of Delete.
func (mr *MockPostServiceMockRecorder) Delete(post any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostService)(nil).Delete), post)
}

// GetAll mocks base method.
func (m *MockPostService) GetAll() []models.Post {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Post)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPostServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPostService)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockPostService) GetByID(id uint) models.Post {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(models.Post)
	return ret0
}

// GetByID indicates an expected call of GetByID.
func (mr *MockPostServiceMockRecorder) GetByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockPostService)(nil).GetByID), id)
}

// Update mocks base method.
func (m *MockPostService) Update(post *models.Post) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", post)
}

// Update indicates an expected call of Update.
func (mr *MockPostServiceMockRecorder) Update(post any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostService)(nil).Update), post)
}
