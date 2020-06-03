package util

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"time"
)

type HttpServer struct { // 目标server类
	Host string
	Weight int
}
func NewHttpServer(host string, weight int) *HttpServer  {
	return &HttpServer{Host:host, Weight: weight}
}

type LoadBalance struct { // 负载均衡类
	Servers []*HttpServer
}
func NewLoadBalance() *LoadBalance {
	return &LoadBalance{Servers:make([]*HttpServer, 0)}
}

func (this *LoadBalance) AddServer(server *HttpServer)  {
	this.Servers=append(this.Servers, server)
}

func (this *LoadBalance) SelectByRand() *HttpServer { // 随机算法
	rand.Seed(time.Now().UnixNano())
	// 0 - 1  随机
	index:=rand.Intn(len(this.Servers))
	return this.Servers[index]
}

func (this *LoadBalance) SelectByIpHash(ip string) *HttpServer { // ip_hash算法
	index:= int(crc32.ChecksumIEEE([]byte(ip))) % len(this.Servers)
	return this.Servers[index]
}

func (this *LoadBalance) SelectByWeightRand() *HttpServer { // 加权随机算法
	rand.Seed(time.Now().UnixNano())
	index:=rand.Intn(len(ServerIndices))
	return this.Servers[ServerIndices[index]]
}
func (this *LoadBalance) SelectByWeightRand2() *HttpServer { // 加权随机算法（改良算法）
	rand.Seed(time.Now().UnixNano())
	sumList:=make([]int, len(this.Servers))
	sum:=0
	for i:=0;i<len(this.Servers);i++ {
		sum+=this.Servers[i].Weight
		sumList[i]=sum
	}
	randInt:=rand.Intn(sum) //[) 左闭右开
	for index, value:=range sumList {
		if randInt<value {
			return this.Servers[index]
		}
	}
	return this.Servers[0]
}
var LB *LoadBalance
var ServerIndices []int
func init()  {
	LB=NewLoadBalance()
	LB.AddServer(NewHttpServer("http://localhost:9091", 5))
	LB.AddServer(NewHttpServer("http://localhost:9092", 15))
	for index,server:=range LB.Servers{
		if server.Weight > 0 {
			for i:=0;i<server.Weight;i++ {
				ServerIndices=append(ServerIndices, index)
			}
		}
	}
	fmt.Println(ServerIndices)
}