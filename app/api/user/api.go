package user

import (
	"douyin/app/model"
	userService "douyin/app/service/user"
	"douyin/boot/global"
	JwtLib "douyin/utils/jwt"
	RegexpLib "douyin/utils/regexp"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userResponse := model.UserResp{2, "密码错误", 0, ""}
	if !RegexpLib.RegSome("AZ09", username, 6, 32) || !RegexpLib.RegSome("AZ09", password, 6, 32) {
		c.JSON(http.StatusBadGateway, userResponse)
		return
	}
	user := model.User{}
	result := global.DYDB.Where(map[string]interface{}{"username": username, "password": password}).Find(&user)
	jwt, err := JwtLib.JwtSignByIdName(user.Id, user.Username)
	if err != nil {

	}
	if result.Error != nil || result.RowsAffected == 1 {
		userResponse.StatusCode = 0
		userResponse.StatusMsg = "登陆成功"
		userResponse.User_id = user.Id
		userResponse.Token = jwt
		fmt.Println("token:", jwt)
		c.JSON(http.StatusOK, userResponse)
		return
	} else {
		c.JSON(http.StatusBadGateway, userResponse)
	}
}
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userResponse := model.UserResp{1, "输入有限制", 0, ""}
	if !RegexpLib.RegSome("AZ09", username, 6, 32) || !RegexpLib.RegSome("AZ09", password, 6, 32) {
		c.JSON(http.StatusOK, userResponse)
		return
	}
	user := model.User{0, 0, 0, "test", username, password}

	userResponse = model.UserResp{1, "重复注册", 0, ""}
	result := global.DYDB.Where("username = ?", username).Find(&user)

	if result.Error != nil || result.RowsAffected > 0 {
		c.JSON(http.StatusOK, userResponse)
		return
	}
	global.DYDB.Create(&user)
	userResponse = model.UserResp{0, "登陆成功", 0, ""}
	c.JSON(http.StatusOK, userResponse)
}
func Info(c *gin.Context) {
	user_id := c.Query("user_id")
	//token := c.Query("token")
	//log.Println(user_id)
	//log.Println(token)
	v, err := strconv.Atoi(user_id)
	if err != nil {
		c.JSON(http.StatusOK, nil)
	}
	value, _ := userService.GetUser(uint(v))
	userTest := model.UserFans{
		value[0].FollowCount,
		value[0].FollowerCount,
		value[0].Id,
		value[0].Name,
		false,
	}
	//result := []model.UserFans{userTest}
	userinfo := model.UserinfoResp{0, "登陆成功", userTest}
	//log.Println(userinfo)
	c.JSON(http.StatusOK, userinfo)
}
