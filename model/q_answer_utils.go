/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-02-25 16:01:32
 * @LastEditTime: 2022-03-06 15:02:33
 */
package model

import (
	"Backstage/global"

	"gorm.io/gorm/clause"
)

/**
 * @description: 答案具体内容
 */
func (a *Answer) ADto() AnswerDto {
	return AnswerDto{
		Username: a.User.Username,
		Answer:   a.Answer,
	}
}

/**
 * @description: 获取答案列表
 */
func (a *Answer) AIntro() AnswerIntro {
	return AnswerIntro{
		Username:   a.User.Username,
		Answer:     a.Answer,
		Id:         int(a.ID),
		Like:       a.Like,
		CreateTime: a.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

/**
 * @description: 答案列表数据传输
 */
func (al AnswerList) Dto() []AnswerIntro {
	var alDto []AnswerIntro
	for _, a := range al {
		alDto = append(alDto, a.AIntro())
	}
	return alDto
}

/**
 * @description: 通过答案ID 查找答案
 */
func (a *Answer) GetByAID(aid interface{}) bool {
	return global.MYSQL_DB.Where("id = ?", aid).Preload(clause.Associations).First(&a).Error == nil
}

/**
 * @description:通过id删除答案
 */
func (a *Answer) Removean(aid interface{}) bool {
	return global.MYSQL_DB.Where("id = ?", aid).Delete(&a).Error == nil
}

/**
 * @description: 在某个问题下，通过username或answer搜索答案
 */
// answer搜索
func (al *AnswerList) Searcha(aid uint, answer string) error {
	if result := global.MYSQL_DB.Where("id=? AND answer like ?", aid, "%"+answer+"%").Order("created_at desc").Preload(clause.Associations).Find(&al); result.Error != nil {
		return result.Error
	}
	return nil
}

// 通过username在user中搜索，取出user_id,通过user_id在answerlist中搜索
func SearchByUid(aid uint, uid uint, al *AnswerList) error {
	if result := global.MYSQL_DB.Where("id=? AND user_id= ?", aid, uid).Order("created_at desc").Preload(clause.Associations).Find(&al); result.Error != nil {
		return result.Error
	}
	return nil
}
