package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("------ print", "test")
}
func main() {
	fmt.Println("------ print", 1)
	func() {
		time.Sleep(time.Second * 2)

		fmt.Println("------ print", 2)
	}()
	fmt.Println("------ print", 3)

}
