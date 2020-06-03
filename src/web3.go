package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("web3"))
	})
	http.ListenAndServe(":9093", nil)
}