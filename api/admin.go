package api

import (
	"Backstage/model"
	"Backstage/response"
	"Backstage/service"

	"github.com/gin-gonic/gin"
)

//显示所有管理员
func AllAdmin(c *gin.Context) {
	userService.AllAdmin(c)
}

//邮件搜索管理员
func SearchAdmin(c *gin.Context) {
	e, ok := c.GetQuery("email")

	if !ok {
		response.Fail(response.ParamErr, c)
		return
	}

	userService.SearchAdmin(e, c)
}

//删除管理员
func DeleteAdmin(c *gin.Context) {
	e, _ := c.GetQuery("email")
	super := c.MustGet("IsSuper").(bool)

	if super == false {
		response.FailWithMessage(403, "你不是超级管理员", c)
		return
	} else {
		userService.DeleteAdmin(e, c)
	}

}

//修改管理员密码
func UpdataAdminPwd(c *gin.Context) {
	var p model.UpdataAdminPwd
	if err := c.ShouldBindJSON(&p); err != nil {
		response.Fail(response.ParamErr, c)
		return
	}
	super := c.MustGet("IsSuper").(bool)
	if super == false {
		response.FailWithMessage(403, "你不是超级管理员", c)
		return
	} else {
		userService.UpdataAdminPwd(p, c)
	}

}
func AddAdmin(c *gin.Context) {
	// 绑定数据
	var r model.AddAdmin
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithDetailed(response.ParamErr, err.Error(), "提交信息非法！", c)
		return
	}

	service.AddAdmin(r, c)
}
