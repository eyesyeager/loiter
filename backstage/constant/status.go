package constant

/**
 * @author eyesYeager
 * @date 2023/9/28 9:14
 */

var Status = status{
	Normal: StatusStructure{
		Code: 1,
		Name: "正常",
	},
	Invalid: StatusStructure{
		Code: 2,
		Name: "失效",
	},
	Delete: StatusStructure{
		Code: 3,
		Name: "删除",
	},
}

type status struct {
	Normal  StatusStructure
	Invalid StatusStructure
	Delete  StatusStructure
}

func (s *status) GetNameByCode(code uint8) string {
	switch code {
	case s.Normal.Code:
		return s.Normal.Name
	case s.Invalid.Code:
		return s.Invalid.Name
	case s.Delete.Code:
		return s.Delete.Name
	default:
		return ""
	}
}

type StatusStructure struct {
	Code uint8
	Name string
}
