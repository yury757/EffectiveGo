package main

import (
	"fmt"
	"context"
	"time"
	"github.com/go-redis/redis/v8"
)

type Person struct {
	name string
	age int
	address Address
}

type Address struct {
	province string
	city string
}

/*
func (person Person) toString() {
	fmt.Println(person)
	fmt.Printf("person的地址为 %p\n", &person)
}
*/

func (person *Person) toString() {
	fmt.Println(*person)
	fmt.Printf("person的地址为 %p\n", person)
}

func test2() {
	fmt.Println("123")
	var address = Address{"jiangxi", "fengcheng"}
	var person = Person{"yury", 27, address}
	var person2 = &Person{"yury", 12, address}
	person.toString()
	person2.toString()
}

func main() {
	var rdb *redis.Client = redis.NewClient(&redis.Options{
		Addr: "18.181.84.103:9221",
		Password: "1yuryamazon",
		DB: 0,
		PoolSize: 2,
	})
	var ctx, cancel = context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var result, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}else {
		fmt.Println(result)
	}
	rdb.Close()
}
