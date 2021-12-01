package main

import (
	"fmt"
	"sync"
)

func task(wg sync.WaitGroup, stopChan chan struct{}) {
	// 异常
	stopChan <- struct{}{}
	// 正常
	// fmt.Println("------ print ok")
	// wg.Done()
}

// 主线程监听goroutine失败
func main() {
	wg := sync.WaitGroup{}
	stopChan := make(chan struct{})
	wgChan := make(chan struct{})
	wg.Add(1)
	go task(wg, stopChan)

	go func() {
		wg.Wait()
		close(wgChan)
	}()

	select {
	case <-wgChan:
		fmt.Println("------ task success done")
	case <-stopChan:
		fmt.Println("------ task cause error")
	}
}
