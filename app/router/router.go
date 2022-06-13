package router

import (
	"douyin/app/api"
	"douyin/app/api/action"
	"douyin/app/api/favorite"

	"douyin/app/api/testa"
	"douyin/app/api/user"
	"douyin/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	tag := testa.NewTag()

	r.Use(middleware.LoggerHandler())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word！！！！！！！！！！！")
	})
	test1 := r.Group("/test/v1")
	{
		test1.POST("/get/tag", tag.Get)
		test1.POST("/post/data", tag.PostData)
	}
	DouyinApi := r.Group("/douyin")
	{
		//用户操作
		//DouyinApi.GET("user",user.Info)
		UserApi := DouyinApi.Group("user")
		{
			UserApi.POST("/login/", user.Login)
			UserApi.POST("/register/", user.Register)
			UserApi.GET("/", user.Info).Use(middleware.TokenCheck())
		}
		CommentApi := DouyinApi.Group("comment")
		{
			CommentApi.POST("/action/", action.AddAction).Use(middleware.TokenCheck())
			CommentApi.GET("/list/", action.ListAction).Use(middleware.TokenCheck())

		}

	}
	douyin := r.Group("/douyin")
	{
		douyin.GET("/feed", api.Getfeed)
		douyin.POST("/publish/action/", api.PublishAction)
		douyin.GET("/publish/list", api.GetPublicList)

	}
	//点赞
	zan := r.Group("/douyin")
	{
		zan.GET("/favorite/list", favorite.List)
		zan.POST("/favorite/action/", favorite.Action)
	}

	return r
}
