package constants

/**
 * 黑白名单类型
 * @author eyesYeager
 * @date 2024/3/23 13:10
 */

var NameList = nameList{
	Black: "black",
	White: "white",
}

type nameList struct {
	Black string
	White string
}
