//登录接口

package api

import (
	"Backstage/model"
	"Backstage/response"
	"Backstage/service"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 用户登录
 * @param {*gin.Context} c
 */
func Login(c *gin.Context) {

	var l model.Login

	if err := c.ShouldBind(&l); err != nil {
		response.FailWithDetailed(response.ParamErr, err.Error(), "邮箱或密码错误", c)
		return
	}

	service.Login(l, c)
}
