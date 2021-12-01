package main

import (
	"fmt"
	"time"
)

type funcWithChan func(chan struct{})

func startWithChannel(stopChan chan struct{}, f funcWithChan) {
	go f(stopChan)
}

func task(stopChan chan struct{}) {
	fmt.Println("------ start task")
	<-stopChan
	fmt.Println("------ start end")

}

// 主进程下发结束信息
func main() {
	stopChan := make(chan struct{})
	startWithChannel(stopChan, task)
	// 结束
	stopChan <- struct{}{}
	time.Sleep(500 * time.Millisecond)
}
