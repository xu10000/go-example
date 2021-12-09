package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

var ProtoSize = 20000

type UnixScoket interface {
	sendMsg(string) string
}

type UnixScoketImpl struct {
	filename string
	bufsize  int
}

func getSendByte(str string, i int) []byte {
	b := []byte(str + strconv.Itoa(i))
	bSize := len(b)
	appendByte := make([]byte, ProtoSize-bSize)
	return append(b, appendByte...)
}

func (u *UnixScoketImpl) sendMsg(str string) string {
	addr, err := net.ResolveUnixAddr("unix", u.filename)
	if err != nil {
		panic("Cannot resolve unix addr: " + err.Error())
	}
	c, err := net.DialUnix("unix", nil, addr)
	// defer c.Close()

	if err != nil {
		panic("DialUnix failed.")
	}
	// var buf []byte
	// var nr int
	for i := 0; i < 5; i++ {
		//写出
		_str := getSendByte(str, i)
		fmt.Println("------ printxx", len(_str))
		_, err = c.Write(_str)
		if err != nil {
			panic("Writes failed.")
		}
		fmt.Println("------ client write str+i ok", str+strconv.Itoa(i))

		// time.Sleep(time.Second * 1)
		//读结果
		buf := make([]byte, u.bufsize)
		_, err := io.ReadFull(c, buf)

		if err != nil {
			panic("Read: " + err.Error())
		}
		// time.Sleep(time.Second)
		fmt.Println("------ client read success ", string(buf))
	}
	return "s"
}

func NewUnixSocket(filename string) UnixScoket {
	return &UnixScoketImpl{
		filename,
		ProtoSize,
	}
}
