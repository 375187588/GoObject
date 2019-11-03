package main

import (
	"net/rpc/jsonrpc"
	"log"
	"net"
	"net/rpc"
	"gorobot/rpc"
	"fmt"
)

func main() {
	fmt.Println("1:rpc server main")
	rpc.Register(rpcdemo.DemoService{})
	listener,err := net.Listen("tcp",":1234")
	if err != nil{
		panic(err)
	}
	fmt.Println("2:rpc server main")
	defer listener.Close()
	for{
		conn,err:=listener.Accept()
		if err!= nil{
			log.Printf("accept error: %v",err)
			continue
		}
		go jsonrpc.NewServerCodec(conn)
	}

}
