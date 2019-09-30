package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Felyne/crawler/distributed/service/jsonrpc/rpcsupport"
	"github.com/Felyne/crawler/distributed/service/jsonrpc/worker"
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
