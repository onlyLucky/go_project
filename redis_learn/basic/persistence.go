package basic

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

func RDB2AOF(rdb *redis.Client) {
	ctx := context.Background()

	// 获取redis 的 RDB 最后一次保存时间
	lastSave,err := rdb.LastSave(ctx).Result()
	if err != nil {
		panic(err)
	}
	lastSaveTimeStr := time.Unix(lastSave,0).Format("2006-01-02 15:04:05")
	fmt.Printf("Last RDB save time: %v\n", lastSaveTimeStr)

	// 获取 AOF 相关状态信息
	aofInfo,errAOF := rdb.Info(ctx, "persistence").Result()
	if errAOF != nil {
		panic(errAOF)
	}
	aofInfoLines := strings.Split(aofInfo, "\n") // 将信息按行分割

	for _,line := range aofInfoLines {
		parts := strings.SplitN(line,":",2)// 每行信息按照冒号分割成键和值两部分
		if len(parts) < 2 {
			continue // 忽略无效行（例如空行）
		}
		key := parts[0]
		if strings.HasPrefix(key,"aof_"){
			fmt.Printf("%s == %s\n",key,parts[1])
		}
	}
}