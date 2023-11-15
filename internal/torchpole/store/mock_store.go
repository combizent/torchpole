// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rppkg/torchpole.

package store

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/rppkg/torchpole/internal/pkg/model"
)

// MockIStore 是IStore的mock实现.
type MockIStore struct {
	ctrl     *gomock.Controller
	recorder *MockIStoreMockRecorder
}

type MockIStoreMockRecorder struct {
	mock *MockIStore
}

func NewMockIStore(ctrl *gomock.Controller) *MockIStore {
	mock := &MockIStore{ctrl: ctrl}
	mock.recorder = &MockIStoreMockRecorder{mock}
	return mock
}

func (m *MockIStore) EXPECT() *MockIStoreMockRecorder {
	return m.recorder
}

func (m *MockIStore) DB() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *MockIStoreMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockIStore)(nil).DB))
}

func (m *MockIStore) Users() IUserStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(IUserStore)
	return ret0
}

func (mr *MockIStoreMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockIStore)(nil).Users))
}

// MockUserStore 是IUserStore的mock实现.
type MockUserStore struct {
	ctrl     *gomock.Controller
	recorder *MockUserStoreMockRecorder
}

type MockUserStoreMockRecorder struct {
	mock *MockUserStore
}

func NewMockUserStore(ctrl *gomock.Controller) *MockUserStore {
	mock := &MockUserStore{ctrl: ctrl}
	mock.recorder = &MockUserStoreMockRecorder{mock}
	return mock
}

func (m *MockUserStore) EXPECT() *MockUserStoreMockRecorder {
	return m.recorder
}

func (m *MockUserStore) Create(arg0 context.Context, arg1 *model.User) error {
	return nil
}

func (m *MockUserStore) Delete(arg0 context.Context, arg1 string) error {
	return nil
}

func (m *MockUserStore) Get(arg0 context.Context, arg1 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserStore)(nil).Get), arg0, arg1)
}

func (m *MockUserStore) List(arg0 context.Context, arg1, arg2 int) (int64, []*model.User, error) {
	return 0, nil, nil
}

func (m *MockUserStore) Update(arg0 context.Context, arg1 *model.User) error {
	return nil
}
