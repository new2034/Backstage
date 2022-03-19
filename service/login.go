package service

import (
	"Backstage/global"
	"Backstage/model"
	"Backstage/response"
	"Backstage/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func Login(l model.Login, c *gin.Context) {

	var a model.Admin
	/* 	if err := global.MYSQL_DB.Model(&a).Where("adminname = ?", l.Email).First(&a).Error; err != nil {
	   		response.FailWithMessage(response.ParamErr, "邮箱错误", c)
	   		return
	   	} else if err := global.MYSQL_DB.Model(&a).Where("password = ?", l.Password).Find(&a).Error; err != nil {
	   		response.FailWithMessage(response.ParamErr, "密码错误", c)
	   		return
	   	} */

	if err := global.MYSQL_DB.Where("adminname = ?", l.Email).First(&a).Error; err != nil {

		response.FailWithMessage(response.ParamErr, "邮箱错误", c)
		return
	}

	// 密码验证
	if err := global.MYSQL_DB.Where("password = ?", l.Password).First(&a).Error; err != nil {

		response.FailWithMessage(response.ParamErr, "密码错误", c)
		return
	}
	token, err := utils.ReleaseToken(a)
	if err != nil {
		log.Print("Token generate error:", err)
		response.FailWithMessage(response.TokenErr, "Token分发错误", c)
		return
	}

	response.OkWithDetailed(gin.H{"token": token, "admin": a.Vip}, "登录成功", c)
}
