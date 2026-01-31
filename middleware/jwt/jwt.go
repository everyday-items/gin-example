package jwt

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/everyday-items/gin-example/library/e"
	"github.com/everyday-items/gin-example/library/setting"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT claims
type Claims struct {
	UserID uint64 `json:"userId"`
	Openid string `json:"openid"`
	jwt.RegisteredClaims
}

// Auth 认证中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getTokenFromHeader(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": e.UNAUTHORIZED,
				"msg":  e.GetMsg(e.UNAUTHORIZED),
				"data": struct{}{},
			})
			c.Abort()
			return
		}

		// 解析JWT
		claims, err := ParseJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": e.UNAUTHORIZED,
				"msg":  e.GetMsg(e.UNAUTHORIZED),
				"data": struct{}{},
			})
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set("userID", claims.UserID)
		c.Set("openid", claims.Openid)
		c.Set("token", token)
		c.Next()
	}
}

// ParseJWT 解析JWT
func ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
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

	return nil, fmt.Errorf("invalid token")
}

// getTokenFromHeader 从请求头获取token
func getTokenFromHeader(c *gin.Context) string {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		return ""
	}
	// 支持 "Bearer xxx" 和 "xxx" 两种格式
	if token, found := strings.CutPrefix(auth, "Bearer "); found {
		return token
	}
	return auth
}
