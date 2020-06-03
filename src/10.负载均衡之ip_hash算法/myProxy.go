package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
	util "network/util"
)

type ProxyHandler struct {}
func (* ProxyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err:=recover();err!=nil{
			writer.WriteHeader(500)
			log.Panicln(err)
		}
	}()

	lb:=util.NewLoadBalance()
	lb.AddServer(util.NewHttpServer("http://localhost:9091"))
	lb.AddServer(util.NewHttpServer("http://localhost:9092"))
	//url,_:=url2.Parse(lb.SelectByRand().Host)
	url,_:=url2.Parse(lb.SelectByIpHash(request.RemoteAddr).Host)
	proxy:=httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(writer, request)
}

func main() {
	http.ListenAndServe(":8080", &ProxyHandler{})
}
