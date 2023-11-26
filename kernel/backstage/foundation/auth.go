package foundation

/**
 * 权限服务
 * @author eyesYeager
 * @date 2023/11/26 16:23
 */

type authService struct {
}

var AuthService = authService{}

// TokenAnalysis 令牌解析
func (*authService) TokenAnalysis() {

}

// IdentityVerification 身份认证（邮件）
func (*authService) IdentityVerification() {

}

// emailVerification 身份认证——邮箱验证码
func emailVerification() {

}

// smsVerification 身份认证——短信验证码
func smsVerification() {

}
