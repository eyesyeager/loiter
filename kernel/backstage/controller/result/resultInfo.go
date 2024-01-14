package result

/**
 * 通用错误信息
 * @auth eyesYeager
 * @date 2024/1/4 16:49
 */

var CommonInfo = commonInfo{
	"数据库操作失败，失败位置：%s，错误信息: %s",
}

type commonInfo struct {
	DbOperateError string
}
