/*
 * @Author: fzf404
 * @Date: 2022-02-10 22:57:39
 * @LastEditTime: 2022-02-11 13:29:04
 * @Description: 鉴权相关
 */
package model

import (
	"github.com/golang-jwt/jwt"
)

// jwt
type Claims struct {
	UID      uint      // 用户 ID
	Username string    // 用户名
    IsSuper  bool
	jwt.StandardClaims
}

type Claims1 struct {
	UID      uint   // 用户 ID
	Username string // 用户名
	Vip      bool
	Password string
	jwt.StandardClaims
}