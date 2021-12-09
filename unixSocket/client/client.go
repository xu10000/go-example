package main

import (
	"fmt"
)

func main() {
	filename := "/tmp/us.socket"
	u := NewUnixSocket(filename)
	str := u.sendMsg("xzt")
	fmt.Println("------ print client receive msg ", str)
}
