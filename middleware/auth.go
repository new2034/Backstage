/*
 * @Author: fzf404
 * @Date: 2022-01-22 15:35:53
 * @LastEditTime: 2022-03-14 15:48:46
 * @Description: 鉴权
 */

package middleware

import (
	"Backstage/response"
	"Backstage/utils"

	"github.com/gin-gonic/gin"
)

/**
 * @description: token 鉴权
 */
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 Token
		tokenString := c.GetHeader("x-token")
		if tokenString == "" {
			response.FailWithMessage(response.NoLogin, "未登录", c)
			c.Abort()
			return
		}

		// 解析 Token
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.FailWithMessage(response.TokenExpired, "登录已过期", c)
			c.Abort()
			return
		}

		// 写入上下文
		c.Set("UID", claims.UID)
		c.Next()
	}
}
