package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// WaitForServer 尝试连接URL对应用服务器，
// 在一分钟内使用指数退避策略进行重试
// 所有的尝试失败后返回错误
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			// fmt.Fprintf(os.Stderr, "Site is down:%v\n", err)
			log.Fatalf("Site is down :%v\n", err)
			return nil
		}
		log.Printf("server not responding: %s\n", err)
		sleep := time.Second << uint(tries)
		fmt.Println(sleep)
		time.Sleep(sleep)
	}
	return fmt.Errorf("server %s failed after %v", url, timeout)
}
