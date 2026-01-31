package model

import "time"

// UserToken 用户令牌表
type UserToken struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint64    `json:"userId" gorm:"column:user_id;index"`
	Token     string    `json:"token" gorm:"column:token;type:varchar(512);uniqueIndex"`
	ExpireAt  time.Time `json:"expireAt" gorm:"column:expire_at"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
}

func (UserToken) TableName() string {
	return "user_tokens"
}
