package returnee

/**
 * @auth eyesYeager
 * @date 2024/2/16 15:00
 */

type GetDetailedRequestExtremumLog struct {
	RequestNum    int64 `json:"requestNum"`    // 总请求数
	RunTimeMin    int64 `json:"runTimeMin"`    // 最小响应时间
	RunTimeMax    int64 `json:"runTimeMax"`    // 最大响应时间
	RunTimeAvg    int64 `json:"runTimeAvg"`    // 平均响应时间
	QPSAvg        int64 `json:"qpsAvg"`        // 平均QPS
	RequestReject int64 `json:"requestReject"` // 请求拒绝数
}
