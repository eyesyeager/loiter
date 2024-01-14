package pipeline

import "loiter/kernel/passageway/genre"

/**
 * 管道集合
 * @auth eyesYeager
 * @date 2024/1/9 18:07
 */

// InitPipeline 初始化管道配置
func InitPipeline() map[string]genre.Aisle {
	pipelineMap := make(map[string]genre.Aisle)
	pipelineMap["requestLog"] = RequestLogPipeline
	return pipelineMap
}
