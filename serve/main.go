package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":9090", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "请求参数是：", r.URL.RawQuery)
}
