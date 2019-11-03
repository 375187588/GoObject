package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// ServerRpc
func ServerRpc(host string, service interface{}) error {
	server := rpc.NewServer()
	server.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

	return nil
}

// NewClient
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	client := jsonrpc.NewClient(conn)
	return client, err
}
