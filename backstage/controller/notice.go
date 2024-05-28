package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/constant"
	"loiter/backstage/controller/parser"
	"loiter/backstage/controller/result"
	"loiter/backstage/controller/validator"
	"loiter/backstage/foundation"
	"loiter/backstage/service"
	"loiter/model/receiver"
	"loiter/utils"
	"net/http"
	"strconv"
)

/**
 * 通知控制器
 * @auth eyesYeager
 * @date 2024/2/23 11:13
 */

// GetNoticeList
// @Summary			分页获取通知列表
// @Description		权限：user
// @Tags			notice
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.GetNoticeList				body		string		false		"请求参数"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/notice/getNoticeList [post]
func GetNoticeList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验
	var data receiver.GetNoticeList
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 分页处理
	if err := utils.CheckPageStruct(data.PageStruct); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err, res := service.NoticeService.GetNoticeList(data); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetEmailNoticeContent
// @Summary			获取邮件通知内容
// @Description		权限：user
// @Tags			notice
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.GetEmailNoticeContent		body		string		false		"请求参数"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/notice/getEmailNoticeContent/:id [get]
func GetEmailNoticeContent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验格式转换
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		result.FailAttachedMsg(w, fmt.Sprintf("id格式错误，error：%s", err.Error()))
		return
	}
	if id <= 0 {
		result.FailAttachedMsg(w, fmt.Sprintf("非法参数，id：%d", id))
		return
	}

	// 执行业务
	if err, res := service.NoticeService.GetEmailNoticeContent(uint(id)); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}
