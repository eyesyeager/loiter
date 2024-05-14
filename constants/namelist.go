package constants

/**
 * 黑白名单类型
 * @author eyesYeager
 * @date 2024/3/23 13:10
 */

var NameList = nameList{
	Black: NameListStruct{
		Label: "黑名单",
		Value: "black",
	},
	White: NameListStruct{
		Label: "白名单",
		Value: "white",
	},
}

type nameList struct {
	Black NameListStruct
	White NameListStruct
}

type NameListStruct struct {
	Label string
	Value string
}
