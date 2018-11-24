// fetch 输出从URL获取的内容
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ReadAll: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		fmt.Println(resp.Header)
	}
	fmt.Printf("total time: %.2fs\n", time.Since(start).Seconds())
}
