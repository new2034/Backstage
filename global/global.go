/*
 * @Author: fzf404
 * @Date: 2022-02-10 10:52:05
 * @LastEditTime: 2022-02-20 19:35:01
 * @Description: ćšć±ćé
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
