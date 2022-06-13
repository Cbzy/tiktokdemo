package testa

import (
	"douyin/app/model/testm"
	"douyin/app/service/tests"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	var tags testm.Tags
	if err := c.ShouldBindJSON(&tags); err != nil {
		fmt.Println(err)
	}

	var ts tests.TagService
	if err := ts.CreateTag(tags); err != nil {

		c.JSON(500, gin.H{"message": "创建失败"})
		return
	}
	c.JSON(200, gin.H{"message": "pong1", "data": tags})
	return

}

func (t Tag) PostData(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
