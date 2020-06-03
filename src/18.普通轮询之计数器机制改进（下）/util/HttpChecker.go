package util

import (
	"net/http"
	"time"
)

type HttpChecker struct {
	Servers HttpServers // HttpServers LoadBalance里面的二开HttpServer
	FailMax int
	RecovCount int // 连续成功 到达这个值，标识为UP
}

func NewHttpChecker(servers HttpServers) *HttpChecker  {
	return &HttpChecker{servers, 6, 3}
}

func (this *HttpChecker) Check(timeout time.Duration)  {
	client:=http.Client{}
	for _,server:=range this.Servers{
		res,err:=client.Head(server.Host)
		if res!=nil {
			defer res.Body.Close()
		}
		if err!=nil {
			this.Fail(server)
			continue
		}
		if res.StatusCode>=200 && res.StatusCode<=400 {
			this.Success(server)
		} else {
			this.Fail(server)
		}
	}
}

func (this *HttpChecker) Fail(server *HttpServer)  {
	// 目前的机制是 计数器
	if server.FailCount>=this.FailMax { // 超过阈值
		server.Status="DOWN"
	} else {
		server.FailCount++
	}
	server.SuccessCount = 0
}
func (this *HttpChecker) Success(server *HttpServer)  {
	// 目前的机制是 计数器
	if server.FailCount>0 {
		server.FailCount--
		server.SuccessCount++
		if server.SuccessCount==this.RecovCount {
			server.FailCount=0
			server.Status="UP"
			server.SuccessCount=0
		}
	} else {
		server.Status="UP"
	}
}