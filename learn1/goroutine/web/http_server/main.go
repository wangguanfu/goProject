package main

import (
	"fmt"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "%v\n", r.Form)
	fmt.Fprintf(w, "path:%s\n", r.URL.Path)
	fmt.Fprintf(w, "schema:%s\n", r.URL.Scheme)
	fmt.Fprintf(w, "hello world\n")
}

func main() {
	http.HandleFunc("/", sayhello) // 设置访问路由
	err := http.ListenAndServe(":9000", nil) //设置监听端口
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

}
