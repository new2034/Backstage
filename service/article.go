/*
 * @Author: fzf404
 * @Date: 2022-02-13 17:17:52
 * @LastEditTime: 2022-02-20 20:43:05
 * @Description: description
 */
package service

import (
	"Backstage/global"
	"Backstage/model"
	"Backstage/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

/**
 * @description: 删除文章
 * @param {string} aid
 * @param {uint} uid
 * @param {*gin.Context} c
 */
func Remove(i string, c *gin.Context) {

	var a model.Article

	if err := global.MYSQL_DB.Table("articles").Where("id = ?", i).First(&a).Unscoped().Delete(&a).Error; err != nil {
		response.FailWithMessage(response.ParamErr, "删除失败", c)
		return
	}

	response.OkWithDetailed(gin.H{"id":a.ID,"title":a.Title},"删除成功", c)

}

/**
 * @description: 搜索文章
 * @param {string} title
 * @param {int} page
 * @param {*gin.Context} c
 */
 func Searchtitle(t string, c *gin.Context) {

	var al model.ArticleList
	
	global.MYSQL_DB.Table("articles").Where("title = ?", t).Preload(clause.Associations).Find(&al)
	if len(al) != 0 {
		response.OkWithData(gin.H{"article":al.Dto()}, c)
		return
	}
	response.FailWithMessage(response.ParamErr, "搜索失败", c)
	
	}


func Searchwriter(w string, c *gin.Context) {

	var al model.ArticleList
	var u model.User
	global.MYSQL_DB.Table("users").Where("username = ?", w).Find(&u)
	global.MYSQL_DB.Table("articles").Where("user_id = ?", u.ID).Preload(clause.Associations).Find(&al)
	if len(al) != 0 {
		response.OkWithData(gin.H{"article":al.Dto()}, c)
		return
	}
	response.FailWithMessage(response.ParamErr, "搜索失败", c)
	
}
/**
 * @description: 文章详细信息
 * @param {int} id
 * @param {*gin.Context} c
 */
func Articledata(i string, c *gin.Context) {
	var a model.Article

	global.MYSQL_DB.Table("articles").Where("id = ?", i).Preload(clause.Associations).First(&a)
    b:=model.TimeDto(a.CreatedAt)
	response.OkWithData(gin.H{"id":a.ID,"title":a.Title,"writer":a.User.Username,"createtime":b,"like":a.Like,"pageviews":a.View,"content":a.Content}, c)
}

/**
 * @description: 文章列表
 * @param {int} page
 * @param {*gin.Context} c
 */
func Article(c *gin.Context) {

	// 文章列表
	var al model.ArticleList
	global.MYSQL_DB.Table("articles").Preload(clause.Associations).Find(&al)
	response.OkWithData(gin.H{"article":al.Dto()}, c)
}

func Classall(i string, c *gin.Context) {

	var al model.ArticleList
	var a model.Article
	if i == "1" {
		global.MYSQL_DB.Table("articles").Preload(clause.Associations).Find(&al)
		response.OkWithData(gin.H{"article":al.Dto()}, c)
		return
	}
	if i== "2" {
		global.MYSQL_DB.Table("articles").Where("category = ?", "游戏").Preload(clause.Associations).Find(&a).Scan(&al)
		response.OkWithData(gin.H{"article":al.Dto()}, c)
		return
	}
	if i == "3" {
		global.MYSQL_DB.Table("articles").Where("category = ?", "软件").Preload(clause.Associations).Find(&a).Scan(&al)
		response.OkWithData(gin.H{"article":al.Dto()}, c)
		return
	}
	if i == "4" {
		global.MYSQL_DB.Table("articles").Where("category = ?", "影视").Preload(clause.Associations).Find(&a).Scan(&al)
		response.OkWithData(gin.H{"article":al.Dto()}, c)
		return
	}
	if i== "5"{
		global.MYSQL_DB.Table("articles").Where("category = ?", "网页").Preload(clause.Associations).Find(&a).Scan(&al)
		response.OkWithData(gin.H{"article":al.Dto()}, c)
		return
	}

}
