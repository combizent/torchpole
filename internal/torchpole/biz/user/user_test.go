// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

package user

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"

	"github.com/combizent/torchpole/internal/pkg/model"
	"github.com/combizent/torchpole/internal/torchpole/store"
	v1 "github.com/combizent/torchpole/pkg/api/torchpole/v1"
)

func Test_userBiz_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeUser := &model.User{
		ID:        1,
		Username:  "zhangsan",
		Password:  "zhangsan123",
		Nickname:  "zhangsan",
		Email:     "zhangsan@qq.com",
		Phone:     "18888888888",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(fakeUser, nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().AnyTimes().Return(mockUserStore)

	var want v1.GetUserResponse
	_ = copier.Copy(&want, fakeUser)
	want.CreatedAt = fakeUser.CreatedAt.Format("2006-01-02 15:04:05")
	want.UpdatedAt = fakeUser.UpdatedAt.Format("2006-01-02 15:04:05")

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.GetUserResponse
	}{
		{name: "default", fields: fields{ds: mockStore}, args: args{context.Background(), "zhangsan"}, want: &want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			got, err := b.Get(tt.args.ctx, tt.args.username)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
