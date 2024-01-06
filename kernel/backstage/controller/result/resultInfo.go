package result

/**
 * 通用错误信息
 * @auth eyesYeager
 * @date 2024/1/4 16:49
 */

var ResultInfo = resultInfo{
	"数据库操作失败，error: %s",
}

type resultInfo struct {
	DbOperateError string
}
