/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-25 23:43:42
 * @LastEditTime: 2022-03-06 16:11:19
 */
package api

import (
	"Backstage/model"
	"Backstage/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 获取资源列表
 */
func Resourceall(c *gin.Context) {
	resourceService.Allr(c)
}

/**
 * @description: 搜索资源
 */
// 函数命名为驼峰式 AxxxBxxx
func ResourceSearch(c *gin.Context) {
	res := &model.ResourceSearch{}
	if err := c.Bind(&res); err != nil {
		response.Fail(http.StatusBadRequest, c)
		return
	}
	//re, err := resourceService.Search(res)
	//re, err := resourceService.Search1(res)
	re, err := resourceService.Search2(res)
	if err != nil {
		response.Fail(http.StatusBadRequest, c)
		return
	}
	response.OkWithData(re, c)
}

/**
 * @description: 获取资源内容
 */
func Resourcecontent(c *gin.Context) {
	rid := c.Query("id")
	resourceService.Contentr(rid, c)
}

/**
 * @description: 删除资源
 */
func Resourceremove(c *gin.Context) {
	rid := c.Query("id")
	resourceService.Remover(rid, c)
}
