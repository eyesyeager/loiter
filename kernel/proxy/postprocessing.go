package proxy

import (
	"fmt"
	"loiter/constants"
	"loiter/global"
	"loiter/kernel/aid"
	"loiter/kernel/exception"
	"loiter/kernel/final"
	"loiter/utils"
	"net/http"
)

/**
 * 代理后置处理
 * @auth eyesYeager
 * @date 2024/1/29 15:32
 */

// post 后置处理总入口
func post(w http.ResponseWriter, req *http.Request, resp *http.Response, host string, entrance string, errInfo string, genre string) {
	// 进入响应处理器
	if isResponse(entrance) {
		entryAid(w, req, resp, host, genre)
	}
	// 进入异常处理器
	if isException(entrance) {
		entryException(w, req, host, errInfo, genre)
	}
	// 进入最终处理器
	entryFinal(w, req, resp, host, entrance, errInfo, genre)
}

// entryAid 进入响应处理器
func entryAid(w http.ResponseWriter, req *http.Request, resp *http.Response, host string, genre string) {
	if err := aid.Entry(w, req, resp, host); err != nil {
		errMsg := fmt.Sprintf("aid execution failed. Error message: %s", err.Error())
		statusCode, contentType, content := utils.ResponseTemplate(constants.ResponseTitle.BadGateway, errMsg, genre)
		utils.Response(w, statusCode, contentType, content)
		global.GatewayLogger.Warn(errMsg)
	}
}

// entryException 进入异常处理器
func entryException(w http.ResponseWriter, req *http.Request, host string, errInfo string, genre string) {
	if err := exception.Entry(w, req, host, errInfo); err != nil {
		errMsg := fmt.Sprintf("exception execution failed. Error message: %s", err.Error())
		statusCode, contentType, content := utils.ResponseTemplate(constants.ResponseTitle.BadGateway, errMsg, genre)
		utils.Response(w, statusCode, contentType, content)
		global.GatewayLogger.Warn(errMsg)
	}
}

// entryFinal 进入最终处理器
func entryFinal(w http.ResponseWriter, req *http.Request, resp *http.Response, host string, entrance string, errInfo string, genre string) {
	if err := final.Entry(w, req, resp, host, entrance, errInfo); err != nil {
		errMsg := fmt.Sprintf("final execution failed. Error message: %s", err.Error())
		statusCode, contentType, content := utils.ResponseTemplate(constants.ResponseTitle.BadGateway, errMsg, genre)
		utils.Response(w, statusCode, contentType, content)
		global.GatewayLogger.Warn(errMsg)
	}
}

// isException 判断是否是异常来源
func isException(entrance string) bool {
	return entrance == constants.PostEntrance.Error
}

// isResponse 判断是否是响应来源
func isResponse(entrance string) bool {
	return entrance == constants.PostEntrance.Post
}

// isReject 判断是否是拒绝来源
func isReject(entrance string) bool {
	return entrance == constants.PostEntrance.Reject
}
