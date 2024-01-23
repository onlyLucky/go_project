package basic

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func SetFunc(rdb *redis.Client) {
	ctx := context.Background()
	
	fmt.Println("SetFunc:=======")
	// 1.向集合添加成员（SAdd）
	err := rdb.SAdd(ctx, "mySet", "member1").Err()
	if err != nil {
		log.Fatal(err)
	}
	// 或者同时添加多个成员
	members := []string{"member2","member3"}
	//  []string{} => []interface{}
	interfaceMembers := make([]interface{}, len(members))
	for i, member := range members {
		interfaceMembers[i] = member
	}
	_,err1 := rdb.SAdd(ctx,"mySet",interfaceMembers...).Result()
	if err1 != nil {
		log.Fatal(err1)
	}

	// 2.获取集合中的所有成员（SMembers）
	members2, err2 := rdb.SMembers(ctx,"mySet").Result()
	if err2 != nil{
		log.Fatal(err2)
	}
	fmt.Println(members2)
	/* for _,member2 :=range members2 {
		log.Printf("Member in mySet: %s\n", member2)
	} */

	// 3.判断成员是否存在集合中（SIsMember）
	exists, err3 := rdb.SIsMember(ctx, "mySet", "member1").Result()
	if err3 != nil {
		log.Fatal(err3)
	}
	if exists {
		log.Println("member1 is a member of mySet")
	} else {
		log.Println("member1 is not a member of mySet")
	}

	// 4.从集合中移除一个或多个成员（SRem）
	rdb.SRem(ctx, "mySet", "member1")
	// 或者同时移除多个成员
	membersToRemove := []string{"member2", "member3"}
	interfaceMembersToRemove := make([]interface{}, len(membersToRemove))
	for i, member := range membersToRemove {
		interfaceMembersToRemove[i] = member
	}
	rdb.SRem(ctx, "mySet", interfaceMembersToRemove...)
	// 5.集合的交集、并集和差集操作
	// 计算两个集合的交集
	intersection, err := rdb.SInter(ctx, "mySet", "anotherSet").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(intersection)
	// 计算两个集合的并集
	union, err := rdb.SUnion(ctx, "mySet", "anotherSet").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(union)
	// 计算两个集合的差集（第一个集合与第二个集合之间的差异）
	difference, err := rdb.SDiff(ctx, "mySet", "anotherSet").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(difference)
}