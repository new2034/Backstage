/*
 * @Author: fzf404
 * @Date: 2022-02-10 10:44:20
 * @LastEditTime: 2022-02-20 19:49:50
 * @Description: redis 管理
 */
package initialize

import (
	"Backstage/global"
	"log"

	"github.com/go-redis/redis"
)

func Redis() {
	// 获取 redis 配置
	r := global.CONFIG.Redis
	// 建立连接
	client := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})
	// 测试连接
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Connect redis failed: ", err)
	} else {
		global.REDIS_DB = client
		log.Print("Connect redis success: ", pong)
	}
}
