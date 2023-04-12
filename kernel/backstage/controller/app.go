package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"zliway/kernel/backstage/controller/result"
)

/**
 * 注册应用相关接口
 * @author eyesYeager
 * @date 2023/4/11 17:55
 */

// RegisterApp 注册应用
func RegisterApp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result.SuccessDefault(w, nil)
}
