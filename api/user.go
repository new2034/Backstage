package api

import (
	"Backstage/model"
	"Backstage/response"

	"github.com/gin-gonic/gin"
)

//显示所有用户
func AllUser(c *gin.Context) {
	userService.AllUser(c)
}

//邮件搜索用户
func SearchUser(c *gin.Context) {
	e, ok := c.GetQuery("email")

	if !ok {
		response.Fail(response.ParamErr, c)
		return
	}

	userService.SearchUser(e, c)
}

//删除用户
func DeleteUser(c *gin.Context) {

	vip := c.MustGet("Vip").(bool)
	apassword := c.MustGet("Admin_password").(string)

	var k model.DeleteUser

	if err := c.ShouldBindJSON(&k); err != nil {
		response.Fail(response.ParamErr, c)
		return
	}

	if vip == true {
		userService.DeleteUser(k, c)
	} else if !vip || apassword != k.Password {
		response.FailWithMessage(403, "您不是管理员无权操作,或管理员密码错误", c)
	}

}

//修改用户密码
func UpdataUserPwd(c *gin.Context) {
	var l model.UpdataUserPwd
	vip := c.MustGet("Vip").(bool)

	if err := c.ShouldBindJSON(&l); err != nil {
		response.Fail(response.ParamErr, c)
		return
	}

	if vip == true {
		userService.UpdataUserPwd(l, c)
	} else {
		response.FailWithMessage(403, "您不是管理员无权操作", c)
	}

}
