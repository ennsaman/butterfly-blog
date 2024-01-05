package config

var Conf Config

// Config 配置文件对应结构体
type Config struct {
	Server  Server
	MySQL   MySQL
	Redis   Redis
	Session Session
	JWT     JWT
	Zap     Zap
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

type Redis struct {
	Addr     string
	Password string
	DB       int
}

type Session struct {
	Name       string
	Salt       string
	ExpireTime int64
}

type JWT struct {
	Issuer     string
	SecretKey  string
	ExpireTime int64
}

type Zap struct {
	Level        string
	Format       string
	Prefix       string
	Directory    string
	ShowLine     bool
	LogInConsole bool
}
