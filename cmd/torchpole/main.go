// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

package main

import (
	_ "go.uber.org/automaxprocs"

	"github.com/combizent/torchpole/internal/torchpole"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "tp-cli",
	Short: "Torchpole CLI.",
}

func main() {
	cmd.AddCommand(torchpole.CmdRun)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
