package model

type CommitAction struct {
	Id          uint   `json:"id"  gorm:"<-;primaryKey" ` // 用户id
	Content     string `json:"content"`                   // 用户名称
	Create_date string `json:"create_date"`
	Video_id    string `json:"video_id"`
	Uid         uint   `json:"uid"`
}
type CommitActionPush struct {
	Id          uint   `json:"id"  `    // 用户id
	Content     string `json:"content"` // 用户名称
	Create_date string `json:"create_date"`
	User        *User  `json:"user" `
}
type CommitActionListPush struct {
	StatusCode   int64              `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg    string             `json:"status_msg"`   // 返回状态描述
	Comment_list []CommitActionPush `json:"comment_list"` // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}
type CommitResp struct {
	StatusCode int64             `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string            `json:"status_msg"`  // 返回状态描述
	Comment    *CommitActionPush `json:"comment"`     // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}
type AB struct {
	StatusCode int64 `json:"status_code"` // 状态码，0-成功，其他值-失败
}
