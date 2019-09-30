package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Felyne/crawler/distributed/config"
	"github.com/Felyne/crawler/distributed/service/jsonrpc/persist"
	"github.com/Felyne/crawler/distributed/service/jsonrpc/rpcsupport"
)

var port = flag.Int("port", 0, "server listen port")

//ItemSaver服务端注册
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	listenAddr := fmt.Sprintf(":%d", *port)
	log.Fatal(serveRpc(listenAddr,
		config.ElasticURL,
		config.ElasticIndex))
}

func serveRpc(listenAddr, elasticURL, index string) error {
	s, err := persist.NewItemSaverService(elasticURL, index)
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(listenAddr, s)
}
