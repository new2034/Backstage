/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-25 14:17:22
 * @LastEditTime: 2022-03-06 16:12:33
 */
package model

import "gorm.io/gorm"

/**
 * @description: 问题表
 */
type Question struct {
	gorm.Model
	UserID    uint   `gorm:"foreignKey:UserID"`                 // 用户 ID
	Question  string `json:"question" binding:"min=20,max=200"` // 问题
	Status    bool   `json:"status" gorm:"default:false"`       // 状态
	AnswerNum int    `json:"answer_num" gorm:"default:0"`       // 回答数
	View      int    `json:"view" gorm:"default:0"`             // 浏览量
	Like      int    `json:"like" gorm:"default:0"`             // 点赞

	// 自动关联
	User User // 创建者
}
type QuestionList []Question

/**
 * @description: 问题列表/问题搜索
 */
type QuestionDto struct {
	ID         uint   `json:"id"`          // 问题id
	Username   string `json:"username"`    // 发布者
	Question   string `json:"question"`    // 问题
	AnswerNum  int    `json:"answer_num"`  // 答案数
	View       int    `json:"view"`        // 问题浏览量
	CreateTime string `json:"create_time"` // 发布时间
}

/**
 * @description: 问题详细页
 */
type QuestionTop struct {
	Username  string `json:"username"`   // 发布者
	Question  string `json:"question"`   // 问题
	View      int    `json:"view"`       // 问题浏览量
	Like      int    `json:"like"`       // 问题点赞数
	AnswerNum int    `json:"answer_num"` // 问题回答数
}

/**
 * @description: 获取前端传来的数据
 */
type QuestionDetail struct {
	Username string `form:"username" `
	Question string `form:"question" `
}
