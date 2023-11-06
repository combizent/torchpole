// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

package user

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"

	v1 "github.com/combizent/torchpole/pkg/api/torchpole/v1"
)

// MockUserBiz 需要实现IUserBiz接口.
type MockUserBiz struct {
	ctrl     *gomock.Controller
	recorder *MockUserBizMockRecorder
}

type MockUserBizMockRecorder struct {
	mock *MockUserBiz
}

func NewMockUserBiz(ctrl *gomock.Controller) *MockUserBiz {
	mock := &MockUserBiz{ctrl: ctrl}
	mock.recorder = &MockUserBizMockRecorder{mock}
	return mock
}

func (m *MockUserBiz) EXPECT() *MockUserBizMockRecorder {
	return m.recorder
}

func (m *MockUserBiz) ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error {
	return nil
}

func (m *MockUserBiz) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	return nil, nil
}

func (m *MockUserBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	return nil
}

func (m *MockUserBiz) Get(ctx context.Context, username string) (*v1.GetUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, username)
	ret0, _ := ret[0].(*v1.GetUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserBizMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserBiz)(nil).Get), arg0, arg1)
}

func (m *MockUserBiz) List(ctx context.Context, offset, limit int) (*v1.ListUserResponse, error) {
	return nil, nil
}

func (m *MockUserBiz) Update(ctx context.Context, username string, r *v1.UpdateUserRequest) error {
	return nil
}

func (m *MockUserBiz) Delete(ctx context.Context, username string) error {
	return nil
}
