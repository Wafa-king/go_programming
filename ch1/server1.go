// server1是一个迷你回声服务器
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	// err := http.ListenAndServe("127.0.0.1:8000", nil)
	// if err != nil {
	// 	panic(err)
	// }
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
