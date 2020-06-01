package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str:="an:123"
	base:=base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(base)
	decodeBytes,_:=base64.StdEncoding.DecodeString(base)
	fmt.Println(string(decodeBytes))
}
