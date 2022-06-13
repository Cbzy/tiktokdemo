package service

import (
	"douyin/app/model"
	"douyin/boot/global"
	JwtLib "douyin/utils/jwt"
	"errors"
	"fmt"
	"strconv"
)

func Getfeed(latestTime int64, token string) ([]model.VideoResp, error) {
	var uid int
	fmt.Println(token)
	if len(token) == 0 { //未登录
		uid = 0
	} else { //已经登陆解析token
		uc, err := JwtLib.JwtParse(token)
		if err != nil {
			fmt.Println("用户信息获取失败")
			return nil, err
		}
		uid = int(uc.Uid)
	}

	// 	SELECT
	// videos.public_time,videos.comment_count,videos.cover_url,videos.favorite_count,videos.id as v_id,videos.play_url,videos.title,
	// users.follow_count,users.follower_count,users.id as u_id,users.`name`,(select count(*) from follow where follow.follower_user_id=100001 and videos.user_id=follow.follow_user_id and follow.is_follow=1) as is_follow,(select count(1) from like_videos where like_videos.user_id=100001 and like_videos.video_id=videos.id and is_favorite=1) as is_favorite
	// FROM `videos` left join users on users.id = videos.user_id
	// WHERE videos.public_time<1655002477644 ORDER BY videos.public_time DESC LIMIT 30

	rows, err := global.DYDB.Debug().Table("videos").Select("videos.public_time,videos.comment_count,videos.cover_url,videos.favorite_count,videos.id as v_id,videos.play_url,videos.title,users.follow_count,users.follower_count,users.id as u_id,users.`name`,(select count(*) from follow where follow.follower_user_id="+strconv.Itoa(uid)+" and videos.user_id=follow.follow_user_id and follow.is_follow=1) as is_follow,(select count(1) from like_videos where like_videos.user_id="+strconv.Itoa(uid)+" and like_videos.video_id=videos.id and is_favorite=1) as is_favorite").Joins("left join users on users.id = videos.user_id").Where("videos.public_time<?", latestTime).Order("videos.public_time DESC").Limit(30).Rows()
	if err != nil {
		fmt.Println("视频流查询失败")
		return nil, err
	}
	defer rows.Close()
	var videoArr []model.VideoResp
	for rows.Next() {
		var video model.VideoResp

		err = rows.Scan(&video.PublicTime, &video.CommentCount, &video.CoverURL, &video.FavoriteCount, &video.ID, &video.PlayURL, &video.Title, &video.Author.FollowCount, &video.Author.FollowerCount, &video.Author.Id, &video.Author.Name, &video.Author.Is_follow, &video.IsFavorite)
		if err != nil {
			fmt.Println("数据绑定失败")
			return nil, err
		}
		//fmt.Println(video)
		//fmt.Println(video.Author)
		fmt.Println(video.PublicTime)

		videoArr = append(videoArr, video)
	}

	return videoArr, nil
}

func PublishList(userId string, token string) ([]model.VideoResp, error) {

	if isNum(userId) == false {
		return nil, errors.New("用户id含有非法字符")
	}

	//查询
	rows, err := global.DYDB.Debug().Table("videos").Select("videos.comment_count,videos.cover_url,videos.favorite_count,videos.id as v_id,videos.play_url,videos.title,users.follow_count,users.follower_count,users.id as u_id,users.`name`,(select count(*) from follow where follow.follower_user_id=" + userId + " and videos.user_id=follow.follow_user_id and follow.is_follow=1) as is_follow,(select count(1) from like_videos where like_videos.user_id=" + userId + " and like_videos.video_id=videos.id and is_favorite=1) as is_favorite").Joins("left join users on users.id = videos.user_id").Where("videos.user_id=" + userId).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videoArr []model.VideoResp
	for rows.Next() {
		var video model.VideoResp

		err = rows.Scan(&video.CommentCount, &video.CoverURL, &video.FavoriteCount, &video.ID, &video.PlayURL, &video.Title, &video.Author.FollowCount, &video.Author.FollowerCount, &video.Author.Id, &video.Author.Name, &video.Author.Is_follow, &video.IsFavorite)
		if err != nil {
			fmt.Println("数据绑定失败")
			return nil, err
		}

		videoArr = append(videoArr, video)
	}

	//封装返回值
	return videoArr, nil
}

func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func PublishAction() error {
	
	
	return nil
}
