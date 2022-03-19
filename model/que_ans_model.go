/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-28 23:56:57
 * @LastEditTime: 2022-03-06 15:04:46
 */
package model

/**
 * @description: 问答关系表
 */
type QueAns struct {
	QuestionID uint `gorm:"foreignKey:QuestionID" `
	AnswerID   uint `gorm:"foreignKey:AnswerID" `

	// 自动关联
	Question Question
	Answer   Answer
}
type QueAnsList []QueAns

/**
 * @description: 答案列表
 */
type AnsDto struct {
	Username   string `json:"replier"`    // 回答者
	Answer     string `json:"answer"`     // 答案内容
	Id         int    `json:"answer_id"`  //答案序号
	Like       int    `json:"a_like"`     // 答案点赞数
	CreateTime string `json:"creat_time"` // 发布时间
}
