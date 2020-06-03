package main

import (
	"fmt"
	"time"
)

func main() {
	// 每三秒执行一次
	t:=time.NewTicker(time.Second*3)
	for {
		select {
		case <- t.C:
			fmt.Println("aaaa")
		}
	}
}
