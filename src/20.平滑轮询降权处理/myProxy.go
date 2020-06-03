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

	if request.URL.Path=="/favicon.ico" { // 谷歌会访问一个图标文件，我们不做处理
		return
	}
	url,_:=url.Parse(util.LB.RoundRobinByWeight3().Host)
	//fmt.Println(url)
	proxy:=httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(writer, request)
}

func main() {
	http.ListenAndServe(":8080", &ProxyHandler{})
}
