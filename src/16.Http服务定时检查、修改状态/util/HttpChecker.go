package util

import (
	"net/http"
	"time"
)

type HttpChecker struct {
	Servers HttpServers // HttpServers LoadBalance里面的二开HttpServer
}

func NewHttpChecker(servers HttpServers) *HttpChecker  {
	return &HttpChecker{servers}
}

func (this *HttpChecker) Check(timeout time.Duration)  {
	client:=http.Client{}
	for _,server:=range this.Servers{
		res,err:=client.Head(server.Host)
		if res!=nil {
			defer res.Body.Close()
		}
		if err!=nil {
			server.Status="DOWN"
			continue
		}
		if res.StatusCode>=200 && res.StatusCode<=400 {
			server.Status="UP"
		} else {
			server.Status="DOWN"
		}
	}
}