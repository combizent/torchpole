// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

package user

import (
	"github.com/combizent/torchpole/internal/torchpole/biz"
	"github.com/combizent/torchpole/internal/torchpole/store"
	"github.com/combizent/torchpole/pkg/auth"
	pb "github.com/combizent/torchpole/pkg/proto/torchpole/v1"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	auth *auth.Authz
	biz  biz.IBiz
	pb.UnimplementedTorchPoleServer
}

// New 创建一个 user controller.
func New(s store.IStore, a *auth.Authz) *UserController {
	return &UserController{auth: a, biz: biz.NewBiz(s)}
}
