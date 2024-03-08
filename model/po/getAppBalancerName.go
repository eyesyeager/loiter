package po

/**
 * 容器—查询应用与负载均衡策略名
 * @auth eyesYeager
 * @date 2024/1/8 10:48
 */

type GetAppBalancerName struct {
	Host     string
	Balancer string
	Status   uint8
}
