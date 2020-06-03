package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// 关于crc算法（循环冗余校验）和术语算法实现，暂时忽略
	ip:="127.0.0.1"
	fmt.Println(crc32.ChecksumIEEE([]byte(ip))) // IEEE 多项式返回数据的 CRC-32 校验和
	fmt.Println(0%3) // 0
	fmt.Println(1%3) // 1
	fmt.Println(2%3) // 2
	fmt.Println(3%3) // 0
}
