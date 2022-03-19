package model

import (
	"time"

	"gorm.io/gorm"
)

type ArticleList []Article

type ArticleIntro struct {
	ID         uint   `json:"id"`
	Username   string `json:"writer"`
	Title      string `json:"title"`
	Created_at string `json:"createtime"`
	Like       int    `json:"like"`
	View       int    `json:"pageviews"`
}

type Data struct {
	Date   string `json:"date"`
	Number int64  `json:"number"`
}

/**
 * @description: 数据转换
 */
func (al ArticleList) Dto() []ArticleIntro {

	var alDto []ArticleIntro

	for _, a := range al {
		alDto = append(alDto, a.Intro())
	}

	return alDto
}

func (a *Article) Intro() ArticleIntro {
	return ArticleIntro{
		ID:         a.ID,
		Title:      a.Title,
		Username:   a.User.Username,
		Created_at: TimeDto(a.CreatedAt),
		Like:       a.Like,
		View:       a.View,
	}
}

func TimeDto(t time.Time) string {
	//now := time.Now()
	formatNow := t.Format("2006-01-02 15:04:05")
	return formatNow
}

type Article struct {
	gorm.Model
	UserID uint `gorm:"foreignkey:UserID"` // 用户 ID

	Cover    string `json:"cover"`                              // 封面
	Title    string `json:"title" binding:"min=4,max=40"`       // 标题
	Summary  string `json:"summary" binding:"min=20,max=200"`   // 摘要
	Content  string `json:"content" binding:"min=20,max=10000"` // 内容
	View     int    `json:"view" gorm:"default:0"`              // 浏览量
	Like     int    `json:"like" gorm:"default:0"`              // 点赞数
	Tag      string `json:"tag" binding:"max=5,omitempty"`      // 标签
	Category string `json:"category" binding:"required"`        // 分类

	User User // 创建者
}

type CommentAt struct {
	gorm.Model
	UserID uint `gorm:"foreignkey:UserID"` // 用户 ID

	ContentAt string `json:"content_at" binding:"min=20,max=10000"` // 评论详情

	// 自动关联
	User User // 评论者
}

type CommentRe struct {
	gorm.Model
	UserID uint `gorm:"foreignkey:UserID"` // 用户 ID

	ContentRe string `json:"content_at" binding:"min=20,max=10000"` // 评论详情

	// 自动关联
	User User // 评论者
}

type ArtCom struct {
	ArticleID    uint `gorm:"forignkey:ArticleID"`
	ArticleComID uint `gorm:"forignkey:ArticleComID"`

	// 自动关联
	Article    Article
	ArticleCom ArticleCom
}

type ArticleCom struct {
	gorm.Model
	UserID uint `gorm:"foreignkey:UserID"` // 用户 ID

	Content string `json:"content"`

	// 自动关联
	User User // 创建者
}
