package returnee

/**
 * 获取请求日志详细信息
 * @auth eyesYeager
 * @date 2024/2/14 17:47
 */

type GetDetailedRequestNumLog struct {
	XAxis      []string `json:"xAxis"`
	Series     []int    `json:"series"`
	LastSeries []int    `json:"lastSeries"`
}
