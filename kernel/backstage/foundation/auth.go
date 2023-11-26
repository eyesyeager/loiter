package foundation

import (
	"errors"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/utils"
	"net/http"
	"strconv"
)

/**
 * 权限服务
 * @author eyesYeager
 * @date 2023/11/26 16:23
 */

type authFoundation struct {
	subject   string // jwt主题
	stSignKey []byte // jwt密钥
	expire    int    // 过期时间
}

var AuthFoundation = authFoundation{
	config.Program.Name,
	[]byte(config.Program.JWTSecretKey),
	config.Program.JWTExpire,
}

// TokenAnalysis 令牌解析
func (authFoundation *authFoundation) TokenAnalysis(w http.ResponseWriter, r *http.Request, role uint) (userClaims utils.JwtCustomClaims, err error) {
	token := r.Header.Get(constant.ResponseHead.Token)
	// 处理未携带令牌的情况
	if token == "" {
		if constant.Role.Visitor == role {
			// 接口权限等级为 constant.Role.Visitor 时，无需令牌
			return userClaims, nil
		}
		return userClaims, errors.New("权限不足，请先登录")
	}
	// 解析令牌
	if userClaims, err = utils.ParseToken(authFoundation.stSignKey, token); err != nil {
		global.BackstageLogger.Warn("failed to parse token, error:" + err.Error())
		return userClaims, errors.New("权限失效，请重新登录")
	}
	// 刷新令牌
	if refreshErr := authFoundation.RefreshToken(w, userClaims.Uid, userClaims.Weight); refreshErr != nil {
		// 令牌刷新失败不阻塞业务
		global.BackstageLogger.Warn("token refresh failed for user with uid " + strconv.Itoa(int(userClaims.Uid)) + ", error:" + refreshErr.Error())
	}
	// 权限判断
	if userClaims.Weight < role {
		global.BackstageLogger.Warn("uid(" + strconv.Itoa(int(userClaims.Uid)) + ") has insufficient permissions, with a weight of " + strconv.Itoa(int(userClaims.Weight)) + ", which is less than " + strconv.Itoa(int(role)))
		return userClaims, errors.New("权限不足，请先提升权限")
	}
	return userClaims, err
}

// RefreshToken 刷新令牌
func (authFoundation *authFoundation) RefreshToken(w http.ResponseWriter, uid uint, weight uint) error {
	// 生成令牌
	token, err := utils.GenerateToken(authFoundation.subject, authFoundation.stSignKey, uid, weight, authFoundation.expire)
	if err != nil {
		global.BackstageLogger.Warn("user with uid " + strconv.Itoa(int(uid)) + " and weight " + strconv.Itoa(int(weight)) + " failed to generate token, error:" + err.Error())
		return errors.New("身份令牌生成失败")
	}
	// 写入响应头
	w.Header().Set(constant.ResponseHead.Token, token)
	return nil
}

// IdentityVerification 身份认证（邮件）
func (*authFoundation) IdentityVerification() {

}

// emailVerification 身份认证——邮箱验证码
func emailVerification() {

}

// smsVerification 身份认证——短信验证码
func smsVerification() {

}
