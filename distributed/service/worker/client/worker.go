package client

import (
	"crawler/distributed/config"
	"crawler/distributed/engine"
	"crawler/distributed/service/worker"
	"net/rpc"
)

//worker从池子里拿到client去调用rpc服务
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult

		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
