// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

package torchpole

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/combizent/torchpole/internal/pkg/core"
	"github.com/combizent/torchpole/internal/pkg/log"
	"github.com/combizent/torchpole/internal/pkg/middleware"
	"github.com/combizent/torchpole/pkg/token"
	"github.com/combizent/torchpole/pkg/version/verflag"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "torchpole",

		Short: "The brightest star in the night sky",
		Long: `The brightest star in the night sky

More information at:
	https://github.com/combizent/torchpole#readme`,

		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			// 如果 `--version=true`，则打印版本并退出
			verflag.PrintAndExitIfRequested()

			log.Init(logOptions())

			return run()
		},

		// 不添加命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q",
						cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)

	// 在这里您将定义标志和配置设置。

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "",
		"The path to the torchpole configuration file. Empty string for no configuration file.")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 添加 --version 标志
	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

func run() error {
	// 初始化 store 层
	if err := initStore(); err != nil {
		return err
	}

	// 设置 token 包的签发密钥，用于 token 包 token 的签发和解析
	token.Init(viper.GetString("jwt-secret"), core.XUsernameKey)

	// 设置 Gin 模式
	gin.SetMode(viper.GetString("mode"))

	// 创建 Gin 引擎
	g := gin.New()

	// gin.Recovery() 中间件，用来捕获任何 panic，并恢复
	mws := []gin.HandlerFunc{
		gin.Recovery(), middleware.NoCache, middleware.Cors,
		middleware.Secure, middleware.RequestID(),
	}
	g.Use(mws...)

	// 初始化路由
	if err := installRouters(g); err != nil {
		return err
	}

	// 启动HTTP服务器
	httpsrv := startInsecureServer(g)

	// 创建并运行 HTTPS 服务器
	httpssrv := startSecureServer(g)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	// 创建 ctx 用于通知服务器 goroutine, 它有 10 秒时间完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 10 秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过 10 秒就超时退出
	if err := httpsrv.Shutdown(ctx); err != nil {
		log.ErrWithoutCtx(err).Msg("Insecure Server forced to shutdown")
		return err
	}

	if err := httpssrv.Shutdown(ctx); err != nil {
		log.ErrWithoutCtx(err).Msg("Secure Server forced to shutdown")
		return err
	}

	return nil
}

// startInsecureServer 创建并运行 HTTP 服务器.
func startInsecureServer(g *gin.Engine) *http.Server {
	// 创建 HTTP Server 实例
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	// 运行 HTTP 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	// 打印一条日志，用来提示 HTTP 服务已经起来，方便排障
	log.InfoWithoutCtx().Str("addr", viper.GetString("addr")).Msg("Start to listening the incoming requests on http address")
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.FatalWithoutCtx().Err(err)
		}
	}()

	return httpsrv
}

// startSecureServer 创建并运行 HTTPS 服务器.
func startSecureServer(g *gin.Engine) *http.Server {
	// 创建 HTTPS Server 实例
	httpssrv := &http.Server{Addr: viper.GetString("tls.addr"), Handler: g}

	// 运行 HTTPS 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	// 打印一条日志，用来提示 HTTPS 服务已经起来，方便排障
	log.InfoWithoutCtx().Str("addr", viper.GetString("tls.addr")).Msg("Start to listening the incoming requests on https address")
	cert, key := viper.GetString("tls.cert"), viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			if err := httpssrv.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.FatalWithoutCtx().Err(err)
			}
		}()
	}

	return httpssrv
}
