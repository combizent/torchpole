// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

package user

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/combizent/torchpole/internal/pkg/core"
	"github.com/combizent/torchpole/internal/pkg/errcode"
	"github.com/combizent/torchpole/internal/pkg/log"
	v1 "github.com/combizent/torchpole/pkg/api/torchpole/v1"
	pb "github.com/combizent/torchpole/pkg/proto/torchpole/v1"
)

// List 返回用户列表，只有 root 用户才能获取用户列表.
func (userController *UserController) List(c *gin.Context) {
	log.Info(c).Msg("List user function called")

	var r v1.ListUserRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errcode.ErrBind, nil)

		return
	}

	resp, err := userController.biz.UserBiz().List(c, r.Offset, r.Limit)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}

// ListUser 返回用户列表，只有 root 用户才能获取用户列表.
func (userController *UserController) ListUser(ctx context.Context, r *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	log.Info(ctx).Msg("ListUser function called")

	resp, err := userController.biz.UserBiz().List(ctx, int(r.Offset), int(r.Limit))
	if err != nil {
		return nil, err
	}

	users := make([]*pb.UserInfo, 0, len(resp.Users))
	for _, u := range resp.Users {
		createdAt, _ := time.Parse("2006-01-02 15:04:05", u.CreatedAt)
		updatedAt, _ := time.Parse("2006-01-02 15:04:05", u.UpdatedAt)
		users = append(users, &pb.UserInfo{
			Username:  u.Username,
			Nickname:  u.Nickname,
			Email:     u.Email,
			Phone:     u.Phone,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		})
	}

	ret := &pb.ListUserResponse{
		TotalCount: resp.TotalCount,
		Users:      users,
	}

	return ret, nil
}
