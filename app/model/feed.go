package model

//视频流相应
type Feed struct {
	NextTime   *int64      `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64       `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string     `json:"status_msg"`  // 返回状态描述
	VideoList  []VideoResp `json:"video_list"`  // 视频列表
}

//发布接口相应
type PublishListResp struct {
	StatusCode int64       `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string     `json:"status_msg"`  // 返回状态描述
	VideoList  []VideoResp `json:"video_list"`  // 用户发布的视频列表
}

//投稿接口相应
type PublishActionResp struct {
	StatusCode int64   `json:"status_code"`
	StatusMsg  *string `json:"status_msg"`
}

// Video
type Video struct {
	Author        User   `json:"author"`                            // 视频作者信息
	CommentCount  int64  `json:"comment_count"`                     // 视频的评论总数
	CoverURL      string `json:"cover_url" gorm:"column:cover_url"` // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"`                    // 视频的点赞总数
	ID            int64  `json:"id" gorm:"column:v_id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`                       // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url" gorm:"column:play_url"`   // 视频播放地址
	Title         string `json:"title"`                             // 视频标题
	PublicTime    int64  `json:"public_time"`                       //视频发布时间
}

// Video
type VideoResp struct {
	Author        UserFans `json:"author"`                            // 视频作者信息
	CommentCount  int64    `json:"comment_count"`                     // 视频的评论总数
	CoverURL      string   `json:"cover_url" gorm:"column:cover_url"` // 视频封面地址
	FavoriteCount int64    `json:"favorite_count"`                    // 视频的点赞总数
	ID            int64    `json:"id" gorm:"column:v_id"`             // 视频唯一标识
	IsFavorite    bool     `json:"is_favorite"`                       // true-已点赞，false-未点赞
	PlayURL       string   `json:"play_url" gorm:"column:play_url"`   // 视频播放地址
	Title         string   `json:"title"`                             // 视频标题
	PublicTime    int64    `json:"public_time"`                       //视频发布时间
}

type VideoADD struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"userId"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
	Title         string `json:"title"`
	PublicTime    string `json:"public_time"`
}

func (VideoADD) TableName() string {
	return "videos"
}
