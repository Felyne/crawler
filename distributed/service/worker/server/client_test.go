package main

import (
	"crawler/distributed/config"
	"crawler/distributed/service/rpcsupport"
	"crawler/distributed/service/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const addr = ":9000"
	go rpcsupport.ServeRpc(addr, worker.CrawlService{})
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(addr)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/1314495053",
		SerializedParser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "风中的蒲公英",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
