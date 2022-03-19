package router

import (
	"Backstage/api"
	"Backstage/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	route := r.Group("/admin")
	{
		route.POST("/login", api.Login)
		route.GET("/allUser", api.AllUser)       //显示所有用户
		route.GET("/searchUser", api.SearchUser) //搜索用户

		route.GET("/allAdmin", api.AllAdmin)       //显示所有管理员
		route.GET("/searchAdmin", api.SearchAdmin) //搜索管理员

		//route.Use(middleware.Auth())
		{
			route.POST("/getdata", api.Getdata)
			route.POST("/article", api.Article)
			route.POST("/searchtitle", api.Searchtitle)
			route.POST("/searchwriter", api.Searchwriter)
			route.POST("/articledata", api.Articledata)
			route.POST("/remove", api.Remove)

			route.POST("/deleteAdmin", api.DeleteAdmin)       //删除管理员
			route.POST("/addAdmin", api.AddAdmin)             //添加管理员
			route.POST("/updataAdminPwd", api.UpdataAdminPwd) //修改管理员密码

			route.GET("/question/all", api.Questionall)

			route.GET("/resource/all", api.Resourceall)
			route.GET("/resource/search", api.ResourceSearch)
			route.GET("/resource/content", api.Resourcecontent)
			route.POST("/resource/remove", api.Resourceremove)
			route.POST("/classall", api.Classall)
			route.GET("/question/search", api.Questionsearch)
			route.POST("/question/remove", api.Questionremove)
			route.GET("/question/details", api.Questiondetails)
			route.GET("/question/details/answer", api.Questiondetailsanswer)
			route.POST("/question/details/remove", api.Questiondetailsremove)
			route.GET("/question/details/search", api.QuestionDetailsSearch)
			route.POST("/updataUserPwd", api.UpdataUserPwd) //修改用户密码
			route.POST("/deleteUser", api.DeleteUser)       //删除用户
		}

	}
	return r
}
