package filter

import "net/http"

/**
 * 黑白名单
 * @auth eyesYeager
 * @date 2024/1/11 16:47
 */

func ReviewListFilter(w http.ResponseWriter, r *http.Request, host string) (error, bool) {
	return nil, true
}
