// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rppkg/torchpole.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/rppkg/torchpole/internal/pkg/core"
	"github.com/rppkg/torchpole/internal/pkg/errcode"
	"github.com/rppkg/torchpole/internal/pkg/log"
	v1 "github.com/rppkg/torchpole/pkg/api/torchpole/v1"
)

// Update 更新用户信息.
func (userController *UserController) Update(c *gin.Context) {
	log.Info(c).Msg("Update user function called")

	var r v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errcode.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errcode.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := userController.biz.UserBiz().Update(c, c.Param("name"), &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
