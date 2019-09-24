package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//服务端
func ServeRpc(addr string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Printf("listening on %s", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

//客户端
func NewClient(addr string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
