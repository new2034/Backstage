/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-25 23:14:03
 * @LastEditTime: 2022-03-06 16:37:22
 */
package model

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	UserID   uint   `gorm:"foreignkey:UserID"`                  // 用户 ID
	Cover    string `json:"cover"`                              // 封面
	Title    string `json:"title" binding:"min=4,max=40"`       // 标题
	Summary  string `json:"summary" binding:"min=20,max=200"`   // 摘要
	Content  string `json:"content" binding:"min=20,max=10000"` // 内容
	View     int    `json:"view" gorm:"default:0"`              // 浏览量
	Like     int    `json:"like" gorm:"default:0"`              // 点赞数
	Link     string `json:"link"`                               // 链接
	Author   string `json:"author"`                             // 作者
	Category int    `json:"category" binding:"required"`        // 分类

	// 自动关联
	User User // 创建者
}
type ResourceList []Resource

/**
 * @description: 资源列表
 */
type ResourceDto struct {
	Id         uint   `json:"id"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Category   int    `json:"category"`
	Author     string `json:"author"` // 发布者
	CreateTime string `json:"create_time"`
	Link       string `json:"link"`
}

/**
 * @description:资源内容
 */
type ResourceIntro struct {
	Author  string `json:"author"`                             // 发布者
	Content string `json:"content" binding:"min=20,max=10000"` // 内容
}

/**
 * @description:  获取前端传来数据
 */
type ResourceSearch struct {
	Username string `form:"username"`
	Title    string `form:"title"`
	Category int    `form:"category"`
}
