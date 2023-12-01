package routes

import "blog-server/utils"

// 初始化全局变量
func InitGlobalVariable() {
	utils.InitViper()
	utils.InitMySQL()
}
