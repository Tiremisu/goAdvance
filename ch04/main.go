package main

import(
	"fmt"
	"net"
	"net/http"
	"log"
	"net/rpc"
	server "./server"
	
)

func main() {
	http.HandleFunc("/", home) //DefaultServeMutex
	

	//register a rpc service
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":9090")
	if e != nil {
		log.Fatal("rpc listen error:", e)
	}
	go http.Serve(l, nil)

	client, err := rpc.DialHTTP("tcp", "localhost" + ":9090")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	args := &server.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}
	fmt.Printf("Arith: %d*%d=%d \n", args.A, args.B, reply)

	//start a http server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to RPC world")
}
