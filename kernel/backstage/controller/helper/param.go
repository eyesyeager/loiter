package helper

import (
	"encoding/json"
	"net/http"
	"zliway/global"
	"zliway/kernel/utils"
)

/**
 * 请求参数相关工具
 * @author eyesYeager
 * @date 2023/4/12 11:14
 */

// PostData 获取post请求参数
func PostData(r *http.Request, container any) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(container)
	if err == nil {
		result, _ := json.Marshal(container)
		global.Log.Info("ip:" + utils.GetIp(r) + " browser:" + utils.GetBrowser(r) + " result:" + string(result))
	}
	return err
}
