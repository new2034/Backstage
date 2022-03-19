//main函数调用初始化

package main

import (
	"Backstage/global"
	"Backstage/initialize"
	"Backstage/router"
)

/**
 * @description: 初始化
 */
func init() {
	initialize.Config() // 初始化配置
	initialize.Mysql()  // 初始化数据库
	initialize.Redis()  // 初始化数据库
}

/**
 * @description: 主函数
 */
func main() {
	r := router.InitRouter()               // 初始化路由
	r.Run(":" + global.CONFIG.Common.Port) // 启动服务
}
