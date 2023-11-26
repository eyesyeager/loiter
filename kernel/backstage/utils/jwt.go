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
	Uid              uint
	Weight           uint
	RegisteredClaims jwt.RegisteredClaims
}

func (j JwtCustomClaims) Valid() error {
	return nil
}

// GenerateToken 生成Token
func GenerateToken(subject string, stSignKey []byte, uid uint, weight uint, expire int) (string, error) {
	// 初始化
	iJwtCustomClaims := JwtCustomClaims{
		Uid:    uid,
		Weight: weight,
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
	iJwtCustomClaims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})

	if err == nil && !token.Valid {
		err = errors.New("invalid Token:" + tokenStr)
	}
	return iJwtCustomClaims, err
}
