package container

/**
 * 应用与应用实例容器
 * @auth eyesYeager
 * @date 2024/1/5 15:11
 */

var ServerByAppMap map[string][]ServerWeight

type ServerWeight struct {
	Server string
	Weight uint
}
