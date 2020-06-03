package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机数
	rand.Seed(time.Now().UnixNano())
	// 0 - 1  随机
	fmt.Println(rand.Intn(2))
}
