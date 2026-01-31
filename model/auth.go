package model

// LoginRequest 登录请求
type LoginRequest struct {
	Code string `json:"code" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token      string `json:"token"`
	ExpireTime int64  `json:"expireTime"`
	IsNewUser  bool   `json:"isNewUser"`
	UserID     uint64 `json:"userId"`
}

// AuthCheckResponse 登录态检查响应
type AuthCheckResponse struct {
	UserID    uint64  `json:"userId"`
	Openid    string  `json:"openid"`
	NickName  string  `json:"nickName"`
	AvatarURL *string `json:"avatarUrl"`
}

// WxLoginResponse 微信登录接口响应
type WxLoginResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
