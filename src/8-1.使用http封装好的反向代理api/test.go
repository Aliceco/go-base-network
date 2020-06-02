package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	cfg, err := ini.Load("env")
	if err!=nil{
		log.Println(err)
		return
	}
	sec,err:=cfg.GetSection("proxy")
	//log.Println(sec.GetKey("path"))
	//log.Println(sec.GetKey("aaa")) // 不存在
	secs:=sec.ChildSections()
	for _,sec:=range secs{
		fmt.Println(sec.Name())
		//proxy.a
		//proxy.b
	}
}
