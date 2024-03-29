package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/Felyne/crawler/distributed/config"
	"github.com/Felyne/crawler/distributed/serializer"
	"github.com/Felyne/crawler/distributed/service/jsonrpc/rpcsupport"
	worker2 "github.com/Felyne/crawler/distributed/service/jsonrpc/worker"
)

func TestCrawlService(t *testing.T) {
	const addr = ":9000"
	go rpcsupport.ServeRpc(addr, worker2.CrawlService{})
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(addr)
	if err != nil {
		panic(err)
	}
	req := serializer.Request{
		Url: "http://album.zhenai.com/u/1314495053",
		SerializedParser: serializer.SerializedParser{
			Name: config.ParseProfile,
			Args: "风中的蒲公英",
		},
	}
	var result serializer.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
