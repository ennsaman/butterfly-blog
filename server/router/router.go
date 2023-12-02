package router

import "blog-server/utils"

// 初始化全局变量
func InitGlobalVariable() {
	utils.InitViper()
	db := utils.InitMySQL()
	db.DB()
}
