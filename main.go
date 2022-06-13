package main

import (
	"douyin/app/router"
	"douyin/boot/global"
	"douyin/boot/viper"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = GormMysql1()
	if err != nil {
		log.Fatalf("init.GormMysql1 err: %v", err)
	}

}

func main() {
	//r := gin.Default()
	gin.SetMode(gin.DebugMode)
	viper.NewSetting()
	global.DYDB = global.GormMysql()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	r := router.NewRouter()
	r.Run(global.DyServer.RunAddress + ":" + global.DyServer.HttpPort) // listen and serve on 0.0.0.0:8080
}

func setupSetting() error {
	_, err := viper.NewSetting()
	if err != nil {
		return nil
	}

	return nil
}

func GormMysql1() error {
	global.DYDB = global.GormMysql()
	return nil
}
