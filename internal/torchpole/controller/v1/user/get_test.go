// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/combizent/torchpole/internal/torchpole/biz"
	"github.com/combizent/torchpole/internal/torchpole/biz/user"
	v1 "github.com/combizent/torchpole/pkg/api/torchpole/v1"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func TestUserController_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//  需要mock层的接口，返回的定义
	want := &v1.GetUserResponse{
		Username:  "zhangsan",
		Nickname:  "zhangsan",
		Email:     "zhangsan@qq.com",
		Phone:     "18888888888",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	mockUserBiz := user.NewMockUserBiz(ctrl)
	mockBiz := biz.NewMockIBiz(ctrl)
	// 对IUserBiz接口层的mock， 通过调用Get获取Get方法的返回结果，结果被定义为want
	mockUserBiz.EXPECT().Get(gomock.Any(), gomock.Any()).Return(want, nil).Times(1)
	// 对IBiz接口层的Mock, 通过调用UserBiz获取具体业务接口IUserBiz的实现
	mockBiz.EXPECT().UserBiz().AnyTimes().Return(mockUserBiz)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/v1/users/zhangsan", nil)
	c.Request.Header.Set("Content-Type", "application/json")

	blw := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}
	c.Writer = blw

	type fields struct {
		b biz.IBiz
	}

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.GetUserResponse
	}{
		{
			name:   "default",
			fields: fields{b: mockBiz},
			args: args{
				c: c,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userController := &UserController{
				biz: tt.fields.b,
			}
			userController.Get(tt.args.c)

			var resp v1.GetUserResponse
			err := json.Unmarshal(blw.body.Bytes(), &resp)
			assert.Nil(t, err)
			assert.Equal(t, resp.Username, want.Username)
		})
	}
}
