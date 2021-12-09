package main

import (
	"fmt"
	"io"
	"os"

	"bufio"
)

func _readLine(r *bufio.Reader) (totalLone []byte, err error) {
	var line []byte
	var more bool
	for {
		line, more, err = r.ReadLine()
		// fmt.Println("------ print line", line)
		// 当前一行太长没读完，more是true,还要继续读
		// fmt.Println("------ print more", more)
		// fmt.Println("------ print err", err)
		// 结束
		if err != io.EOF {
			totalLone = append(totalLone, line...)
			return
		}

		// 错误
		if err != nil {
			fmt.Println("err ", err)
			return
		}

		// 空文件
		if line == nil && !more {
			return
		}
		// 一行还没读完
		if more {
			totalLone = append(totalLone, line...)
			continue
		}

	}
}

func readFile(r *bufio.Reader) ([]byte, error) {
	text := []byte{}
	for {
		line, err := _readLine(r)
		// 读完了
		if err == io.EOF {
			fmt.Println("------ read over /n")
			return text, nil
		}
		// 错误
		if err != nil {
			fmt.Println("err ", err)
			return text, err
		}
		text = append(text, line...)
	}
}

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	read := bufio.NewReader(file)

	text, _ := readFile(read)
	fmt.Println("------ print text", string(text))
}
