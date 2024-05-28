package returnee

/**
 * @auth eyesYeager
 * @date 2024/2/17 13:47
 */

type GetDetailedRequestQPSLog struct {
	XAxis      []string `json:"xAxis"`
	Series     []int    `json:"series"`
	LastSeries []int    `json:"lastSeries"`
}
