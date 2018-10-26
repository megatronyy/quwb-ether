package main

import (
	"net"
	"compress/zlib"
	"bufio"
	"github.com/twfx7758/quwb-ether/common/ioutils"
)

func main() {
	/*
	在这个时候，文件已经不再局限于io，可以是一个内存buffer，
	也可以是一个计算hash的对象，甚至是一个计数器，流量限速器。
	golang灵活的接口机制为我们提供了无限可能。
	*/
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	//加上一个zip压缩，还可以利用加上crypto/aes来个AES加密...
	zip := zlib.NewWriter(conn)
	//对socket加上一个buffer来增加吞吐量
	bufconn := bufio.NewWriter(zip)
	ioutils.EncodePacket3(bufconn, []byte("hello,client \r\n"))
}