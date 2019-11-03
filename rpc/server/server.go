package main

import (
	"fmt"
	rpcdemo "gorobot/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	server := rpc.NewServer()
	fmt.Println("1:rpc server main")
	server.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	fmt.Println("2:rpc server main")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		fmt.Println("rpc server start success")
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
