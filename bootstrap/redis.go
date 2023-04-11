package bootstrap

import (
	"fmt"
	"github.com/go-redis/redis"
	"zliway/global"
)

/**
 * redis启动文件
 * @author eyesYeager
 * @date 2023/4/10 13:42
 */

// initializeRDB 初始化Redis
func initializeRDB() {
	fmt.Println("start connecting to Redis server...")

	global.RDB = redis.NewClient(&redis.Options{
		Addr:     global.Config.Persistent.Redis.Host + ":" + global.Config.Persistent.Redis.Port,
		Password: global.Config.Persistent.Redis.Password,
		DB:       global.Config.Persistent.Redis.Db,
	})

	_, err := global.RDB.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("failed to start Redis: %s", err))
	}

	fmt.Println("successfully connected to Redis server")
}
