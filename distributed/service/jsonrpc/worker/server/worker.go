package main

import (
	"crawler/distributed/service/jsonrpc/rpcsupport"
	"crawler/distributed/service/jsonrpc/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "server listen port")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
