package userService

import (
	"douyin/app/model"
	"douyin/boot/global"
	"gorm.io/gorm"
)

func GetUser(user_id uint) ([]model.User, *gorm.DB) {
	var user []model.User
	result := global.DYDB.Where("id = ?", user_id).Find(&user)
	for i, _ := range user {
		user[i].Password = ""
	}
	return user, result
}
