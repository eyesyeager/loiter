package genre

import "net/http"

/**
 * 通道类型
 * @auth eyesYeager
 * @date 2024/1/9 18:09
 */

// Aisle 通道方法类型
type Aisle func(http.ResponseWriter, *http.Request, string) (error, bool)
