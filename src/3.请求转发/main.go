package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

type webHandler struct {
}
func (webHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	writer.Write([]byte("<h1>web1</h1>"))
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
