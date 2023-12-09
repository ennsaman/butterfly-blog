package router

import (
	"blog-server/config"
	"blog-server/utils"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var DB *gorm.DB

// InitGlobalVariable 初始化全局变量
func InitGlobalVariable() {
	utils.InitViper()
	DB = utils.InitMySQL()
	_, err := DB.DB()
	if err != nil {
		fmt.Println("get *sql.DB fail：", err)
	}
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
