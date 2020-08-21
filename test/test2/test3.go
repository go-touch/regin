package main

import (
	"fmt"
	"log"
)

type LogWriter struct {
	//io.Writer
}

func (lr *LogWriter) Write(p []byte) (n int, err error) {

	fmt.Println(string(p))
	fmt.Println("hello logging...", `\n`)

	return n, err
}

func main() {
	logger := log.New(&LogWriter{},"测试日志",3)


	logger.Printf("我在打印日志")

}
