/*
 * @Author: fzf404
 * @Date: 2022-01-22 14:28:51
 * @LastEditTime: 2022-03-10 14:09:42
 * @Description: 用户相关
 */

package model

import (
	"Backstage/global"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

/**
 * @description: users 数据表
 */
type User struct {
	gorm.Model
	UUID        uuid.UUID // uuid
	Username    string    // 用户名
	Password    string    // 密码
	Email       string    // 邮箱
	IsSuper     bool      `gorm:"default:false"` // 是否为超级管理员
	QuestionNum int       `gorm:"default:0"`     // 问题数
	AnswerNum   int       `gorm:"default:0"`     // 回答数
	ArticleNum  int       `gorm:"default:0"`     // 文章数
	ResourceNum int       `gorm:"default:0"`     // 资源数
}

/**
 * @description: 管理员表
 */
// type Admin struct {
// 	gorm.Model
// 	Adminname string
// 	Password  string
// 	IsSuper   bool `gorm:"default:false"` // 是否为超级管理员
// }

type UserDto struct {
	ID       uint      `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	IsSuper  bool      `json:"is_super"`
	Avatar   string    `json:"avatar"`
}


type UserDetail struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"` // 邮箱
	IsSuper  bool   `json:"is_super"`
}

type UserDetail2 struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"` // 邮箱
	Password string `json:"password"`
}
type UpdataUserPwd struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type DeleteUser struct {
	ID       uint   `json:"nid"`
	Password string `json:"password"`
}

/**
 * @description: 通过 ID 查找用户
 * @param {uint} id
 */
func (u *User) GetByID(id interface{}) bool {
	return global.MYSQL_DB.Where("id = ?", id).First(&u).Error == nil
}

/**
 * @description: 通过用户名查找用户
 * @param {string} username
 */
//通过username查找用户
func (u *User) GetByUsername(username string) error {
	if result := global.MYSQL_DB.Where("username = ?", username).First(&u); result.Error != nil {
		fmt.Printf("err:%s", result.Error)
		return result.Error
	}
	return nil
}

/**
 * @description: 通过邮箱查找用户
 * @param {string} email
 */
func (u *User) GetByEmail(email string) bool {
	return global.MYSQL_DB.Where("email = ?", email).First(&u).Error == nil
}

/**
 * @description: 更新用户信息
 */
func (u *User) Update() bool {
	return global.MYSQL_DB.Save(&u).Error == nil
}

/**
 * @description: 转换为 Dto
 */
func (u *User) Dto() UserDto {
	return UserDto{
		ID:       u.ID,
		Username: u.Username,
		IsSuper:  u.IsSuper,
		//	Avatar:   u.Avatar,
	}
}
func (u *User) Create() {
	global.MYSQL_DB.Create(&u)
}