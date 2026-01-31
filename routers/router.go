package routers

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/everyday-items/gin-example/app/api/v1"
	"github.com/everyday-items/gin-example/middleware/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查接口（无需认证）
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/check", v1.Check)
		apiv1.POST("/check", v1.Check)
	}

	// 认证接口（无需认证）
	authApi := r.Group("/api/auth")
	{
		authApi.POST("/login", v1.Login)
		authApi.POST("/check", v1.AuthCheck)
		authApi.POST("/logout", v1.Logout)
	}

	// 需要认证的接口
	userApi := r.Group("/api/user")
	userApi.Use(jwt.Auth())
	{
		// 用户信息
		userApi.GET("/info", v1.GetUserInfo)
	}

	return r
}
