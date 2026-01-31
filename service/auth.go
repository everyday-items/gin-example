package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/everyday-items/gin-example/dao"
	"github.com/everyday-items/gin-example/library/setting"
	"github.com/everyday-items/gin-example/model"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// Claims JWT claims
type Claims struct {
	UserID uint64 `json:"userId"`
	Openid string `json:"openid"`
	jwt.RegisteredClaims
}

// AuthService 认证服务层
type AuthService struct {
	userDao *dao.UserDao
}

// NewAuthService 创建实例
func NewAuthService() *AuthService {
	return &AuthService{
		userDao: dao.NewUserDao(),
	}
}

// Login 小程序登录
func (s *AuthService) Login(code string) (*model.LoginResponse, error) {
	// 调用微信接口获取openid
	wxResp, err := s.wxLogin(code)
	if err != nil {
		return nil, err
	}

	if wxResp.ErrCode != 0 {
		return nil, fmt.Errorf("微信登录失败: %s", wxResp.ErrMsg)
	}

	// 查找或创建用户
	isNewUser := false
	user, err := s.userDao.GetByOpenid(wxResp.Openid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 新用户，创建账号
			user = &model.User{
				Openid:     wxResp.Openid,
				SessionKey: &wxResp.SessionKey,
			}
			if wxResp.UnionID != "" {
				user.UnionID = &wxResp.UnionID
			}
			if err := s.userDao.Create(user); err != nil {
				return nil, err
			}
			isNewUser = true
		} else {
			return nil, err
		}
	} else {
		// 老用户，更新session_key
		user.SessionKey = &wxResp.SessionKey
		if err := s.userDao.Update(user); err != nil {
			return nil, err
		}
	}

	// 更新最后登录时间
	_ = s.userDao.UpdateLastLogin(user.ID)

	// 生成JWT
	token, expireAt, err := s.generateJWT(user.ID, user.Openid)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token:      token,
		ExpireTime: expireAt.UnixMilli(),
		IsNewUser:  isNewUser,
		UserID:     user.ID,
	}, nil
}

// Check 检查登录态
func (s *AuthService) Check(token string) (*model.AuthCheckResponse, error) {
	// 解析JWT
	claims, err := s.ParseJWT(token)
	if err != nil {
		return nil, errors.New("登录态已失效，请重新登录")
	}

	// 获取用户信息
	user, err := s.userDao.GetByID(claims.UserID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	return &model.AuthCheckResponse{
		UserID:    user.ID,
		Openid:    user.Openid,
		NickName:  user.NickName,
		AvatarURL: user.AvatarURL,
	}, nil
}

// Logout 退出登录
// JWT是无状态的，客户端删除token即可
// 如需服务端主动失效，可实现token黑名单
func (s *AuthService) Logout(token string) error {
	// JWT无状态，无需服务端处理
	// 可选：将token加入黑名单（需要额外存储）
	return nil
}

// GetUserIDByToken 根据token获取用户ID
func (s *AuthService) GetUserIDByToken(token string) (uint64, error) {
	claims, err := s.ParseJWT(token)
	if err != nil {
		return 0, errors.New("登录态已失效")
	}
	return claims.UserID, nil
}

// GetUserByID 根据用户ID获取用户信息
func (s *AuthService) GetUserByID(userID uint64) (*model.User, error) {
	return s.userDao.GetByID(userID)
}

// ParseJWT 解析JWT
func (s *AuthService) ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(setting.WechatSetting.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// generateJWT 生成JWT
func (s *AuthService) generateJWT(userID uint64, openid string) (string, time.Time, error) {
	expireHours := setting.WechatSetting.TokenExpire
	if expireHours <= 0 {
		expireHours = 168 // 默认7天
	}
	expireAt := time.Now().Add(time.Duration(expireHours) * time.Hour)

	claims := &Claims{
		UserID: userID,
		Openid: openid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-example",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(setting.WechatSetting.JwtSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expireAt, nil
}

// wxLogin 调用微信登录接口
func (s *AuthService) wxLogin(code string) (*model.WxLoginResponse, error) {
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		setting.WechatSetting.AppID,
		setting.WechatSetting.AppSecret,
		code,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求微信接口失败: %v", err)
	}
	defer resp.Body.Close()

	var wxResp model.WxLoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&wxResp); err != nil {
		return nil, fmt.Errorf("解析微信响应失败: %v", err)
	}

	return &wxResp, nil
}
