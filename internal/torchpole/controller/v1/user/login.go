// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rppkg/torchpole.

package user

import (
	"github.com/gin-gonic/gin"

	"github.com/rppkg/torchpole/internal/pkg/core"
	"github.com/rppkg/torchpole/internal/pkg/errcode"
	"github.com/rppkg/torchpole/internal/pkg/log"
	v1 "github.com/rppkg/torchpole/pkg/api/torchpole/v1"
)

// Login 登录 torchpole 并返回一个 JWT Token.
func (userController *UserController) Login(c *gin.Context) {
	log.Info(c).Msg("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errcode.ErrBind, nil)

		return
	}

	resp, err := userController.biz.UserBiz().Login(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}
