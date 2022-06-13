package favorite

import (
	"douyin/app/model"
	favorite2 "douyin/app/service/favorites"
	"fmt"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)

func List(c *gin.Context) {

	var favorite *model.Fav
	userId, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	favorite = new(model.Fav)
	favorite.Token = token
	favorite.UserId = uint(userId)
	var fav favorite2.Favorite
	varArr := fav.GetFavoriteList(favorite)
	var favoriteList model.FavoriteListResp
	favoriteList.StatusCode = 200
	favoriteList.StatusMsg = "成功"
	favoriteList.VideoList = varArr[:]
	fmt.Println(favoriteList)
	c.JSON(http.StatusOK, favoriteList)
}

func Action(c *gin.Context) {

	video_id, _ := strconv.Atoi(c.Query("video_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))
	token := c.Query("token")
	var res model.Action
	res.ActionType = action_type
	res.VideoId = uint(video_id)
	var fav favorite2.Favorite
	data := fav.Action(res, token)
	fmt.Println(data)
	var ActionResp model.ActionResp
	ActionResp.StatusCode = 0
	ActionResp.StatusMsg = "成功"
	c.JSON(http.StatusOK, ActionResp)
}
