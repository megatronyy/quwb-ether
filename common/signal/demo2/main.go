package main

import (
	"os"
	"os/signal"
	"fmt"
)
//监听指定信号
func main() {
	//合建chan
	c := make(chan os.Signal)
	//监听所有信号
	signal.Notify(c, os.Interrupt, os.Kill)
	//阻塞直到有信号传入
	fmt.Println("启动了")
	s := <-c
	fmt.Println("退出信号：", s)
}
