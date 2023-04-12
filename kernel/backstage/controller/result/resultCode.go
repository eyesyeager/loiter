package result

/**
 * @author eyesYeager
 * @date 2023/4/11 22:01
 */

type customResult struct {
	code int
	msg  string
}

type customResults struct {
	DefaultSuccess   customResult
	DefaultFail      customResult
	ValidateError    customResult
	AuthError        customResult
	ServerError      customResult
	AuthInsufficient customResult
}

// Results TODO: 记得修改状态码
var Results = customResults{
	DefaultSuccess:   customResult{200, "success"},
	DefaultFail:      customResult{400, "fail"},
	ServerError:      customResult{10000, "网关似乎出问题了"},
	ValidateError:    customResult{20000, "请求参数错误"},
	AuthError:        customResult{20001, "权限错误"},
	AuthInsufficient: customResult{20001, "权限不足"},
}