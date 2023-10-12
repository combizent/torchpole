// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

package main

import (
	"os"

	_ "go.uber.org/automaxprocs"

	"github.com/combizent/torchpole/internal/torchpole"
)

func main() {
	if err := torchpole.NewCmd().
		Execute(); err != nil {
		os.Exit(1)
	}
}
