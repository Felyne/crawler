package client

import (
	"net/rpc"

	"github.com/Felyne/crawler/distributed/config"
	"github.com/Felyne/crawler/distributed/engine"
	"github.com/Felyne/crawler/distributed/serializer"
)

//worker从池子里拿到client去调用rpc服务
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := serializer.SerializeRequest(req)
		var sResult serializer.ParseResult

		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return serializer.DeserializeResult(sResult), nil
	}
}
