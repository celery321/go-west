package main

import (
	"fmt"
	"go-west/pkg/cache/redis"
)

func main() {



	conn, err := redis.Dial("tcp", "192.168.50.16:6379")
	if err != nil {
		fmt.Printf("err===%v\n", err)
	}
	fmt.Printf("con===%v\n", conn)
	defer conn.Close()
}
