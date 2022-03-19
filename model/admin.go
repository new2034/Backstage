package model

import (
	"gorm.io/gorm"
)

/**
 * @description: 用户表
 */

type Admin struct {
	gorm.Model
	Adminname string `json:"adminname"`     // 用户名
	Password  string `json:"password"`      // 密码
	Vip       bool   `gorm:"default:false"` // 是否为管理员
}

type AdminDto struct {
	ID        uint   `json:"uuid"`
	Adminname string `json:"adminname"`
	Vip       bool   `json:"vip"`
}

type AdminDetail3 struct {
	ID        uint   `json:"uuid"`
	Adminname string `json:"adminname"`
	Email     string `json:"email"` // 邮箱
	Password  string `json:"password"`
}
type UpdataAdminPwd struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AddAdmin struct {
	Adminname string `json:"adminname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
