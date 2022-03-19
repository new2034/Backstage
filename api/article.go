//文章相关接口

package api

import (
	"Backstage/global"
	"Backstage/model"
	"Backstage/response"
	"Backstage/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* /**
 * @description: 删除文章
 * @param {*gin.Context} c
 */
func Remove(c *gin.Context) {
	i := c.PostForm("id")

	service.Remove(i, c)
}

/**
 * @description: 搜索文章
 * @param {*gin.Context} c
 */
 func Searchtitle(c *gin.Context) {

	t := c.PostForm("title")

	service.Searchtitle(t, c)
}


func Searchwriter(c *gin.Context) {

	w := c.PostForm("writer")

	service.Searchwriter(w, c)
}
/**
 * @description: 获取文章详细信息
 * @param {*gin.Context} c
 */
func Articledata(c *gin.Context) {

	i := c.PostForm("id")
	service.Articledata(i, c)
}

/**
 * @description: 获取文章总数
 * @param {*gin.Context} c
 */
func Getdata(c *gin.Context) {
	var u model.User
	var ui int64
	global.MYSQL_DB.Model(&u).Count(&ui)

	var a model.Article
	var ai int64
	global.MYSQL_DB.Model(&a).Count(&ai)

	var r model.Resource
	var ri int64
	global.MYSQL_DB.Model(&r).Count(&ri)

	var q model.Question
	var qi int64
	global.MYSQL_DB.Model(&q).Count(&qi)
	 
	var j int64
	Time("2022-01-01 00:00:00","2022-01-31 23:59:59").Count(&j)

 	var f int64
	Time("2022-02-01 00:00:00","2022-02-31 23:59:59").Count(&f)

	var m int64
	Time("2022-03-01 00:00:00","2022-03-31 23:59:59").Count(&m)

	var ap int64
	Time("2022-04-01 00:00:00","2022-04-31 23:59:59").Count(&ap)

	var ma int64
	Time("2022-05-01 00:00:00","2022-05-31 23:59:59").Count(&ma)

	var ju int64
	Time("2022-06-01 00:00:00","2022-06-31 23:59:59").Count(&ju)

	var jui int64
	Time("2022-07-01 00:00:00","2022-07-31 23:59:59").Count(&jui)

	var au int64
	Time("2022-08-01 00:00:00","2022-08-31 23:59:59").Count(&au)

	var s int64
	Time("2022-09-01 00:00:00","2022-09-31 23:59:59").Count(&s)

	var o int64
	Time("2022-10-01 00:00:00","2022-10-31 23:59:59").Count(&o)

	var n int64
	Time("2022-11-01 00:00:00","2022-11-31 23:59:59").Count(&n)

	var d int64
	Time("2022-12-01 00:00:00","2022-12-31 23:59:59").Count(&d)

	ddd := []model.Data{
	model.Data{
	"1月",
	j,
	},
	 model.Data{
	"2月",
	f,
	 },
	model.Data{
	"3月",
 	m,
	},
	model.Data{
 	"4月",
 	ap,
	},
	model.Data{
	"5月",
	ma,
	},
	model.Data{
 	"6月",
 	ju,
	},
	model.Data{
	"7月",
	jui,
	},
	model.Data{
	"8月",
 	au,
	},
	model.Data{
 	"9月",
 	s,
	},
	model.Data{
	"10月",
	o,
	 },
	model.Data{
	"11月",
	 n,
 	},
	model.Data{
	"12月",
	 d,
	},
}
	response.OkWithDetailed(gin.H{
		"Data":gin.H{
			"userdata": ui,"articledata":ai,"resourcedata":ri,"questiondata":qi},
		"rows":ddd}, "返回数据成功", c)
}

func Time(timea string,timeb string) *gorm.DB {
	var u model.User
	return global.MYSQL_DB.Model(&u).Where("created_at > ?", timea).Where("created_at < ?",timeb).Find(&u)
  }
/**
 * @description: 获取文章列表
 * @param {*gin.Context} c
 */
func Article(c *gin.Context) {

	service.Article(c)
}

func Classall(c *gin.Context) {
	i := c.PostForm("id")
	service.Classall(i, c)
}