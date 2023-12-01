package utils

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitMySQL() {

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
