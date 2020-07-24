package middleware

import (
	"common/jwt"
	R "common/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) < 7 {
			c.Abort()
			R.Response(c, http.StatusUnauthorized, "未登录", nil, http.StatusUnauthorized)
			return
		}
		jwtUserInfo := jwt.UserInfo{}
		err := jwtUserInfo.ParseToken(token[7:])
		if err != nil {
			R.Response(c, http.StatusUnauthorized, "未登录", nil, http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Set("jwtUserInfo", jwtUserInfo)
		c.Next()
		return
	}
}
