package main

import (
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
		newRequest, _:=http.NewRequest(request.Method, "http://localhost:9091", request.Body)
		newResponse,_:=http.DefaultClient.Do(newRequest)
		defer newResponse.Body.Close()

		resCont, _:= ioutil.ReadAll(newResponse.Body)
		writer.Write(resCont)
		return
	}
	writer.Write([]byte("default index"))
}

func main() {
	http.ListenAndServe(":8080", &ProxyHandler{})
}
