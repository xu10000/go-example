package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var ProtoSize = 20000

type UnixScoket interface {
	handle(net.Conn)
	startServer()
}

type UnixScoketImpl struct {
	filename string
	bufsize  int
	errCh    chan struct{}
}

func (u *UnixScoketImpl) handle(conn net.Conn) {

	defer conn.Close()
	go func() chan struct{} {

		// time.Sleep(time.Second * 5)
		for {

			buf := make([]byte, u.bufsize)
			_, err := io.ReadFull(conn, buf)
			if err != nil {
				fmt.Println("Read:(maybe client close the conn) " + err.Error())
				u.errCh <- struct{}{}
			}

			fmt.Println("------server read success", string(buf))
			// bug! 如下写，buf len是1048，转string也是1048，并不是字符串实际长度！！！
			// 在通过+字符串，接收的buf长度不再一致！
			// result := "hello, " + string(buf)

			// 额外延伸，与此bug无关！，string实际len计算方式： len(rune[](str))

			// 正式的交互是返回协议的头部和body，然后根据头部读取body长度，即可返回"hello xxx"格式
			result := "hello"

			_, err = conn.Write(buf)
			if err != nil {
				fmt.Println("err Write ", err)
				u.errCh <- struct{}{}
			}

			fmt.Println("------ server write success", result)

		}
	}()

	fmt.Println("------ print wait")

	select {
	case <-time.After(time.Second * 3):
		fmt.Println("------ time deadline")
	case <-u.errCh:
		fmt.Println("------ routine err")
	}

}

func (u *UnixScoketImpl) startServer() {
	os.Remove(u.filename)
	addr, err := net.ResolveUnixAddr("unix", u.filename)
	if err != nil {
		fmt.Println("err ResolveUnixAddr ", err)
		return
	}

	listener, err := net.ListenUnix("unix", addr)

	if err != nil {
		fmt.Println("err ListenUnix ", err)
		return
	}

	defer listener.Close()

	fmt.Println("Listening on", listener.Addr())

	for {
		c, err := listener.Accept()
		if err != nil {
			panic("Accept: " + err.Error())
		}
		go u.handle(c)
	}

}

func NewUnixSocket(filename string) UnixScoket {
	return &UnixScoketImpl{
		filename,
		ProtoSize,
		make(chan struct{}),
	}
}
