/*
 * @Author: fzf404
 * @Date: 2022-01-22 13:52:00
 * @LastEditTime: 2022-02-15 22:43:11
 * @Description: 配置管理
 */

package initialize

import (
	"Backstage/global"
	"log"

	"github.com/spf13/viper"
)

/**
 * @description: 初始化配置文件
 */
func Config() {
	// 文件路径 
	viper.AddConfigPath("./config")
	// 文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Read config failed: ", err)
	}
	// 解析至结构体
	if err := viper.Unmarshal(&global.CONFIG); err != nil {
		log.Fatal("Unmarshal config failed: ", err)
	}
}
