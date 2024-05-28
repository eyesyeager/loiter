package constants

/**
 * 处理器类型常量
 * @auth eyesYeager
 * @date 2024/2/12 18:47
 */

var Processor = processor{
	Filter: ProcessorStructure{
		Code: "filter",
		Name: "过滤器",
	},
	Aid: ProcessorStructure{
		Code: "aid",
		Name: "响应处理器",
	},
	Exception: ProcessorStructure{
		Code: "exception",
		Name: "异常处理器",
	},
	Final: ProcessorStructure{
		Code: "final",
		Name: "最终处理器",
	},
}

type processor struct {
	Filter    ProcessorStructure
	Aid       ProcessorStructure
	Exception ProcessorStructure
	Final     ProcessorStructure
}

type ProcessorStructure struct {
	Code string
	Name string
}
