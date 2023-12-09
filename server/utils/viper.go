package utils

import (
	"blog-server/config"
	"github.com/spf13/viper"
	"log"
)

// InitViper 初始化 viper 配置
func InitViper() {
	v := viper.New()
	v.SetConfigName("config")    // 指定配置文件名
	v.SetConfigType("yaml")      // 指定配置文件类型
	v.AddConfigPath("./config")  // main指定查找配置文件的路径
	v.AddConfigPath("../config") // 测试目录查找配置文件的路径

	err := v.ReadInConfig() // 读取配置文件

	if err != nil {
		log.Panic("读取配置文件失败：", err)
	}

	// 将读取的配置文件保存到全局变量Conf
	if err := v.Unmarshal(&config.Conf); err != nil {
		log.Panic("解析配置文件失败：", err)
	}

	log.Println("解析配置文件成功！")
}
