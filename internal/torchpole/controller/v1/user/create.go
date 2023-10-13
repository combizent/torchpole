// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/combizent/torchpole/internal/pkg/core"
	"github.com/combizent/torchpole/internal/pkg/errcode"
	"github.com/combizent/torchpole/internal/pkg/log"
	v1 "github.com/combizent/torchpole/pkg/api/torchpole/v1"
)

const defaultMethods = "(GET)|(POST)|(PUT)|(DELETE)"

// Create 创建一个新的用户.
func (ctrl *UserController) Create(c *gin.Context) {
	log.Info(c).Msg("Create user function called")

	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errcode.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errcode.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.biz.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	if _, err := ctrl.auth.AddNamedPolicy("p", r.Username, "/v1/users/"+r.Username, defaultMethods); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
