package model

import "Backstage/global"

//email搜索用户
func (u *User) SearchUser(email string) bool {
	return global.MYSQL_DB.Where("email like ?", "%"+email+"%").Find(&u).Error == nil
}
func (u *User) SearchUserID(id uint) bool {
	return global.MYSQL_DB.Where("id= ?", id).Find(&u).Error == nil
}

//email搜索管理员
func (t *Admin) SearchAdmin(email string) bool {
	return global.MYSQL_DB.Where("email like ?", "%"+email+"%").Find(&t).Error == nil
}

//name搜索管理员
func (t *Admin) SearchAdmin1(name string) bool {
	return global.MYSQL_DB.Where("adminname = ?", name).First(&t).Error == nil
}

//管理员更新
func (t *Admin) Update() bool {
	return global.MYSQL_DB.Save(&t).Error == nil
}

func (t *Admin) Create() bool {
	return global.MYSQL_DB.Create(&t).Error == nil
}
