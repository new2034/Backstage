package service

import (
	"Backstage/response"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//密码单独加密接口
func Hash(password string, c *gin.Context) string {

	// 密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Print("Password hash error:", err)
		response.FailWithMessage(response.ServerErr, "密码加密错误", c)
	}
	Password := string(hashPassword)
	return Password

}
