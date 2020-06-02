package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CloneHeader(src http.Header, dest *http.Header)  {
	for k,v:=range src {
		dest.Set(k, v[0])
	}
}

func RequestUrl(writer http.ResponseWriter, request *http.Request, url string)  {
	newRequest, _:=http.NewRequest(request.Method, url, request.Body)

	CloneHeader(request.Header, &newRequest.Header) // 拷贝请求头

	newRequest.Header.Add("x-forwarded-for", request.RemoteAddr)

	newResponse, err:=http.DefaultClient.Do(newRequest)

	if newResponse==nil && err!=nil {
		fmt.Println("--------")
		fmt.Println(err)
		return
	}
	getHeader:=writer.Header()

	fmt.Println(newResponse)
	fmt.Println(getHeader)
	CloneHeader(newResponse.Header, &getHeader) // 拷贝响应头给客户端

	writer.WriteHeader(newResponse.StatusCode) // 写入http status

	defer newResponse.Body.Close()

	resCont, _:= ioutil.ReadAll(newResponse.Body)
	writer.Write(resCont) // 写入响应体给客户端
}