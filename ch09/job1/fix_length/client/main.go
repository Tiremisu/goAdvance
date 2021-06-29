package main

import (
	"bytes"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000") //connect to the tcp server which serve at localhost;3000
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	// start 20 connections to the tcp server
	for i := 0; i < 10; i++ {
		msg := `Hello, i. How are you?`
		buf := make([]byte, 1024)
		br := bytes.NewReader([]byte(msg)) //从msg中读数据
		br.Read(buf[:])                    //

		conn.Write(buf)
	}
}
