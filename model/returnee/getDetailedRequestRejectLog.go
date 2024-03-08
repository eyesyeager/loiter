package returnee

/**
 * @auth eyesYeager
 * @date 2024/2/17 16:45
 */

type GetDetailedRequestRejectLog struct {
	XAxis        []string `json:"xAxis"`
	RejectSeries []int    `json:"rejectSeries"`
	ErrorSeries  []int    `json:"errorSeries"`
}
