package helper

import (
	"encoding/json"
	"net/http"
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
	return err
}
