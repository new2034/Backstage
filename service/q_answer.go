/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-26 10:24:08
 * @LastEditTime: 2022-03-17 21:51:01
 */
package service

import (
	"Backstage/model"
	"Backstage/response"

	"github.com/gin-gonic/gin"
)

/**
 * @description:搜索答案
 */
func (as *AnswerService) Search(search *model.QuestionAnswer) (data interface{}, err error) {
	qa := &model.QueAns{}
	al := &model.AnswerList{}
	qal := &model.QueAnsList{}
	if err = qal.SearchByQID(search.QuestionId); err != nil {
		return
	}
	if len(search.Replier) > 0 {
		u := &model.User{}
		if err = u.GetByUsername(search.Replier); err != nil {
			return
		}
		if err = model.SearchByUid(qa.AnswerID, u.ID, al); err != nil {
			return
		}
	} else {
		if err = al.Searcha(qa.AnswerID, search.Answer); err != nil {
			return
		}
	}
	return qal.ADto(), nil
}

/**
 * @description: 通过答案id查看答案详细
 */
func (as *AnswerService) Detaila(aid string, c *gin.Context) {
	var a model.Answer
	if !a.GetByAID(aid) {
		response.FailWithMessage(response.Success, "暂无答案", c)
		return
	}
	response.OkWithData(a.ADto(), c)
}

/**
 * @description: 通过答案id删除答案
 */
func (as *AnswerService) Removea(aid string, c *gin.Context) {
	var a model.Answer
	var qa model.QueAns
	if !a.GetByAID(aid) {
		response.FailWithMessage(response.Success, "暂无答案", c)
		return
	}
	if a.Removean(aid) && qa.Removea(aid) {
		response.OkWithMessage("删除成功", c)
	} else {
		response.Fail(response.SQLErr, c)
	}
}
