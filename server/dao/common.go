package dao

import (
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// Create 通用创建
func Create[T any](data *T) {
	err := DB.Create(&data).Error
	if err != nil {
		log.Println("创建失败：", err)
		panic(err)
	}
}

// GetOne 通用查询单行
func GetOne[T any](data T, query string, args ...any) T {
	err := DB.Where(query, args...).Limit(1).Find(&data).Error
	if err != nil {
		log.Println("查询失败：", err)
		panic(err)
	}
	return data
}

// Update 通用更新单行
func Update[T any](data *T, cols ...string) {
	// 判断是否传入字段名参数
	if len(cols) == 0 {
		DB.Model(&data).Updates(&data)
		return
	}
	err := DB.Model(&data).Select(cols).Updates(&data).Error
	if err != nil {
		log.Println("更新失败：", err)
		panic(err)
	}
}

// UpdateMap 更新 map
func UpdateMap[T any](data *T, mp map[string]any, query string, args ...any) {
	err := DB.Model(&data).Where(query, args).Updates(mp).Error
	if err != nil {
		log.Println("更新失败：", err)
		panic(err)
	}
}

// Delete 通用删除
func Delete[T any](data T, query string, args ...any) {
	err := DB.Where(query, args).Delete(&data).Error
	if err != nil {
		log.Println("删除失败：", err)
		panic(err)
	}
}

// Count 通用计数
func Count[T any](data T, query string, args ...any) int64 {
	var count int64
	db := DB.Model(&data)
	if query == "" {
		db = db.Where(query, args)
	}
	if err := db.Count(&count).Error; err != nil {
		log.Println("计数失败：", err)
		panic(err)
	}
	return count
}
