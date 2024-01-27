package namelist

import (
	"errors"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"loiter/global"
	"loiter/kernel/backstage/controller/result"
)

/**
 * 白名单
 * @auth eyesYeager
 * @date 2024/1/23 18:04
 */

type WhiteNameList struct {
	host        string             // 主机地址
	bloomFilter *bloom.BloomFilter // 布隆过滤器
}

// Check 检查ip是否允许通行
func (l *WhiteNameList) Check(ip string) (error, bool) {
	// 布隆过滤器校验，没命中则一定不存在
	if !l.bloomFilter.Test([]byte(ip)) {
		return nil, false
	}
	// 布隆过滤器命中，但不一定存在，需要查库校验
	var res int
	if err := global.MDB.Raw(`SELECT 1 
					FROM app a, name_list nl 
					WHERE a.host = ? AND a.id = nl.app_id AND nl.genre = ? AND nl.ip = ?`, l.host, WhiteList, ip).Scan(&res).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, "WhiteList.Check()-Raw", err.Error())), false
	}
	return nil, res == 1
}

// Refresh 更新名单
func (l *WhiteNameList) Refresh() error {
	err, bloomFilter := buildBloomFilter(l.host, WhiteList)
	if err != nil {
		return err
	}
	l.bloomFilter = bloomFilter
	return nil
}

// NewWhiteNameList 创建黑名单结构体
func NewWhiteNameList(host string) (error, *WhiteNameList) {
	err, bloomFilter := buildBloomFilter(host, WhiteList)
	if err != nil {
		return err, nil
	}
	return nil, &WhiteNameList{
		host:        host,
		bloomFilter: bloomFilter,
	}
}
