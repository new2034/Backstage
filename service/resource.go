/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-25 23:59:48
 * @LastEditTime: 2022-03-19 09:27:23
 */
package service

import (
	"Backstage/model"
	"Backstage/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

/**
 * @description:资源列表
 */
func (rs *ResourceService) Allr(c *gin.Context) {
	var rl model.ResourceList
	if !rl.Get() {
		response.Fail(response.SQLErr, c)
		return
	}
	if len(rl) == 0 {
		response.FailWithMessage(response.Success, "暂无资源", c)
		return
	}
	response.OkWithData(rl.RDto(), c)
}

/**
 * @description:搜索资源
 */
func (rs *ResourceService) Search(search *model.ResourceSearch) (data interface{}, err error) {
	rl := &model.ResourceList{}
	//if len(search.Username) > 0 {
	//	u := &model.User{}
	//	if err = u.GetByUsername(search.Username); err != nil {
	//		return
	//	}
	//}
	//	if err = model.GetOnUid(u.ID, rl); err != nil {
	//		return
	//	}
	//} else if err = rl.Searchrt(search.Title); err != nil {
	//	return
	//} else {
	//	if err = rl.Searchrc(search.Category); err != nil {
	//		return
	//	}
	//}

	// sqlString:= "select * from resources where  1=1"
	//
	//if UserID.Text.Lenght > 0 {
	//	QuerySqlStr += "and user_id=" + "'UserID.Text'"
	//}
	//if Category.Text.Lenght > 0 {
	//	QuerySqlStr += "and category=" + "'Category.Text'"
	//}
	//
	//if Title.Text.Lenght > 0 {
	//	QuerySqlStr += "and title=" + "'Title.Text'"
	//}
	return rl.RDto(), nil
}

func (rs *ResourceService) Search1(search *model.ResourceSearch) (data interface{}, err error) {
	rl := &model.ResourceList{}
	var params []interface{}
	whereSql := " 1=1 "
	if len(search.Username) > 0 {
		whereSql += " AND user_id = ? "
		params = append(params, search.Username)
	}

	if search.Category > 0 {
		whereSql += " AND category = ? "
		params = append(params, search.Category)
	}

	if len(search.Title) > 0 {
		whereSql += " AND title LIKE '%"+search.Title+"%'"
	}
	if err = model.SearchMethod("resources", whereSql, params, &rl); err != nil {
		fmt.Println(err)
		return
	}
	return rl.RDto(), nil
}

func (rs *ResourceService) Search2(search *model.ResourceSearch) (data interface{}, err error) {
	rl := &model.ResourceList{}
	var params []interface{}
	whereSql2 := "SELECT * FROM resources WHERE 1=1 "
	if len(search.Username) > 0 {
		whereSql2 += " AND user_id = ? "
		params = append(params, search.Username)
	}

	if search.Category > 0 {
		whereSql2 += " AND category = ? "
		params = append(params, search.Category)
	}

	if len(search.Title) > 0 {
		whereSql2 += " AND title LIKE '%"+search.Title+"%'"
	}

	if err = rl.SearchMethod(whereSql2, params); err != nil {
		fmt.Println(err)
		return
	}
	return rl.RDto(), nil
}


/**
 * @description:根据id移除资源
 */
func (rs *ResourceService) Remover(rid string, c *gin.Context) {
	var r model.Resource
	if !r.GetByID(rid) {
		response.FailWithMessage(response.Success, "资源不存在", c)
		return
	}
	if r.Delete(rid) {
		response.OkWithMessage("删除成功", c)
	} else {
		response.Fail(response.SQLErr, c)
	}
}

/**
 * @description:资源内容
 */
func (rs *ResourceService) Contentr(rid string, c *gin.Context) {
	var r model.Resource
	if r.GetByID(rid) {
		response.OkWithData(r.Intro(), c)
	} else {
		response.FailWithMessage(response.Success, "资源内容不存在", c)
	}
}
