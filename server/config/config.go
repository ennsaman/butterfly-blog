package config

var Conf Config

// 配置文件对应结构体
type Config struct {
	Server Server
	MySQL  MySQL
}

type Server struct {
	AppMode   string
	BackPort  string
	FrontPort string
}

type MySQL struct {
	Host     string
	Port     string
	DateBase string
	UserName string
	Password string
}
