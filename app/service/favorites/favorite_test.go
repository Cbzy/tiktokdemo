package favorite

import (
	"douyin/boot/global"
	"gorm.io/gorm"
	"testing"
)

type Author struct {
	Id            uint   `json:"id"`             //用户id
	Name          string `json:"name"`           //用户名称
	FollowCount   uint   `json:"follow_count"`   //关注总数
	FollowerCount uint   `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      //是否关注
}

func TestName(t *testing.T) {
	var entities []Author
	err := global.DYDB.Select("id").Preload("Videos", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "title")
	}).Find(&entities).Error
	if err != nil {
		t.Error(err)
		return
	}
	return
}
