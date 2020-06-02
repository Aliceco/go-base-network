package main

import (
	"fmt"
	util "network/util"
	"io/ioutil"
	"log"
	"net/http"
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

	if request.URL.Path=="/a" {
		//log.Panicln(request.RemoteAddr)
		newRequest, _:=http.NewRequest(request.Method, "http://localhost:9091", request.Body)

		util.CloneHeader(request.Header, &newRequest.Header) // 拷贝请求头

		newRequest.Header.Add("x-forwarded-for", request.RemoteAddr)

		newResponse,_:=http.DefaultClient.Do(newRequest)

		getHeader:=writer.Header()
		fmt.Println(newResponse.Header)
		fmt.Println(getHeader)
		util.CloneHeader(newResponse.Header, &getHeader) // 拷贝响应头给客户端

		writer.WriteHeader(newResponse.StatusCode) // 写入http status

		defer newResponse.Body.Close()
		resCont, _:= ioutil.ReadAll(newResponse.Body)
		writer.Write(resCont) // 写入响应体给客户端
		return
	}
	writer.Write([]byte("default index"))
}

func main() {
	http.ListenAndServe(":8080", &ProxyHandler{})
}
