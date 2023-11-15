// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rppkg/torchpole.

package torchpole

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/rppkg/torchpole/internal/pkg/core"
	"github.com/rppkg/torchpole/internal/pkg/errcode"
	"github.com/rppkg/torchpole/internal/pkg/log"
	"github.com/rppkg/torchpole/internal/pkg/middleware"
	"github.com/rppkg/torchpole/internal/torchpole/controller/v1/user"
	"github.com/rppkg/torchpole/internal/torchpole/store"
	"github.com/rppkg/torchpole/pkg/auth"
)

// installRouters 安装 torchpole 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errcode.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.Info(c).Msg("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	// 注册 pprof 路由
	pprof.Register(g)

	// 初始化授权器: authz是角色权限控制，middleware.Authn()是token权限控制
	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	// 构建UserController
	uc := user.New(store.S, authz)
	g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)                             // 创建用户
			userv1.PUT(":name/change-password", uc.ChangePassword) // 更改用户密码
			userv1.Use(middleware.Authn(), middleware.Authz(authz))
			userv1.GET(":name", uc.Get)       // 获取用户详情
			userv1.PUT(":name", uc.Update)    // 更新用户
			userv1.GET("", uc.List)           // 列出用户列表，只有 root 用户才能访问
			userv1.DELETE(":name", uc.Delete) // 删除用户
		}
	}

	return nil
}
