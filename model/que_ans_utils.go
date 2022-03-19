/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-03-01 15:07:25
 * @LastEditTime: 2022-03-06 15:07:57
 */
package model

import (
	"Backstage/global"
	"fmt"

	"gorm.io/gorm/clause"
)

/**
 * @description: 答案列表
 */
func (qa *QueAns) A() AnsDto {
	return AnsDto{
		Username:   qa.Answer.User.Username,
		Answer:     qa.Answer.Answer,
		Id:         int(qa.Answer.ID),
		Like:       qa.Answer.Like,
		CreateTime: qa.Answer.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

/**
 * @description: 答案列表 数据传输
 */
func (qal QueAnsList) ADto() []AnsDto {
	var alDto []AnsDto
	for _, qa := range qal {
		alDto = append(alDto, qa.A())
	}
	return alDto
}

/**
 * @description: 通过 questionid 查找   问题详情列表
 */
// 嵌套预加载，查询出username
func (qal *QueAnsList) GetByQID(qid interface{}) bool {
	return global.MYSQL_DB.Where("question_id = ?", qid).Preload("Answer.User").Preload("Question.User").Preload(clause.Associations).Find(&qal).Error == nil
}

/**
 * @description: 删除问题或答案时，同步删除关系表中的数据
 */
// 删除answer_id
func (qa *QueAns) Removea(aid interface{}) bool {
	return global.MYSQL_DB.Where("answer_id = ?", aid).Delete(&qa).Error == nil
}

// 删除question_id
func (qa *QueAns) Removeq(qid interface{}) bool {
	return global.MYSQL_DB.Where("question_id = ?", qid).Delete(&qa).Error == nil
}

/**
 * @description: 通过qid搜索对应的aid  确保是该问题下的答案
 */
func (qal *QueAnsList) SearchByQID(qid interface{}) error {
	if result := global.MYSQL_DB.Table("que_ans").Where("question_id=?", qid).Preload("Answer.User").Preload(clause.Associations).Find(&qal); result.Error != nil {
		fmt.Printf("err:%s", result.Error)
		return result.Error
	}
	return nil
}
