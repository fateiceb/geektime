package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var (
	datatype = flag.String("type", "wrong", "input data type")
	rdb      = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx = context.Background()
)

func main() {
	flag.Parse()

	fmt.Println(*datatype)
	switch *datatype {
	case "string":
		String()
	case "list":
		List()
	case "hash":
		Hash()
	case "Set":
		Set()
	case "zset":
		ZSET()
	case "hyperloglog":
		HyperLogLog()
	default:
		fmt.Println("wrong type")
	}
}

//redis字符串操作
func String() {
	err := rdb.Set(ctx, "sds", "1", 0).Err()
	if err != nil {
		panic(err)
	}
	res, err := rdb.Get(ctx, "sds").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

//redis列表操作
func List() {
	rdb.LPush(ctx, "list", 1)
	rdb.RPush(ctx, "list", 2)
	res, err := rdb.LRange(ctx, "list", 0, 10).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("list", res)
}

//redis 哈希表操作
func Hash() {
	err := rdb.HSet(ctx, "h1ash", "KE11111111", "111").Err()
	if err != nil {
		panic(err)
	}
	res, err := rdb.HGet(ctx, "h1ash", "KE11111111").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	_, err = rdb.HMSet(ctx, "aaa", "1key", "1val", "2key", "2val").Result()
	if err != nil {
		panic(err)
	}
	m, err := rdb.HGetAll(ctx, "aaa").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", m)

}

//redis 集合操作
func Set() {
	err := rdb.SAdd(ctx, "set", "monkey", "monkey2", "monkey").Err()
	if err != nil {
		panic(err)
	}
	res, err := rdb.SMembers(ctx, "set").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

//redis 有序集合操作
func ZSET() {
	err := rdb.ZAdd(ctx, "1et", &redis.Z{11.0, "mysql"}, &redis.Z{2.0, "12"}, &redis.Z{2.0, "123"}).Err()
	if err != nil {
		panic(err)
	}
	res, err := rdb.ZRange(ctx, "1et", 0, 20).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
func HyperLogLog() {
	err := rdb.PFAdd(ctx, "hyper", 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 6).Err()
	if err != nil {
		panic(err)
	}
	cnt, err := rdb.PFCount(ctx, "hyper").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(cnt)
}
