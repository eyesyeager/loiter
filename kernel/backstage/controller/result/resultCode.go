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

var Results = customResults{
	DefaultSuccess: customResult{200, "success"},
	DefaultFail:    customResult{400, "fail"},
	ServerError:    customResult{10000, "server error"},
	AuthError:      customResult{20001, "identity authentication failed"},
}
