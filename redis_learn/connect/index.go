package connect

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var DB *redis.Client
func ConnectFunc() {
	// 创建一个 Redis 客户端实例
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 本地 Redis 的地址和端口
		Password: "redis",          // 如果有密码则填写
		DB:       0,                // 数据库索引号，默认是0
	})
	// 测试连接是否成功
	ctx := context.Background()
	fmt.Println(ctx,"ctx:======")
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong) // 输出：PONG
	DB = rdb
	// 使用完后关闭连接（通常在一个长生命周期的服务里，客户端会一直保持打开状态）
	defer rdb.Close()
}