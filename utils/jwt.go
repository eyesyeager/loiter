package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

/**
 * 身份令牌相关工具
 * @author eyesYeager
 * @date 2023/9/27 11:11
 */

// JwtCustomClaims 注册声明是JWT声明集的结构化版本，仅限于注册声明名称
type JwtCustomClaims struct {
	Uid  uint
	Role string
	jwt.RegisteredClaims
}

// GenerateToken 生成Token
func GenerateToken(subject string, stSignKey []byte, uid uint, role string, expire int) (string, error) {
	// 初始化
	iJwtCustomClaims := JwtCustomClaims{
		Uid:  uid,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire) * time.Minute)),
			// 令牌颁发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 主题
			Subject: subject,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)
	return token.SignedString(stSignKey)
}

// ParseToken 解析token
func ParseToken(stSignKey []byte, tokenStr string) (JwtCustomClaims, error) {
	claims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return claims, errors.New("that's not even a token, " + tokenStr)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return claims, errors.New("token is expired, " + tokenStr)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return claims, errors.New("token not active yet, " + tokenStr)
			} else {
				return claims, errors.New("couldn't handle this token, " + tokenStr)
			}
		}
	}
	if !token.Valid {
		err = errors.New("couldn't handle this token")
	}
	return claims, err
}
