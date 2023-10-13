// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/combizent/torchpole/internal/pkg/core"
	"github.com/combizent/torchpole/internal/pkg/errcode"
	"github.com/combizent/torchpole/pkg/token"
)

// Authn 是认证中间件，用来从 gin.Context 中提取 token 并验证 token 是否合法，
// 如果合法则将 token 中的 sub 作为<用户名>存放在 gin.Context 的 XUsernameKey 键中.
func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errcode.ErrTokenInvalid, nil)
			c.Abort()

			return
		}

		c.Set(core.XUsernameKey, username)
		c.Next()
	}
}
