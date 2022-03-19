/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-26 10:20:37
 * @LastEditTime: 2022-03-05 00:13:42
 */
package api

import (
	"Backstage/model"
	"Backstage/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 用过回答者用户名或答案搜索答案
 */
func QuestionDetailsSearch(c *gin.Context) {
	reas := &model.QuestionAnswer{}
	if err := c.Bind(&reas); err != nil {
		response.Fail(http.StatusBadRequest, c)
		return
	}
	rea, err := answerService.Search(reas)
	if err != nil {
		response.Fail(http.StatusBadRequest, c)
		return
	}
	response.OkWithData(rea, c)
}

/**
 * @description: 通过答案id查看答案详细
 */
func Questiondetailsanswer(c *gin.Context) {
	aid := c.Query("answer_id")
	answerService.Detaila(aid, c)
}

/**
 * @description: 通过答案id删除答案
 */
func Questiondetailsremove(c *gin.Context) {
	aid := c.Query("answer_id")
	answerService.Removea(aid, c)
}
