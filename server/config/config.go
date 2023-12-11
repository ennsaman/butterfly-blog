package config

var Conf Config

// Config 配置文件对应结构体
type Config struct {
	Server Server
	MySQL  MySQL
	JWT    JWT
}

type Server struct {
	AppMode   string
	BackPort  string
	FrontPort string
}

type MySQL struct {
	Host     string
	Port     string
	DataBase string
	UserName string
	Password string
}

type JWT struct {
	Issuer     string
	SecretKey  string
	ExpireTime int64
}
