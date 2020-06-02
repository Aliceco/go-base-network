package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	util "network/util"
	"regexp"
)

type ProxyHandler struct {}
func (* ProxyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err:=recover();err!=nil{
			writer.WriteHeader(500)
			log.Panicln(err)
		}
	}()
	for k,v:=range util.ProxyConfigs {
		if matched, _:=regexp.MatchString(k, request.URL.Path);matched==true {
			//util.RequestUrl(writer, request, v) // 开始反向代理（自己用http实现）

			// 直接使用go http封装的好api
			target, _ := url.Parse(v) // 目标网址
			proxy:=httputil.NewSingleHostReverseProxy(target)
			//proxy.Transport // 自定义网络超时等等。。。
			proxy.ServeHTTP(writer, request)
			return
		}
	}
	writer.Write([]byte("default index"))
}

func main() {
	http.ListenAndServe(":8080", &ProxyHandler{})
}
