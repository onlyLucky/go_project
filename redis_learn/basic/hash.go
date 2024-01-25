package basic

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func HashFunc(rdb *redis.Client) {
	// 1.向哈希添加字段和值（HSet）
	ctx := context.Background()
	rdb.HSet(ctx,"myHash","field1","value1")
	// 或者同时添加多个字段和值
	fieldsAndValues := map[string]interface{}{
    "field2": "value2",
    "field3": "value3",
	}
	rdb.HMSet(ctx, "myHash", fieldsAndValues)

	// 2.获取整个哈希的所有字段和值（HGetAll）
	hashMap, err2 := rdb.HGetAll(ctx, "myHash").Result()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("hashMap: =====",hashMap)
	// 3.获取哈希中特定字段的值（HGet）
	val, err3 := rdb.HGet(ctx, "myHash","field1").Result()
	if err3 == redis.Nil {
		log.Println("field1 does not exist in myHash")
	} else if err3 != nil {
		log.Fatal(err3)
	} else {
		log.Printf("Value of field1 in myHash: %s\n", val)
	}
	// 4.判断哈希中是否存在某个字段（HEXISTS）
	exists, err4 := rdb.HExists(ctx, "myHash", "field1").Result()
	if err4 != nil {
		log.Fatal(err4)
	}
	if exists {
		log.Println("field1 exists in myHash")
	} else {
		log.Println("field1 does not exist in myHash")
	}
	// 5.从哈希中删除一个或多个字段（HDel）
	rdb.HDel(ctx, "myHash", "field1")
	// 删除多个字段
	fieldsToRemove := []string{"field2", "field3"}
	rdb.HDel(ctx, "myHash", fieldsToRemove...)
}