package utils

import (
	"blog-server/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// InitMySQL 初始化 MySQL
func InitMySQL() *gorm.DB {
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

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

// gormConfig gorm 配置
func gormConfig() *gorm.Config {
	return &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info), // 日志级别
		SkipDefaultTransaction: true,                                // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // 表名前缀，“User”的表将为“t_users”
			SingularTable: true,  // 使用单数表名，“User”的表将是“user”，启用此选项后
			NoLowerCase:   false, // 跳过名称的蛇形大小写
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用特性：GORM 在 AutoMigrate 或 CreateTable 时自动创建外键约束
	}
}
