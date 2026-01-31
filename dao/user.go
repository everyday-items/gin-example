package dao

import (
	"github.com/everyday-items/gin-example/library/db"
	"github.com/everyday-items/gin-example/model"
)

// UserDao 用户数据访问层
type UserDao struct{}

// NewUserDao 创建实例
func NewUserDao() *UserDao {
	return &UserDao{}
}

// GetByID 根据ID获取用户
func (d *UserDao) GetByID(id uint64) (*model.User, error) {
	var user model.User
	err := db.GetDB().Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByOpenid 根据openid获取用户
func (d *UserDao) GetByOpenid(openid string) (*model.User, error) {
	var user model.User
	err := db.GetDB().Where("openid = ?", openid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create 创建用户
func (d *UserDao) Create(user *model.User) error {
	return db.GetDB().Create(user).Error
}

// Update 更新用户
func (d *UserDao) Update(user *model.User) error {
	return db.GetDB().Save(user).Error
}

// UpdateLastLogin 更新最后登录时间
func (d *UserDao) UpdateLastLogin(userID uint64) error {
	return db.GetDB().Model(&model.User{}).Where("id = ?", userID).
		Update("last_login_at", db.GetDB().NowFunc()).Error
}
