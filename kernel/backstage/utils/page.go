package utils

import (
	"errors"
	"loiter/kernel/model/structure"
	"strconv"
)

/**
 * 分页器工具类
 * @auth eyesYeager
 * @date 2024/1/4 15:30
 */

// 但页最大数量
var maxPageSize = 300

// CheckPageStruct 校验PageStruct结构
func CheckPageStruct(pageStruct structure.PageStruct) error {
	if pageStruct.PageNo <= 0 {
		return errors.New("非法参数值 'pageNo': " + strconv.Itoa(pageStruct.PageNo))
	}
	if pageStruct.PageSize < 0 {
		return errors.New("非法参数值 'pageSize': " + strconv.Itoa(pageStruct.PageSize))
	}
	if pageStruct.PageSize > maxPageSize {
		return errors.New("非法参数 'pageSize', 最大值为" + strconv.Itoa(maxPageSize))
	}
	return nil
}

// BuildPageSearch 构建分页查询条件
func BuildPageSearch(pageStruct structure.PageStruct) (int, int) {
	return pageStruct.PageSize, pageStruct.PageSize * (pageStruct.PageNo - 1)
}
