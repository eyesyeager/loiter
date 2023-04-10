package test

/**
 * 测试入口
 * @author eyesYeager
 * @date 2023/4/9 20:45
 */

func Web() {
	go StartWebA()
	go StartWebB()
}
