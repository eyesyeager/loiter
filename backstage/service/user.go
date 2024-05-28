package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"loiter/app/capability"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/foundation"
	"loiter/config"
	"loiter/constants"
	"loiter/constants/template"
	"loiter/global"
	"loiter/model/entity"
	"loiter/model/po"
	"loiter/model/receiver"
	"loiter/model/returnee"
	"loiter/utils"
	"net/http"
	"strconv"
)

/**
 * @author eyesYeager
 * @date 2023/9/26 15:33
 */

type userService struct {
}

var UserService = userService{}

// DoLogin 用户登录
func (*userService) DoLogin(w http.ResponseWriter, r *http.Request, data receiver.DoLogin) error {
	errorMsg := "账号或密码错误"

	// 获取用户密码
	var checkUser po.LoginUserRole
	if tx := global.MDB.Raw(
		"SELECT user.id Uid, user.password, user.status, role.name Role FROM user, role WHERE user.username = ? AND user.rid = role.id",
		data.Username).Scan(&checkUser); tx.RowsAffected != 1 {
		global.AppLogger.Warn("a user with username ", data.Username, " does not exist")
		return errors.New(errorMsg)
	}

	// 密码校验
	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(data.Password))
	if err != nil {
		global.AppLogger.Warn("the password with username ", data.Username, " is not the given value")
		return errors.New(errorMsg)
	}

	// 状态校验
	if checkUser.Status != constant.Status.Normal.Code {
		return errors.New("账户被冻结，不可登录")
	}

	// 生成token
	if err = foundation.AuthFoundation.RefreshToken(w, checkUser.Uid, checkUser.Role); err != nil {
		global.AppLogger.Warn("token generation failed for user with username ", data.Username, ", error:"+err.Error())
		return errors.New("令牌生成失败，请联系管理员处理")
	}

	// 添加操作日志
	go LogService.Universal(r, checkUser.Uid,
		constant.BuildUniversalLog(constant.LogUniversal.DoLogin, utils.GetIp(r), utils.GetBrowser(r)))
	return err
}

// DoRegister 开通新账号
func (*userService) DoRegister(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.DoRegister) (err error) {
	// 校验操作可行性，高级别用户只可创建低级别用户
	var compareResult int
	if err, compareResult = foundation.RoleFoundation.CompareRole(userClaims.Role, data.Role); err != nil {
		global.AppLogger.Warn("permission judgment error, incorrect data present, error:", err.Error())
		return errors.New("角色非法，系统不存在类型为 " + data.Role + " 的角色")
	}
	if compareResult <= 0 {
		return errors.New("您的权限不足以创建类型为 " + data.Role + " 的角色")
	}

	// 生成随机密码
	initialPsd := utils.GenerateRandString(config.Program.InitialPsdLen)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(initialPsd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	// 创建用户
	_, rid := foundation.RoleFoundation.GetRidByRole(data.Role)
	user := entity.User{
		Username: data.Username,
		Password: string(passwordHash),
		Rid:      rid,
		Email:    data.Email,
		Remarks:  data.Remarks,
	}
	if err = global.MDB.Create(&user).Error; err != nil {
		global.AppLogger.Warn("failed to create user, error:", err.Error())
		return errors.New("创建用户失败，请联系管理员")
	}

	// 发送邮件通知被创建的用户
	// 此处不能异步，因为邮件中含有初始密码，因此必须保证邮件发送成功
	emailTemplate := template.GetRegisterEmailTemplate(data.Username, initialPsd)
	err = capability.NoticeFoundation.SendEmailWithHTMLAndCC(constants.Talisman.WithoutApp, emailTemplate.Subject, []string{data.Email}, []string{}, emailTemplate.Content, true)
	if err != nil {
		global.AppLogger.Error("email sending failed when registering a new user,",
			"username:", data.Username,
			";email:", data.Email,
			";error:", err.Error())
		return errors.New("邮件通知新用户失败！请记住初始密码：" + initialPsd)
	}

	// 记录操作日志
	go LogService.Universal(r, userClaims.Uid,
		constant.BuildUniversalLog(constant.LogUniversal.DoRegister, user.Username, user.Email, user.Remarks))
	return err
}

// GetUserInfo 获取用户信息
func (*userService) GetUserInfo(userClaims utils.JwtCustomClaims) (error, returnee.GetUserInfo) {
	var checkUser entity.User
	if err := global.MDB.First(&checkUser, userClaims.Uid).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), returnee.GetUserInfo{}
	}
	err, weight := foundation.RoleFoundation.GetWeightByRole(userClaims.Role)
	if err != nil {
		return err, returnee.GetUserInfo{}
	}
	return nil, returnee.GetUserInfo{
		Uid:      userClaims.Uid,
		Username: checkUser.Username,
		Weight:   weight,
	}
}

// GetAllUser 获取所有用户
func (*userService) GetAllUser() (err error, res []returnee.GetDictionary) {
	var userList []entity.User
	if err = global.MDB.Where(&entity.User{Status: constant.Status.Normal.Code}).Find(&userList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	for _, item := range userList {
		res = append(res, returnee.GetDictionary{
			Label: item.Username,
			Value: strconv.Itoa(int(item.ID)),
		})
	}
	return err, res
}
