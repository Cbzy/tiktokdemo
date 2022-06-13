package model

// 视频作者信息
type User struct {
	FollowCount   int64  `json:"follow_count"`              // 关注总数
	FollowerCount int64  `json:"follower_count"`            // 粉丝总数
	Id            uint   `json:"id"  gorm:"<-;primaryKey" ` // 用户id
	Name          string `json:"name"`                      // 用户名称
	Username      string `json:"username"`
	Password      string `json:"password"`
}

type UserFans struct {
	FollowCount   int64  `json:"follow_count"`              // 关注总数
	FollowerCount int64  `json:"follower_count"`            // 粉丝总数
	Id            uint   `json:"id"  gorm:"<-;primaryKey" ` // 用户id
	Name          string `json:"name"`
	Is_follow bool `json:"is_follow"`
}

type UserResp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	User_id    uint   `json:"user_id"`     // 视频列表
	Token      string `json:"token"`       // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}
type UserinfoResp struct {
	StatusCode int64    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `json:"status_msg"`  // 返回状态描述
	User       UserFans `json:"user"`
}
