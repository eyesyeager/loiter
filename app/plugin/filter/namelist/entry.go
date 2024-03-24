package namelist

import (
	"errors"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"loiter/backstage/controller/result"
	"loiter/config"
	"loiter/constants"
	"loiter/global"
)

/**
 * 黑白名单整合
 * @auth eyesYeager
 * @date 2024/1/23 18:00
 */

// CheckNameListGenre 校验名单类型是否合法
func CheckNameListGenre(genre string) bool {
	return genre == constants.NameList.Black || genre == constants.NameList.White
}

// INameList 黑白名单接口
type INameList interface {
	// Check 校验IP
	Check(ip string) (error, bool)
	// Refresh 刷新布隆过滤器
	Refresh() error
}

// NewNameList 创建黑白名单实例
func NewNameList(host string, genre string) (error, INameList) {
	if genre == constants.NameList.Black {
		return NewBlackNameList(host)
	} else if genre == constants.NameList.White {
		return NewWhiteNameList(host)
	} else {
		return errors.New(fmt.Sprintf("there is no nameList of type %s", genre)), nil
	}
}

// buildBloomFilter 构建布隆过滤器
func buildBloomFilter(host string, genre string) (error, *bloom.BloomFilter) {
	if !CheckNameListGenre(genre) {
		return errors.New(fmt.Sprintf("there is no nameList of type %s", genre)), nil
	}
	// 查询名单
	var nameList []string
	if err := global.MDB.Raw(`SELECT nl.ip 
						FROM app a, name_list nl 
						WHERE a.host = ? AND a.id = nl.app_id AND nl.genre = ?`, host, genre).Scan(&nameList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil
	}
	// 构建布隆过滤器
	bloomFilter := bloom.NewWithEstimates(config.Program.PluginConfig.NameListBloomCapacity, config.Program.PluginConfig.NameListBloomMisjudgmentRate)
	for _, item := range nameList {
		bloomFilter.Add([]byte(item))
	}
	return nil, bloomFilter
}
