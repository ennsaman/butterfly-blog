package utils

import (
	"blog-server/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func InitMySQL() {
	mysqlConf := config.Conf.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.UserName,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DataBase,
	)
	db, err := gorm.Open(mysql.Open(dsn), gormConfig())
	if err != nil {
		log.Fatal("MySQL 连接失败：", err)
	}
	log.Println("MySQL 连接成功！")

	// 数据库连接数量相关设置
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

func gormConfig() *gorm.Config {
	return &gorm.Config{
		SkipDefaultTransaction: true, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，“User”的表将为“t_users”
			SingularTable: true, // 使用单数表名，“User”的表将是“user”，启用此选项后
			NoLowerCase:   true, // 跳过名称的蛇形大小写
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用特性：GORM 在 AutoMigrate 或 CreateTable 时自动创建外键约束
	}
}
