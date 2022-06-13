package favorite

import (
	"douyin/app/model"
	"douyin/boot/global"
	JwtLib "douyin/utils/jwt"
	"gorm.io/gorm"
	"time"

	//JwtLib "douyin/utils/jwt"
	"fmt"
)

type Favorite struct{}

func (f *Favorite) GetFavoriteList(favorite *model.Fav) (list []model.VideoList) {
	//fmt.Println(favorite)
	//err = global.DYDB.Preload("Videos", "video_id").Find($)

	var resArr []model.VideoList
	//rows, err := global.DYDB.Debug().Table("like_videos").Select("like_videos.video_id,videos.play_url,videos.cover_url,videos.favorite_count,videos.comment_count,videos.title").Joins("left join videos on like_videos.video_id = videos.id").Where("like_videos.user_id=?", favorite.UserId).Order("like_videos.like_time DESC").Rows()
	rows, err := global.DYDB.Debug().Table("like_videos").Select("like_videos.video_id,videos.play_url,videos.cover_url,videos.favorite_count,videos.comment_count,videos.title,like_videos.is_favorite,users.name,users.id,users.follow_count,users.follower_count,users.is_follow").Joins("left join videos on like_videos.video_id = videos.id").Joins("left join users on like_videos.user_id = users.id").Where("like_videos.user_id=? And is_favorite=?", favorite.UserId, 1).Order("like_videos.like_time DESC").Rows()

	if err != nil {
		fmt.Println(111)
		fmt.Println(err)
	}
	for rows.Next() {
		var res model.VideoList
		err = rows.Scan(&res.Id, &res.PlayUrl, &res.CoverUrl, &res.FavoriteCount, &res.CommentCount, &res.Title, &res.IsFavorite, &res.Author.Name, &res.Author.Id, &res.Author.FollowCount, &res.Author.FollowerCount, &res.Author.IsFollow)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
		resArr = append(resArr, res)
	}
	defer rows.Close()
	//fmt.Println(result)
	fmt.Println(resArr)
	return resArr
}

//db.Where("state = ?", "active").Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
//// SELECT * FROM users WHERE state = 'active';
//// SELECT * FROM orders WHERE user_id IN (1,2) AND state NOT IN ('cancelled');

func (f *Favorite) Action(res model.Action, token string) bool {
	//
	jwtData, err := JwtLib.JwtParse(token)
	if err != nil {
		fmt.Println(err)
	}
	uid := jwtData.Uid

	var LikeVideos model.LikeVideos
	LikeVideos.UserId = uid
	LikeVideos.VideoId = res.VideoId
	fmt.Println("状态")
	fmt.Println(res.ActionType)

	var videos model.VideoADD
	videos.Id = res.VideoId
	if res.ActionType == 1 {
		var res1 model.LikeVideos
		LikeVideos.IsFavorite = 1
		global.DYDB.Where("user_id = ? AND video_id = ?", LikeVideos.UserId, LikeVideos.VideoId).Find(&res1)
		if res1.LikeTime == 0 {
			fmt.Println("测试0")
			LikeVideos.LikeTime = time.Now().Unix()
			result := global.DYDB.Create(&LikeVideos)
			fmt.Println(result)
		} else {
			global.DYDB.Model(&LikeVideos).Where("user_id = ? AND video_id = ?", LikeVideos.UserId, LikeVideos.VideoId).Update("is_favorite", "1")
			global.DYDB.Model(&videos).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
			fmt.Println("改为1")
		}
	} else {
		global.DYDB.Model(&LikeVideos).Where("user_id = ? AND video_id = ?", LikeVideos.UserId, LikeVideos.VideoId).Update("is_favorite", "2")
		global.DYDB.Model(&videos).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
	}

	//if is == nil {
	//	LikeVideos.LikeTime = time.Now().Unix()
	//	result := global.DYDB.Create(&LikeVideos)
	//	fmt.Println(result)
	//} else {
	//	LikeVideos.LikeTime = time.Now().Unix()
	//	result := global.DYDB.Create(&LikeVideos)
	//	fmt.Println(result)
	//}
	fmt.Println(LikeVideos)
	//fmt.Println(jwtData)

	return false
}
