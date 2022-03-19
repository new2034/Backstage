/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-25 14:17:59
 * @LastEditTime: 2022-03-06 15:32:58
 */
package model

import "gorm.io/gorm"

/**
 * @description: 回答表
 */
type Answer struct {
	gorm.Model
	UserID uint   `gorm:"foreignKey:UserID"`                 // 用户 ID
	Answer string `json:"answer" binding:"min=20,max=10000"` // 回答
	Like   int    `json:"like" gorm:"default:0"`             // 点赞

	// 自动关联
	User User // 创建者
}
type AnswerList []Answer

/**
 * @description: 答案详细
 */
type AnswerDto struct {
	Username string `json:"replier"` // 回答者
	Answer   string `json:"answer"`
}

/**
 * @description: 答案数据
 */
type AnswerIntro struct {
	Username   string `json:"replier"`     // 回答者
	Answer     string `json:"answer"`      // 答案内容
	Id         int    `json:"answer_id"`   //答案序号
	Like       int    `json:"a_like"`      // 答案点赞数
	CreateTime string `json:"create_time"` // 发布时间
}

/**
 * @description: 接收前端数据
 */
type QuestionAnswer struct {
	QuestionId string `form:"id" binding:"required"`
	Replier    string `form:"replier"`
	Answer     string `form:"answer"`
}
