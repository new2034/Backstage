package model

import (
	"Backstage/global"
)

// //通过id查找用户
// func (u *User) GetByID(id interface{}) bool {
// 	return global.MYSQL_DB.Where("id = ?", id).First(&u).Error == nil
// }

//通过email查找管理员
func (t *Admin) GetByEmail(email string) bool {
	return global.MYSQL_DB.Where("email = ?", email).First(&t).Error == nil
}

//通过id查找管理员
func (t *Admin) GetByID(id uint) bool {
	return global.MYSQL_DB.Where("id = ?", id).First(&t).Error == nil
}

//通过email删除管理员
func (t *Admin) DeleteByEmail(email string) bool {
	return global.MYSQL_DB.Where("email = ?", email).Delete(&t).Error == nil
}

//通过id删除用户
func (u *User) DeleteByID(id interface{}) bool {
	return global.MYSQL_DB.Where("id = ?", id).Delete(&u).Error == nil
}

func (u *User) Detail2() UserDetail2 {
	return UserDetail2{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}
func (t *Admin) Detail3() AdminDetail3 {
	return AdminDetail3{
		ID:        t.ID,
		Adminname: t.Adminname,
		Password:  t.Password,
	}
}

//全部用户数据返回
func AllBack(ua *[]User) bool {
	return global.MYSQL_DB.Find(&ua).Error == nil
}

//全部管理员数据返回
func AllBack2(ta *[]Admin) bool {
	return global.MYSQL_DB.Find(&ta).Error == nil
}
