package v1

import (
	"net/http"
	"strings"

	"github.com/everyday-items/gin-example/library/app"
	"github.com/everyday-items/gin-example/library/e"
	"github.com/everyday-items/gin-example/model"
	"github.com/everyday-items/gin-example/service"
	"github.com/gin-gonic/gin"
)

// AuthApi 认证控制器
type AuthApi struct {
	svc *service.AuthService
}

// NewAuthApi 创建实例
func NewAuthApi() *AuthApi {
	return &AuthApi{
		svc: service.NewAuthService(),
	}
}

var authApi = NewAuthApi()

// Login 小程序登录接口
// @Summary 小程序登录接口
// @Description 使用微信code登录，返回token
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body model.LoginRequest true "登录请求"
// @Success 200 {object} model.LoginResponse
// @Router /api/mini/auth/login [post]
func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	data, err := authApi.svc.Login(req.Code)
	if err != nil {
		appG.ResponseWithMsg(http.StatusOK, e.ERROR_AUTH_LOGIN_FAIL, err.Error(), nil)
		return
	}

	appG.ResponseWithMsg(http.StatusOK, e.SUCCESS, "登录成功", data)
}

// AuthCheck 登录有效性校验接口
// @Summary 登录有效性校验接口
// @Description 校验当前登录态是否有效
// @Tags 认证
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} model.AuthCheckResponse
// @Router /api/mini/auth/check [post]
func AuthCheck(c *gin.Context) {
	appG := app.Gin{C: c}

	token := getTokenFromHeader(c)
	if token == "" {
		appG.ResponseWithMsg(http.StatusOK, e.UNAUTHORIZED, "登录态已失效，请重新登录", nil)
		return
	}

	data, err := authApi.svc.Check(token)
	if err != nil {
		appG.ResponseWithMsg(http.StatusOK, e.UNAUTHORIZED, err.Error(), nil)
		return
	}

	appG.ResponseWithMsg(http.StatusOK, e.SUCCESS, "登录态有效", data)
}

// Logout 退出登录接口
// @Summary 退出登录接口
// @Description 退出登录，使token失效
// @Tags 认证
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} map[string]interface{}
// @Router /api/mini/auth/logout [post]
func Logout(c *gin.Context) {
	appG := app.Gin{C: c}

	token := getTokenFromHeader(c)
	if token == "" {
		appG.ResponseWithMsg(http.StatusOK, e.SUCCESS, "退出登录成功", nil)
		return
	}

	_ = authApi.svc.Logout(token)
	appG.ResponseWithMsg(http.StatusOK, e.SUCCESS, "退出登录成功", nil)
}

// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Description 获取当前登录用户的基本信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} model.User
// @Router /api/user/info [get]
func GetUserInfo(c *gin.Context) {
	appG := app.Gin{C: c}

	userID := c.MustGet("userID").(uint64)

	user, err := authApi.svc.GetUserByID(userID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, user)
}

// getTokenFromHeader 从请求头获取token
func getTokenFromHeader(c *gin.Context) string {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		return ""
	}
	// 支持 "Bearer xxx" 和 "xxx" 两种格式
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return auth
}
