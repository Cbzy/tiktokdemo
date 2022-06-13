package main

import (
	userService "douyin/app/service/user"
	"douyin/boot/global"
	"douyin/boot/viper"
	"log"
	"testing"
)

func TestGetUser(t *testing.T) {
	viper.NewSetting()
	global.DYDB = global.GormMysql()
	//log.Println(global.DYDB)
	a, b := userService.GetUser(100003)
	log.Println(a)
	log.Println(b)
}
