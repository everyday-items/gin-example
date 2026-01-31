package model

import "time"

// User 用户表
type User struct {
	ID          uint64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Openid      string     `json:"openid" gorm:"column:openid;type:varchar(64);uniqueIndex"`
	UnionID     *string    `json:"unionId" gorm:"column:union_id;type:varchar(64);index"`
	SessionKey  *string    `json:"-" gorm:"column:session_key;type:varchar(128)"`
	NickName    string     `json:"nickName" gorm:"column:nick_name;type:varchar(64);default:'微信用户'"`
	AvatarURL   *string    `json:"avatarUrl" gorm:"column:avatar_url;type:varchar(512)"`
	Gender      int8       `json:"gender" gorm:"column:gender;default:0"`
	Phone       *string    `json:"phone" gorm:"column:phone;type:varchar(20)"`
	Status      int8       `json:"status" gorm:"column:status;default:1"`
	LastLoginAt *time.Time `json:"lastLoginAt" gorm:"column:last_login_at"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
