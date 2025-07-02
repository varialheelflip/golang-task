package middleWare

import (
	"blog_system/pkg/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 放行登录和注册接口
		if c.Request.URL.Path == "/api/v1/users/login" || c.Request.URL.Path == "/api/v1/users/register" {
			c.Next()
			return
		}
		authHeader := c.GetHeader("token")
		if authHeader == "" {
			c.Abort()
			response.BadRequest(c, "请先登录!")
			return
		}
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			// todo 配置化
			return []byte("your_secret_key"), nil
		})

		if err != nil || !token.Valid {
			c.Abort()
			response.BadRequest(c, "无效的token")
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("userID", claims["id"])
		}
		c.Next()
	}
}
