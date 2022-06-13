package model

type FavoriteListResp struct {
	StatusCode int64       `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string      `json:"status_msg"`  // 返回状态描述
	VideoList  []VideoList `json:"video_list"`  //用户点赞视频列表
}

type VideoList struct {
	Id            uint   //视频唯一标识
	Author        Author `json:"author"`         // 视频作者信息
	PlayUrl       string `json:"play_url"`       // 播放地址
	CoverUrl      string `json:"cover_url"`      //视频封面地址
	FavoriteCount uint   `json:"favorite_count"` //视频点赞总数
	CommentCount  uint   `json:"comment_count"`  //视频评论总数
	IsFavorite    bool   `json:"is_favorite"`    //是否点赞
	Title         string `json:"title"`          //视频标题
}

type Author struct {
	Id            uint   `json:"id"`             //用户id
	Name          string `json:"name"`           //用户名称
	FollowCount   uint   `json:"follow_count"`   //关注总数
	FollowerCount uint   `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      //是否关注
}

type Fav struct {
	UserId uint   `json:"user_id"` //用户id
	Token  string `json:"token"`   //鉴权
}

type Action struct {
	VideoId    uint `json:"video_id"`
	ActionType int  `json:"action_type"`
}

type LikeVideos struct {
	UserId     uint  `json:"user_id"`
	VideoId    uint  `json:"video_id"`
	IsFavorite int   `json:"is_favorite"`
	LikeTime   int64 `json:"like_time"`
}

type ActionResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func (LikeVideos) TableName() string {
	return "like_videos"
}
