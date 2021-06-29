package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//var buf [1024]byte

	for {
		//n, err := reader.Read(buf[:]) //read conn data to buf
		bytes, err := reader.ReadBytes('\n')

		//line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		//recvStr := string(buf[:n]) //data type transfer
		recvStr := string(bytes[:]) //data type transfer
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000") //start a tcp server
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept() //listen to accept client connection
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) //start a goroutine to handler connection
	}
}
