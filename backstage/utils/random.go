package utils

import (
	"encoding/base64"
	"math/rand"
	"time"
)

/**
 * 随机数工具
 * @author eyesYeager
 * @date 2023/11/26 20:39
 */

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandString 生成指定长度的随机字符串
// 伪随机，安全性不高
func GenerateRandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// GenerateRandStringCrypto 生成指定长度的随机字符串
// 真随机，安全性较高，可以用于密钥生成
func GenerateRandStringCrypto(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
