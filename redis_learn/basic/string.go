package basic

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func StringFuc(rdb *redis.Client) {
	ctx := context.Background()

	// 设置一个字符串类型的键值对
	rdb.Set(ctx,"myKey","Hello Redis!!!",0)

	// 获取键值
	val,err := rdb.Get(ctx,"myKey").Result()
	if err == redis.Nil || err != nil {
		log.Fatal(err)
	}
	log.Printf("Value at myKey: %s\n", val)
	
	// 判断键是否存在（Exists）
	exists, err := rdb.Exists(ctx, "myKey").Result()
	if err != nil {
    log.Fatal(err)
	}
	log.Printf("Key 'myKey' exists: %v\n", exists)

	// 删除键（Delete）
	deleted, err := rdb.Del(ctx, "myKey").Result()
	if err != nil {
    log.Fatal(err)
	}
	log.Printf("Deleted keys: %d\n", deleted)

	// 自增/自减数值（Increment/Decrement）
	value, err := rdb.Incr(ctx, "counter").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Increased counter: %d\n", value)

	decValue, err := rdb.Decr(ctx, "counter").Result()
	if err != nil {
			log.Fatal(err)
	}
	log.Printf("Decreased counter: %d\n", decValue)

	// 获取并设置（Get and Set）
	oldVal, err := rdb.GetSet(ctx, "myKey", "new value").Result()
	log.Printf("Old value was: %s, New value is now set to: %s\n", oldVal, "new value")
	if err != nil {
		log.Fatal(err)
	}
}