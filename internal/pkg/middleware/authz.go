// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rppkg/torchpole.

package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/rppkg/torchpole/internal/pkg/core"
	"github.com/rppkg/torchpole/internal/pkg/errcode"
	"github.com/rppkg/torchpole/internal/pkg/log"
)

// Auther 用来定义授权接口实现.
// sub: 操作主题，obj：操作对象, act：操作
type Auther interface {
	Authorize(sub, obj, act string) (bool, error)
}

// Authz 是 Gin 中间件，用来进行请求授权.
func Authz(a Auther) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetString(core.XUsernameKey)
		obj := c.Request.URL.Path
		act := c.Request.Method

		log.DebugWithoutCtx().Str("sub", sub).Str("obj", obj).Str("act", act).Msg("Build authorize context")
		if allowed, _ := a.Authorize(sub, obj, act); !allowed {
			core.WriteResponse(c, errcode.ErrUnauthorized, nil)
			c.Abort()
			return
		}
	}
}
