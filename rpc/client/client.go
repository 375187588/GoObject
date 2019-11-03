package main

import (
	"gorobot/rpc"
	"net"
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {
	fmt.Println("1:rpc client main")
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	fmt.Println("2:rpc client main")
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div",
	rpcdemo.Args{10,3},&result)
	fmt.Println("3:rpc client main")
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println(result)
	}

	err = client.Call("DemoService.Div",
	rpcdemo.Args{10,0},&result)
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println(result)
	}
}
