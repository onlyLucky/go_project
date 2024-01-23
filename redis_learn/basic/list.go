package basic

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func ListFunc(rdb *redis.Client) {
	fmt.Println("ListFunc:========")
	ctx := context.Background()
	// 1.向列表头部添加元素（LPush）
	err := rdb.RPush(ctx,"myList","item1").Err()
	if err != nil {
    log.Fatal(err)
	}
	// 或者同时添加多个元素
	rdb.RPush(ctx,"myList","item2","item3")

	// 2. 向列表尾部添加元素（RPush）
	err1 := rdb.LPush(ctx, "myList", "item4R").Err()
	if err1 != nil {
		log.Fatal(err1)
	}
	// 或者同时添加多个元素
	rdb.LPush(ctx, "myListR", "item5R", "item6R")

	// 3.从列表头部弹出元素（LPop）
	rdb.LPop(ctx, "myList")

	// 4.从列表尾部弹出元素（RPop）
	rdb.RPop(ctx, "myList")

	// 5.获取列表范围内的元素（LRANGE）
	values, err5 := rdb.LRange(ctx,"myList",0,-1).Result()
	if err5 != nil {
		log.Fatal(err5)
	}
	fmt.Println(values)
	/* for _, v := range values {
		log.Printf("Value in myList: %s\n", v)
	} */

	// 6.列表长度（LLEN）
	length, err6 := rdb.LLen(ctx, "myList").Result()
	if err6 != nil {
			log.Fatal(err6)
	}
	log.Printf("Length of myList: %d\n", length)

	// 7.移除列表中的一个或多个值（LREM）
	rdb.LRem(ctx, "myList",0,"item1")
}