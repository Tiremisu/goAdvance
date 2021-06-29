package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReaderSize(conn, 1024) //读取connection中固定的数据到流中
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:1024]) //write connection data to buf
		if err != nil {
			return
		}
		fmt.Println("收到client发来的数据：", string(buf[:]))
		fmt.Println("收到数据： ", n)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) //处理connection
	}
}
