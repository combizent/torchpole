// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

package biz

import (
	"reflect"

	"github.com/golang/mock/gomock"

	"github.com/combizent/torchpole/internal/torchpole/biz/user"
)

type MockIBiz struct {
	ctrl     *gomock.Controller
	recorder *MockIBizMockRecorder
}

type MockIBizMockRecorder struct {
	mock *MockIBiz
}

func NewMockIBiz(ctrl *gomock.Controller) *MockIBiz {
	mock := &MockIBiz{ctrl: ctrl}
	mock.recorder = &MockIBizMockRecorder{mock}
	return mock
}

func (m *MockIBiz) EXPECT() *MockIBizMockRecorder {
	return m.recorder
}

// UserBiz mocks base method.
func (m *MockIBiz) UserBiz() user.IUserBiz {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserBiz")
	ret0, _ := ret[0].(user.IUserBiz)
	return ret0
}

// UserBiz indicates an expected call of Users.
func (mr *MockIBizMockRecorder) UserBiz() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserBiz", reflect.TypeOf((*MockIBiz)(nil).UserBiz))
}
