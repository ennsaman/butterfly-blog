package main

import (
	"blog-server/router"
	"golang.org/x/sync/errgroup"
	"log"
)

var g errgroup.Group

func main() {
	// 初始化全局变量
	router.InitGlobalVariable()

	// 启动后台服务
	g.Go(func() error {
		return router.BackendServer().ListenAndServe()
	})

	// 启动前台服务
	g.Go(func() error {
		return router.FrontendServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal("服务运行发生错误：", err)
	}
}
