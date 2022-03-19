package service

import (
	"Backstage/model"
	"Backstage/response"
	"Backstage/utils"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//所有管理员返回
func (us *UserService) AllAdmin(c *gin.Context) {

	var ta []model.Admin
	var tm []model.AdminDetail3

	model.AllBack2(&ta)
	if len(ta) == 0 {
		response.FailWithMessage(404, "暂无管理员数据", c)
	} else {

		for g := 0; g < len(ta); g++ {
			tm = append(tm, ta[g].Detail3())
		}

		response.OkWithData(gin.H{"adminarray": tm}, c)
	}
}

//搜索管理员
func (us *UserService) SearchAdmin(e string, c *gin.Context) {
	var t model.Admin
	var tm model.AdminDetail3
	t.SearchAdmin(e)
	tm = t.Detail3()
	response.OkWithData(gin.H{"admin": tm}, c)

}

//删除管理员
func (us *UserService) DeleteAdmin(email string, c *gin.Context) {
	var t model.Admin
	if t.DeleteByEmail(email) {
		response.OkWithMessage("删除成功", c)
	} else {
		response.FailWithMessage(404, "删除失败", c)
	}

}

//修改管理员密码
func (tb *UserService) UpdataAdminPwd(p model.UpdataAdminPwd, c *gin.Context) {
	var t model.Admin
	if t.SearchAdmin(p.Email) == false {
		response.FailWithMessage(404, "未找到管理员", c)
		return
	}
	Passwords := Hash(p.Password, c)
	t.Password = Passwords
	if t.Update() {
		response.OkWithMessage("修改成功", c)
	} else {
		response.OkWithMessage("修改失败", c)
	}
}

    //添加管理员
func AddAdmin(r model.AddAdmin, c *gin.Context) {
	// 判断管理员名是否存在
	if utils.IsAdminnameExist(r.Adminname) {
		response.FailWithMessage(response.ParamErr, "用户名已存在", c)
		return
	}
	// 判断邮箱是否存在
	if utils.IsEmailExist2(r.Email) {
		response.FailWithMessage(response.ParamErr, "邮箱已存在", c)
		return
	}

	// 密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print("Password hash error:", err)
		response.FailWithMessage(response.ServerErr, "密码加密失败, 请勿使用违禁字符", c)
		return
	}
	r.Password = string(hashPassword)

	u := model.Admin{
		Adminname: r.Adminname,
		Password:  r.Password,
		Vip:       true,
	}

	u.Create()

	if u.ID > 0 {
		response.OkWithMessage("添加成功", c)
	} else {
		response.FailWithMessage(response.ParamErr, "添加失败, 请稍后尝试", c)
	}

}
