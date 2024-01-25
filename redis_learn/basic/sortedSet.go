package basic

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func SortedSet(rdb *redis.Client) {
	ctx := context.Background()
	// 1.向有序集合添加成员和分数（ZAdd）
	rdb.ZAdd(ctx,"myZSet",&redis.Z{Score: 1.0, Member: "member1"})
	// 同时添加多个成员和分数
	members := []*redis.Z{
		{Score: 2.0, Member: "member2"},
		{Score: 3.0, Member: "member3"},
	}
	rdb.ZAdd(ctx, "myZSet", members...)

	// 2.获取有序集合中的所有成员及分数（ZRANGE / ZRANGEBYSCORE）
	// 获取整个有序集合（默认升序）
	membersWithScores, err2 := rdb.ZRangeWithScores(ctx, "myZSet", 0, -1).Result()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("membersWithScores:====",membersWithScores)
	// 根据分数范围获取成员
	rangeMembers, err := rdb.ZRangeByScoreWithScores(ctx, "myZSet", &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
		Offset: 0,
		Count: 10,
	}).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rangeMembers:====",rangeMembers)
	// 3.判断成员是否存在于有序集合中（ZScore）
	score, err3 := rdb.ZScore(ctx, "myZSet", "member1").Result()
	if err3 == redis.Nil {
		log.Println("member1 does not exist in myZSet")
	} else if err3 != nil {
		log.Fatal(err3)
	} else {
		log.Printf("Score of member1 in myZSet: %.2f\n", score)
	}
	// 4.从有序集合中删除一个或多个成员（ZRem）
	count, err4 := rdb.ZRem(ctx, "myZSet", "member1").Result()
	if err4 != nil {
		log.Fatal(err4)
	}
	log.Printf("Removed %d member(s) from myZSet\n", count)

	// 删除多个成员
	membersToRemove := []string{"member2", "member3"}
	interfaceMembers := make([]interface{}, len(membersToRemove))
	for i, member := range members {
		interfaceMembers[i] = member
	}
	rdb.ZRem(ctx, "myZSet", interfaceMembers...)
}