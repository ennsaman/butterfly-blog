package router

import (
	"blog-server/config"
	"blog-server/dao"
	"blog-server/utils"
	"log"
	"net/http"
	"time"
)

// InitGlobalVariable 初始化全局变量
func InitGlobalVariable() {
	// 初始化 Viper
	utils.InitViper()
	// 初始化 Logger
	utils.InitLogger()
	// 初始化 Redis
	//utils.InitRedis()
	// 初始化数据库 DB
	dao.DB = utils.InitMySQL()

}

// BackendServer 后台服务
func BackendServer() *http.Server {
	backPort := config.Conf.Server.BackPort
	log.Printf("后台服务启动于 %s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      BackRouter(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}
}

// FrontendServer 前台服务
func FrontendServer() *http.Server {
	frontPort := config.Conf.Server.FrontPort
	log.Printf("前台服务启动于 %s 端口", frontPort)
	return &http.Server{
		Addr:         frontPort,
		Handler:      FrontRouter(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}
}
