package client

import (
	"crawler/distributed/config"
	"crawler/distributed/engine"
	"crawler/distributed/service/common"
	"net/rpc"
)

//worker从池子里拿到client去调用rpc服务
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := common.SerializeRequest(req)
		var sResult common.ParseResult

		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return common.DeserializeResult(sResult), nil
	}
}
