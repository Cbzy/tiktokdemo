package testm

type Tags struct {
	//ID   int    `gorm:"primarykey"` // 主键ID
	Uname string `json:"uname" form:"uname" gorm:"comment:标签名"`
}
