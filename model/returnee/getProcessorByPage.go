package returnee

import "loiter/model/structure"

/**
 * @auth eyesYeager
 * @date 2024/3/1 15:29
 */

type GetProcessorByPage struct {
	structure.PageStruct                           // 分页参数
	Total                int64                     `json:"total"` // 总数
	Data                 []GetProcessorByPageInner `json:"data"`  // 数据
}

type GetProcessorByPageInner struct {
	AppId         uint     `json:"appId"`
	AppName       string   `json:"appName"`
	FilterStr     string   `json:"-"`
	AidStr        string   `json:"-"`
	ExceptionStr  string   `json:"-"`
	FinalStr      string   `json:"-"`
	Filter        []string `json:"filter"`
	Aid           []string `json:"aid"`
	Exception     []string `json:"exception"`
	Final         []string `json:"final"`
	FilterCode    []string `json:"filterCode"`
	AidCode       []string `json:"aidCode"`
	ExceptionCode []string `json:"exceptionCode"`
	FinalCode     []string `json:"finalCode"`
}
