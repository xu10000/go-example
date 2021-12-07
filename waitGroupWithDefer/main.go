package main

import (
	"fmt"
	"sync"
	"time"
)

// wg chan defer 结合，优雅的实现自上而下通道关闭
func main() {
	wg := sync.WaitGroup{}
	stopChan := make(chan struct{})
	Len := 2
	wg.Add(Len)
	for i := 0; i < Len; i++ {
		go startWithChannel(&wg, stopChan)
	}
	defer wg.Wait() // 等待子routine先关闭
	defer close(stopChan)

	// 一秒后结束主进程，优雅关闭子routine
	time.Sleep(time.Second)

}

func startWithChannel(wg *sync.WaitGroup, stopChan chan struct{}) {
	fmt.Println("------ print", "task start")
	<-stopChan
	fmt.Println("------ print", "task stop")
	wg.Done()
}
