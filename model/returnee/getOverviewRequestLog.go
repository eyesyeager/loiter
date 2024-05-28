package returnee

/**
 * 获取请求日志概览
 * @auth eyesYeager
 * @date 2024/2/14 14:53
 */

type GetOverviewRequestLog struct {
	RequestNum uint64 `json:"requestNum"` // 请求总数
	AvgRunTime uint64 `json:"avgRunTime"` // 请求平均耗时(ms)
	VisitorNum uint64 `json:"visitorNum"` // 访客总数
	RejectNum  uint64 `json:"rejectNum"`  // 请求拒绝总数
}
