/*
 * @Author: fzf404
 * @Date: 2022-01-22 13:56:31
 * @LastEditTime: 2022-03-10 11:24:21
 * @Description: mysql 管理
 */

package initialize

import (
	"Backstage/global"
	"Backstage/model"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @description: 初始化数据库
 */

func Mysql() {
	// 获取 mysql 配置
	mCfg := global.CONFIG.Mysql.Dsn()
	// 连接
	if db, err := gorm.Open(mysql.Open(mCfg),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}); err != nil {
		log.Fatal("Connect mysql failed: ", err)
	} else {
		// 自动建表
		db.AutoMigrate(&model.User{})  // 用户
		db.AutoMigrate(&model.Admin{}) // 管理员

		// 文章
		db.AutoMigrate(&model.Article{})    // 文章
		db.AutoMigrate(&model.ArticleCom{}) // 文章评论
		db.AutoMigrate(&model.ArtCom{})     // 文章与评论关系表

		// 资源
		db.AutoMigrate(&model.Resource{}) // 资源
		// db.AutoMigrate(&model.ResourceCom{}) // 资源评论
		// db.AutoMigrate(&model.ResCom{})      //资源与评论关系表

		// 问答
		db.AutoMigrate(&model.Question{}) // 问题
		db.AutoMigrate(&model.Answer{})   // 回答
		db.AutoMigrate(&model.QueAns{})   // 问答关系表
		global.MYSQL_DB = db
		log.Print("Connect mysql success: ", mCfg)
	}
}
