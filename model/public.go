/*
 * @Description:
 * @Autor: Zouly
 * @Date: 2022-03-06 16:35:10
 * @LastEditTime: 2022-03-06 16:39:11
 */
package model

import (
	"database/sql/driver"
	"encoding/json"
)

/**
 * @description: tag 存入前转为 string
 * @param {interface{}} value
 */
func (t *Tag) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

// 读出前转为 json
func (t Tag) Value() (driver.Value, error) {
	return json.Marshal(t)
}
