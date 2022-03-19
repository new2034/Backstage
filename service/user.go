package service

import (
	"Backstage/model"
	"Backstage/response"

	"github.com/gin-gonic/gin"
)

//所有用户返回
func (us *UserService) AllUser(c *gin.Context) {
	var ua []model.User
	var um []model.UserDetail2

	model.AllBack(&ua)

	if len(ua) == 0 {
		response.FailWithMessage(404, "暂无用户数据", c)
	} else {

		for i := 0; i < len(ua); i++ {
			um = append(um, ua[i].Detail2())
		}

		response.OkWithData(gin.H{"userarray": um}, c)
	}
}

//搜索用户
func (us *UserService) SearchUser(e string, c *gin.Context) {
	var u model.User
	var um model.UserDetail2
	u.SearchUser(e)
	if u.Email == "" {
		response.FailWithMessage(404, "未找到用户", c)
		return
	}

	um = u.Detail2()
	response.OkWithData(gin.H{"user": um}, c)

}

//删除用户
func (us *UserService) DeleteUser(k model.DeleteUser, c *gin.Context) {

	var u model.User
	u.SearchUserID(k.ID)
	if u.Username == "" {
		response.FailWithMessage(404, "未找到用户", c)
		return
	}
	if u.DeleteByID(k.ID) {
		response.OkWithMessage("删除成功", c)
	} else {
		response.FailWithMessage(404, "删除失败", c)
	}

}

//修改用户密码
func (us *UserService) UpdataUserPwd(l model.UpdataUserPwd, c *gin.Context) {
	var u model.User
	u.SearchUser(l.Email)
	Passwords := Hash(l.Password, c)
	u.Password = Passwords
	if u.Update() {
		response.OkWithMessage("修改密码成功", c)
	} else {
		response.OkWithMessage("修改密码失败", c)
	}

}
