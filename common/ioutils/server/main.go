package main

import (
	"net"
	"github.com/twfx7758/quwb-ether/common/ioutils"
)

func main() {
	/*
	在这个时候，文件已经不再局限于io，可以是一个内存buffer，
	也可以是一个计算hash的对象，甚至是一个计数器，流量限速器。
	golang灵活的接口机制为我们提供了无限可能。
	*/
	lister, _ := net.Listen("tcp", ":8020")
	defer lister.Close()
	for {
		conn, _ := lister.Accept()

		//向客户端发送数据，并关闭连接
		ioutils.DecodePacket(conn)
		conn.Close()
	}
}