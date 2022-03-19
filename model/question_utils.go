/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-25 15:22:05
 * @LastEditTime: 2022-03-04 21:19:04
 */
package model

import (
	"Backstage/global"

	"gorm.io/gorm/clause"
)

/**
 * @description: 问题列表/问题搜索
 */
func (q *Question) QDto() QuestionDto {
	return QuestionDto{
		ID:         q.ID,
		Username:   q.User.Username,
		Question:   q.Question,
		AnswerNum:  q.AnswerNum,
		View:       q.View,
		CreateTime: q.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

/**
 * @description: 问题详细
 */
func (q *Question) QTop() QuestionTop {
	return QuestionTop{
		Username:  q.User.Username,
		Question:  q.Question,
		View:      q.View,
		Like:      q.Like,
		AnswerNum: q.AnswerNum,
	}
}

/**
 * @description: 获取问题列表，数据传输
 */
func (ql QuestionList) QQDto() []QuestionDto {
	var qlDto []QuestionDto
	for _, q := range ql {
		qlDto = append(qlDto, q.QDto())
	}
	return qlDto
}

/**
 * @description: 获取全部问题
 */
func (ql *QuestionList) Get() bool {
	return global.MYSQL_DB.Order("created_at desc").Preload(clause.Associations).Find(&ql).Error == nil
}

/**
 * @description: 问题搜索
 */
//question搜索
func (ql *QuestionList) Searchq(question string) error {
	if result := global.MYSQL_DB.Where("question like ?", "%"+question+"%").Order("created_at desc").Preload(clause.Associations).Find(&ql); result.Error == nil {
		return result.Error
	}
	return nil
}

// userid查询question(取消question封装)改用函数调用
func GetByUid(uid uint, q *QuestionList) error {
	if result := global.MYSQL_DB.Where("user_id= ?", uid).Order("created_at desc").Preload(clause.Associations).Find(&q); result.Error == nil {
		return result.Error
	}
	return nil
}

/**
 * @description: 通过 questionid 查找问题
 */
func (q *Question) GetByQID(qid interface{}) bool {
	return global.MYSQL_DB.Where("id = ?", qid).Order("created_at desc").Preload(clause.Associations).First(&q).Error == nil
}

/**
 * @description: 通过id删除问题
 */
func (q *Question) Removeq(qid interface{}) bool {
	return global.MYSQL_DB.Where("id = ?", qid).Delete(&q).Error == nil
}
