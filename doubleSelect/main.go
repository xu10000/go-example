package main

import (
	"fmt"
	"sync"
	"time"
)

func loopTask(wg *sync.WaitGroup, stopChan chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			return
		default:
		}

		fmt.Println("------ print", "running")

		select {
		case <-stopChan:
			return
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	stopChan := make(chan struct{})

	wg := sync.WaitGroup{}
	defer wg.Wait()
	defer close(stopChan)

	wg.Add(1)
	go loopTask(&wg, stopChan)
	time.Sleep(time.Second * 3)
}
