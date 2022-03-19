/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-25 23:19:53
 * @LastEditTime: 2022-03-19 08:57:11
 */
package model

import (
	"Backstage/global"

	"gorm.io/gorm/clause"
)

/**
 * @description: 资源列表
 */
func (r *Resource) Dto() ResourceDto {
	return ResourceDto{
		Id:         r.ID,
		Link:       r.Link,
		Title:      r.Title,
		Summary:    r.Summary,
		Category:   r.Category,
		Author:     r.User.Username,
		CreateTime: r.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

/**
 * @description: 资源内容详情
 */
func (r *Resource) Intro() ResourceIntro {
	return ResourceIntro{
		Author:  r.User.Username,
		Content: r.Content,
	}
}

/**
 * @description: 数据传输
 */
func (rl ResourceList) RDto() []ResourceDto {
	var rlDto []ResourceDto
	for _, r := range rl {
		rlDto = append(rlDto, r.Dto())
	}
	return rlDto
}

/**
 * @description: 通过资源ID查找资源
 */
func (r *Resource) GetByID(rid interface{}) bool {
	return global.MYSQL_DB.Where("id = ?", rid).Preload(clause.Associations).First(&r).Error == nil
}

/**
 * @description: 获取全部资源列表
 */
func (rl *ResourceList) Get() bool {
	return global.MYSQL_DB.Order("created_at desc").Preload(clause.Associations).Find(&rl).Error == nil
}

/**
 * @description: 搜索资源
 */
// title搜索
func (rl *ResourceList) Searchrt(title string) error {
	if result := global.MYSQL_DB.Where("title like ?  ", "%"+title+"%").Order("created_at desc").Preload(clause.Associations).Find(&rl); result.Error != nil {
		return result.Error
	}
	return nil
}

// category 搜索
func (rl *ResourceList) Searchrc(category int) error {
	if result := global.MYSQL_DB.Where(" category = ? ", category).Order("created_at desc").Preload(clause.Associations).Find(&rl); result.Error != nil {
		return result.Error
	}
	return nil
}


/*
 ResourceList类方法
	whereSql   查询语句
	params     查询参数
*/

func (rl *ResourceList) SearchMethod(sqlString string, params []interface{}) error {
	if result := global.MYSQL_DB.Raw(sqlString, params).Scan(&rl); result.Error != nil {
		return result.Error
	}
	return nil
}

/*
 公共方法
	tableName  表名
	whereSql   查询语句
	params     查询参数
	res        返回参数
*/

func SearchMethod(tableName, whereSql string, params []interface{},res interface{}) error {
	if result := global.MYSQL_DB.Table(tableName).Where(whereSql, params...).Scan(&res); result.Error != nil {
		return result.Error
	}
	return nil
}

// username搜索，间接采用user_id搜索
func GetOnUid(uid uint, rl *ResourceList) error {
	if result := global.MYSQL_DB.Where("user_id= ?", uid).Order("created_at desc").Preload(clause.Associations).Find(&rl); result.Error != nil {
		return result.Error
	}
	return nil
}

/**
 * @description: 通过id删除资源
 */
func (r *Resource) Delete(rid interface{}) bool {
	return global.MYSQL_DB.Where("id = ?", rid).Delete(&r).Error == nil
}
