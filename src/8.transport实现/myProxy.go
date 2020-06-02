package main

import (
	"log"
	"net/http"
	util "network/util"
	"regexp"
)

type ProxyHandler struct {}
func (* ProxyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//http://localhost:8080
	//fmt.Println(request.RequestURI)
	//fmt.Println(request.URL.Path)
	//writer.Write([]byte("prxoy"))

	defer func() {
		if err:=recover();err!=nil{
			writer.WriteHeader(500)
			log.Panicln(err)
		}
	}()

	//if request.URL.Path=="/a" {
	//	//log.Panicln(request.RemoteAddr)
	//	util.RequestUrl(writer, request, "http://localhost:9091")
	//	return
	//} else if request.URL.Path=="/b" {
	//	util.RequestUrl(writer, request, "http://localhost:9092")
	//	return
	//}

	for k,v:=range util.ProxyConfigs {
		if matched, _:=regexp.MatchString(k, request.URL.Path);matched==true {
			util.RequestUrl(writer, request, v)
		}
	}

	writer.Write([]byte("default index"))
}

func main() {
	http.ListenAndServe(":8080", &ProxyHandler{})
}
