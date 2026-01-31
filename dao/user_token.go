package dao

import (
	"time"

	"github.com/everyday-items/gin-example/library/db"
	"github.com/everyday-items/gin-example/model"
)

// UserTokenDao 用户令牌数据访问层
type UserTokenDao struct{}

// NewUserTokenDao 创建实例
func NewUserTokenDao() *UserTokenDao {
	return &UserTokenDao{}
}

// GetByToken 根据token获取令牌记录
func (d *UserTokenDao) GetByToken(token string) (*model.UserToken, error) {
	var userToken model.UserToken
	err := db.GetDB().Where("token = ?", token).First(&userToken).Error
	if err != nil {
		return nil, err
	}
	return &userToken, nil
}

// GetValidByToken 获取有效的token（未过期）
func (d *UserTokenDao) GetValidByToken(token string) (*model.UserToken, error) {
	var userToken model.UserToken
	err := db.GetDB().Where("token = ? AND expire_at > ?", token, time.Now()).
		First(&userToken).Error
	if err != nil {
		return nil, err
	}
	return &userToken, nil
}

// Create 创建令牌
func (d *UserTokenDao) Create(userToken *model.UserToken) error {
	return db.GetDB().Create(userToken).Error
}

// DeleteByToken 删除令牌
func (d *UserTokenDao) DeleteByToken(token string) error {
	return db.GetDB().Where("token = ?", token).Delete(&model.UserToken{}).Error
}

// DeleteByUserID 删除用户所有令牌
func (d *UserTokenDao) DeleteByUserID(userID uint64) error {
	return db.GetDB().Where("user_id = ?", userID).Delete(&model.UserToken{}).Error
}

// DeleteExpired 清理过期令牌
func (d *UserTokenDao) DeleteExpired() error {
	return db.GetDB().Where("expire_at < ?", time.Now()).Delete(&model.UserToken{}).Error
}
