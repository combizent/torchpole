// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rppkg/torchpole.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 可以升级为gorilla/mux
	mux := http.NewServeMux()
	mux.Handle("/hello", &helloHandler{})

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    30 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ErrorLog:       nil,     // 使用默认日志记录器
	}

	// 创建系统信号接收器
	done := make(chan os.Signal, 1)
	// os.Interrupt 和 syscall.SIGINT 都表示 Ctrl+C 中断信号.
	// syscall.SIGTERM 则表示程序终止请求。kill -15
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()

	log.Println("Starting HTTP server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}

type helloHandler struct{}

func (*helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	contentType := r.Header.Get("Content-Type")

	fmt.Println(hasFirst, first, hasSecond, second, contentType)

	w.Header().Set("Hello", "World")
	fmt.Fprint(w, "OK.")
}
