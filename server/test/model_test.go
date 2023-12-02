package test

import (
	"blog-server/model"
	"blog-server/utils"
	"fmt"
	"gorm.io/gorm"
	"log"
	"testing"
)

var db *gorm.DB

func init() {
	utils.InitViper()
	db = utils.InitMySQL()
}

func TestInsertUserInfoByOne(t *testing.T) {

	userInfo := &model.UserInfo{
		Nickname: "张三",
		Avatar:   "zhangsan avatar",
		Email:    "zhangsan@qq.com",
		Intro:    "我是张三",
	}
	res := db.Select("Nickname", "Avatar", "Email", "Intro").Create(userInfo)

	err := res.Error
	if err != nil {
		log.Fatal("插入数据失败：", err)
	}
	log.Println("插入数据成功，影响行数：", res.RowsAffected)
}

func TestInsertUserInfoByList(t *testing.T) {
	userInfoList := []*model.UserInfo{
		{Nickname: "zs", Avatar: "zs"},
		{Nickname: "ls", Avatar: "ls"},
	}

	// 切片返回值是第一个元素的指针，所以这里可以不用传入指针
	res := db.Select("Nickname", "Avatar").Create(userInfoList)

	err := res.Error
	if err != nil {
		log.Fatal("插入多条数据失败：", err)
	}
	log.Println("插入数据成功，影响行数：", res.RowsAffected)
	log.Println(userInfoList[0].Id, userInfoList[0].CreateTime) // 插入会返回数据的主键id， 其他的不能获取到
}

func TestSelectUserInfoByOne(t *testing.T) {
	userInfo := &model.UserInfo{Id: 1}
	res := db.Where("id = ? AND nickname = ?", "1", "管理员").Find(userInfo)
	err := res.Error
	if err != nil {
		log.Fatal("查询数据失败：", err)
	}
	log.Println("查询数据成功，影响行数：", res.RowsAffected)
	log.Printf("%d %s %s %s", userInfo.Id, userInfo.Nickname, userInfo.Avatar, userInfo.CreateTime)

}

func TestSelectUserInfoByList(t *testing.T) {
	userInfoList := []*model.UserInfo{}
	res := db.Find(&userInfoList)
	err := res.Error
	if err != nil {
		log.Fatal("查询数据失败：", err)
	}
	log.Println("查询列表成功！")
	for i := 0; i < len(userInfoList); i++ {
		fmt.Printf("id: %d, name: %s\n", userInfoList[i].Id, userInfoList[i].Nickname)
	}

}

func TestUpdateUserInfoByOne(t *testing.T) {
	userInfo := &model.UserInfo{}
	userInfo.Nickname = "zhangsan"
	userInfo.Avatar = "wangwu avatar"
	userInfo.Email = "wangwu@qq.com"

	res := db.Model(&userInfo).Where("id", "37").Updates(map[string]interface{}{"nickname": "zhangsan", "avatar": "zhangsan avatar"})
	err := res.Error
	if err != nil {
		log.Fatal("更新失败：", err)
	}
	log.Println("更新成功！")
	log.Printf("%d %s %s", userInfo.Id, userInfo.Nickname, userInfo.Email)
}

func TestDeleteUserInfo(t *testing.T) {
	userInfo := &model.UserInfo{}
	userInfo.Id = 46
	userInfo.Nickname = "zs"

	res := db.Where("nickname = ?", userInfo.Nickname).Delete(&userInfo)
	err := res.Error
	if err != nil {
		log.Fatal("删除失败：", err)
	}
	log.Println("删除成功！")
}
