/*
 * @Author: fzf404
 * @Date: 2022-01-22 14:53:01
 * @LastEditTime: 2022-03-10 11:40:00
 * @Description: token 处理
 */
package utils

import (
	"Backstage/global"
	"Backstage/model"
	"time"

	"github.com/golang-jwt/jwt"
)

/**
 * @description: 签名密钥
 */
var jwtKey = []byte(global.CONFIG.JWT.Key)

/**
 * @description: 分发 Token
 * @param {model.User} user
 */
func ReleaseToken(a model.Admin) (string, error) {

	// token 结构生成
	claims := &model.Claims{
		// 使用 ID、Username 作为有效载荷
		UID: a.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.Expires, // 签名过期时间
			NotBefore: time.Now().Unix() - 1000,                      // 签名生效时间
			Issuer:    global.CONFIG.JWT.Issuer,                      // 签名发行人
		},
	}

	// 将 Claims 加密存储为 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/**
 * @description: 解析 Token
 * @param {string} tokenString
 */
func ParseToken(tokenString string) (*jwt.Token, *model.Claims, error) {
	claims := &model.Claims{}
	// 解码
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
func ParseToken1(tokenString string) (*jwt.Token, *model.Claims1, error) {
	claims1 := &model.Claims1{}
	// 解码
	token, err := jwt.ParseWithClaims(tokenString, claims1, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims1, err
}
