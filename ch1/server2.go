// server2是一个迷你的回声和计数服务器
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/counter", Counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Handler 处理回显请求的URL路径部分
func Handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.PATH= %q\n", r.URL.Path)
}

// Counter 回显目前为止调用的次数
func Counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Counter: %d\n", count)
	mu.Unlock()
}
