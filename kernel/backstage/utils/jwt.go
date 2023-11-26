package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"loiter/config"
	"time"
)

/**
 * 身份令牌相关工具
 * @author eyesYeager
 * @date 2023/9/27 11:11
 */

// 把签发的秘钥 抛出来
var stSignKey = []byte(config.Program.JWTSecretKey)

// JwtCustomClaims 注册声明是JWT声明集的结构化版本，仅限于注册声明名称
type JwtCustomClaims struct {
	Id               uint
	Role             uint
	RegisteredClaims jwt.RegisteredClaims
}

func (j JwtCustomClaims) Valid() error {
	return nil
}

// GenerateToken 生成Token
func GenerateToken(uid uint, weight uint) (string, error) {
	// 初始化
	iJwtCustomClaims := JwtCustomClaims{
		Id:   uid,
		Role: weight,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Program.JWTExpire) * time.Minute)),
			// 令牌颁发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 主题
			Subject: config.Program.Name,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)
	return token.SignedString(stSignKey)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (JwtCustomClaims, error) {
	iJwtCustomClaims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})

	if err == nil && !token.Valid {
		err = errors.New("invalid Token")
	}
	return iJwtCustomClaims, err
}

// IsTokenValid 检查token是否有效
func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	return err == nil
}
