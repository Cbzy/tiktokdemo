package action

import (
	"douyin/app/model"
	"douyin/boot/global"
	JwtLib "douyin/utils/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func AddAction(c *gin.Context) {
	jwtvalue, _ := JwtLib.JwtParse(c.Query("token"))
	user := model.User{}
	result := global.DYDB.Where(map[string]interface{}{"id": jwtvalue.Uid}).Find(&user)
	log.Println(result)
	action := model.CommitAction{0, c.Query("comment_text"), time.Now().Format("2006-01-02 15:04:05.000"), c.Query("video_id"), user.Id}
	global.DYDB.Create(&action)
	log.Println(action)
	comment := model.CommitActionPush{action.Id, action.Content, action.Create_date, &user}
	pushResult := model.CommitResp{0, "评论成功", &comment}
	c.JSON(http.StatusOK, pushResult)
}

func ListAction(c *gin.Context) {
	jwtvalue, _ := JwtLib.JwtParse(c.Query("token"))
	action := []model.CommitAction{}
	result := global.DYDB.Where(map[string]interface{}{"uid": jwtvalue.Uid, "video_id": c.Query("video_id")}).Find(&action)
	log.Println(result)
	var list []model.CommitActionPush
	for _, value := range action {
		tmp := value
		user := model.User{}
		userResult := global.DYDB.Where(map[string]interface{}{"id": tmp.Uid}).Find(&user)
		if userResult.Error == nil {
			list = append(list, model.CommitActionPush{value.Id, value.Content, value.Create_date, &user})
		}
	}
	//user := model.User{}
	//result := global.DYDB.Where(map[string]interface{}{"id": jwtvalue.Uid}).Find(&user)
	//log.Println(result)
	//action := model.CommitAction{0, c.Query("comment_text"), time.Now().Format("2006-01-02 15:04:05.000")}
	//global.DYDB.Create(&action)
	//log.Println(action)
	//comment := model.CommitActionPush{action.Id, action.Content, action.Create_date, &user}
	//pushResult := model.CommitResp{0, "评论成功", &comment}
	c.JSON(http.StatusOK, model.CommitActionListPush{0, "查询评论", list})
}
func CheckNumber(uid, videoid int) int {
	action := []model.CommitAction{}
	result := global.DYDB.Where(map[string]interface{}{"uid": uid, "video_id": videoid}).Find(&action)
	if result.Error == nil {
		return len(action)
	}
	return 0

}
