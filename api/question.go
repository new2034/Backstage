/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-25 17:41:51
 * @LastEditTime: 2022-03-06 16:10:46
 */
package api

import (
	"Backstage/model"
	"Backstage/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 获取问题列表
 */
func Questionall(c *gin.Context) {
	questionService.Allq(c)
}

/**
 * @description: username，question搜索问题
 */
func Questionsearch(c *gin.Context) {
	reqs := &model.QuestionDetail{}
	if err := c.Bind(&reqs); err != nil {
		response.Fail(http.StatusBadRequest, c)
		return
	}
	req, err := questionService.Search(reqs)
	if err != nil {
		response.Fail(http.StatusBadRequest, c)
		return
	}
	response.OkWithData(req, c)
}

/**
 * @description: id删除问题
 */
func Questionremove(c *gin.Context) {
	qid := c.Query("id")
	questionService.Removeq(qid, c)
}

/**
 * @description: 获取问题详细
 */
func Questiondetails(c *gin.Context) {
	qid := c.Query("id")
	questionService.Detailq(qid, c)
}
