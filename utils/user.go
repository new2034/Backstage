/*
 * @Author: fzf404
 * @Date: 2022-02-10 13:38:36
 * @LastEditTime: 2022-02-11 00:42:46
 * @Description: 用户管理
 */
package utils

import (
	"Backstage/global"
	"Backstage/model"
)

/**
 * @description: 判断邮箱是否存在
 * @param {string} email
 */
func IsEmailExist(email string) bool {
	var count int64
	global.MYSQL_DB.Model(model.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
func IsEmailExist2(email string) bool {
	var count int64
	global.MYSQL_DB.Model(model.Admin{}).Where("email = ?", email).Count(&count)
	return count > 0
}
func IsUsernameExist(username string) bool {
	var count int64
	global.MYSQL_DB.Model(model.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}
func IsAdminnameExist(adminname string) bool {
	var count int64
	global.MYSQL_DB.Model(model.Admin{}).Where("adminname = ?", adminname).Count(&count)
	return count > 0
}