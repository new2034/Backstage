/*
 * @Author: fzf404
 * @Date: 2022-02-10 10:52:05
 * @LastEditTime: 2022-02-20 19:35:01
 * @Description: 全局变量
 */
package global

import (
	"Backstage/config"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	MYSQL_DB *gorm.DB
	REDIS_DB *redis.Client
	CONFIG   config.Config
)
