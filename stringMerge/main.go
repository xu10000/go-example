package main

import "fmt"

func main() {

	orderByArr := []string{"name", "desc", "age", "asc"}
	Len := len(orderByArr) / 2
	str := ""
	for i := 0; i < Len; i++ {
		str = str + " " + orderByArr[2*i] + " " + orderByArr[2*i+1]
	}
	fmt.Println("------ print", str)
}
