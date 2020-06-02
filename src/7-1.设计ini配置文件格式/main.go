package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

type webHandler struct {
}
func(webHandler) GetIP(request *http.Request) string{
	ips:=request.Header.Get("x-forwarded-for")
	if ips!=""{
		ipsList:= strings.Split(ips, ",")
		if len(ipsList)>0 && ipsList[0]!=""{
			return ipsList[0]
		}
	}
	return request.RemoteAddr
}
func (this webHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	auth:=request.Header.Get("Authorization")
	if auth=="" {
		writer.Header().Set("WWW-Authenticate", `Basic realm="您必须输入用户名和密码"`)
		writer.WriteHeader(401)
		return
	}
	// Authorization: Basic xxx
	authList:=strings.Split(auth, " ")
	fmt.Println(authList)
	if len(authList)==2 && authList[0]=="Basic" {
		res, err:=base64.StdEncoding.DecodeString(authList[1])
		fmt.Println(string(res))
		if err==nil && string(res)=="an:123" {
			writer.Write([]byte(fmt.Sprintf("<h1>web1, 来自于:%s</h1>", this.GetIP(request))))
			//writer.Write([]byte("<h1>web1</h1>"))
			return
		}
	}
	writer.Write([]byte("用户名密码错误"))
}

type web2Handler struct {
}
func (web2Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	writer.Write([]byte("web2"))
}

func main() {
	// 1.监听信号
	// 2.channel

	c:=make(chan os.Signal)

	go (func() {
		http.ListenAndServe(":9091", webHandler{})
	})()
	go (func() {
		http.ListenAndServe(":9092", web2Handler{})
	})()

	signal.Notify(c, os.Interrupt)
	s:=<-c // 阻塞（除非监听到信号值）
	log.Panicln(s)
}
