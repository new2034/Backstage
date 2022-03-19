/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-25 18:11:55
 * @LastEditTime: 2022-03-17 21:52:16
 */
package service

import (
	"Backstage/model"
	"Backstage/response"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 获取问题列表
 */
func (qs *QuestionService) Allq(c *gin.Context) {
	var ql model.QuestionList
	if !ql.Get() {
		response.Fail(response.SQLErr, c)
		return
	}
	if len(ql) == 0 {
		response.FailWithMessage(response.Success, "暂无问题", c)
		return
	}
	response.OkWithData(ql.QQDto(), c)
}

/**
 * @description: 问题搜索
 */
func (qs *QuestionService) Search(search *model.QuestionDetail) (data interface{}, err error) {
	ql := &model.QuestionList{}
	if len(search.Username) > 0 {
		u := &model.User{}
		if err = u.GetByUsername(search.Username); err != nil {
			return
		}
		if err = model.GetByUid(u.ID, ql); err != nil {
			return
		}
	} else {
		if err = ql.Searchq(search.Question); err != nil {
			return
		}
	}

	return ql.QQDto(), nil
}

/**
 * @description:  id删除问题
 */
func (qs *QuestionService) Removeq(qid string, c *gin.Context) {
	var qt model.Question
	var qa model.QueAns
	if !qt.GetByQID(qid) {
		response.FailWithMessage(response.Success, "问题不存在", c)
		return
	}
	if qt.Removeq(qid) && qa.Removeq(qid) {
		response.OkWithMessage("删除成功", c)
	} else {
		response.Fail(response.SQLErr, c)
	}
}

/**
 * @description: 获取问题详细页
 */
func (qs *QuestionService) Detailq(qid string, c *gin.Context) {
	var qt model.Question
	var qal model.QueAnsList
	if !qt.GetByQID(qid) {
		response.FailWithMessage(response.Success, "问题不存在", c)
		return
	}
	if !qal.GetByQID(qid) {
		response.Fail(response.SQLErr, c)
		return
	}
	if len(qal) == 0 {
		response.FailWithMessage(response.Success, "暂无答案", c)
		return
	}
	c.JSON(200, gin.H{
		"code":     response.Success,
		"Question": qt.QTop(),
		"Answers":  qal.ADto(),
		"msg":      "操作成功",
	})
}
