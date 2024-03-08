package returnee

/**
 * @auth eyesYeager
 * @date 2024/2/17 14:21
 */

type GetDetailedRequestVisitorLog struct {
	XAxis      []string `json:"xAxis"`
	Series     []int    `json:"series"`
	LastSeries []int    `json:"lastSeries"`
}
