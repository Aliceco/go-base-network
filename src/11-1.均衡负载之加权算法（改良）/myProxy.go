package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

	//url,_:=url2.Parse(lb.SelectByRand().Host)

	//url,_:=url.Parse(util.LB.SelectByIpHash(request.RemoteAddr).Host)
	url,_:=url.Parse(util.LB.SelectByWeightRand2().Host)
	proxy:=httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(writer, request)
}

func main() {
	http.ListenAndServe(":8080", &ProxyHandler{})
}
