package foundation

import (
	"errors"
	"loiter/global"
	"loiter/model/entity"
)

/**
 * 角色服务
 * @author eyesYeager
 * @date 2023/11/26 18:20
 */

type roleFoundation struct {
	WeightByRoleMap map[string]uint // map[role]weight 结构
	RidByRoleMap    map[string]uint // map[role]rid 结构
}

var RoleFoundation = roleFoundation{}

// InitRoleContainer 初始化角色容器
func (roleFoundation *roleFoundation) InitRoleContainer() error {
	// 获取角色信息
	var roleList []entity.Role
	tx := global.MDB.Find(&roleList)
	// 获取角色信息异常
	if tx.Error != nil {
		return errors.New("role information query failed, error:" + tx.Error.Error())
	}
	// role表没有角色数据
	if tx.RowsAffected == 0 {
		return errors.New("there is no role data in the role table. Please initialize the database before starting it")
	}

	// 初始化各类容器
	roleFoundation.initWeightByRoleMap(roleList)
	roleFoundation.initRidByRoleMap(roleList)
	return nil
}

// InitWeightByRoleMap 初始化 weightByRoleMap
func (roleFoundation *roleFoundation) initWeightByRoleMap(roleList []entity.Role) {
	// 初始化/清空 weightByRoleMap
	roleFoundation.WeightByRoleMap = make(map[string]uint)
	// 填充 weightByRoleMap
	for _, roleEntity := range roleList {
		roleFoundation.WeightByRoleMap[roleEntity.Name] = roleEntity.Weight
	}
}

// initRidByRoleMap 初始化 ridByRoleMap
func (roleFoundation *roleFoundation) initRidByRoleMap(roleList []entity.Role) {
	// 初始化/清空 ridByRoleMap
	roleFoundation.RidByRoleMap = make(map[string]uint)
	// 填充 ridByRoleMap
	for _, roleEntity := range roleList {
		roleFoundation.RidByRoleMap[roleEntity.Name] = roleEntity.ID
	}
}

// GetWeightByRole 获取角色对应的权重
func (roleFoundation *roleFoundation) GetWeightByRole(role string) (error, uint) {
	weight := roleFoundation.WeightByRoleMap[role]
	if weight == 0 {
		return errors.New("role " + role + " not defined"), 0
	}
	return nil, weight
}

// GetRidByRole 获取角色对应的id
func (roleFoundation *roleFoundation) GetRidByRole(role string) (error, uint) {
	rid := roleFoundation.RidByRoleMap[role]
	if rid == 0 {
		return errors.New("role " + role + " not defined"), 0
	}
	return nil, rid
}

// CompareRole 比较角色大小
// int 大于 0 则 role1 大于 role2
// int 等于 0 则 role1 等于 role2
// int 小于 0 则 role1 小于 role2
func (roleFoundation *roleFoundation) CompareRole(role1 string, role2 string) (error, int) {
	weight1 := roleFoundation.WeightByRoleMap[role1]
	if weight1 == 0 {
		return errors.New("role " + role1 + " not defined"), 0
	}
	weight2 := roleFoundation.WeightByRoleMap[role2]
	if weight2 == 0 {
		return errors.New("role " + role2 + " not defined"), 0
	}
	return nil, int(weight1 - weight2)
}
