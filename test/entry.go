package test

/**
 * 测试服务器
 * @author eyesYeager
 * @date 2023/4/10 11:24
 */

func Start() {
	go StartWebA()
	go StartWebB()
	go StartWebC()
}
