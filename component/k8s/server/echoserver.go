package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
	"io"
)

func main() {
	port := fmt.Sprintf(":%s", os.Args[1])
	prefix := os.Args[2]

	//在给定的端口上创建tcp监听
	listener, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}

	fmt.Printf("listening on %s, prefix: %s\n", listener.Addr(), prefix)

	// 监听新的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}

		// 启用goroutine处理连接
		go handleConnection(conn, prefix)
	}
}

func handleConnection(conn net.Conn, prefix string) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request: %s\n", bytes)
		// 添加前缀作为response返回
		line := fmt.Sprintf("%s %s", prefix, bytes)
		fmt.Printf("response: %s\n", line)

		conn.Write([]byte(line))
	}
}
