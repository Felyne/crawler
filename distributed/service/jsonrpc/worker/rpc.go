package worker

import (
	"github.com/Felyne/crawler/distributed/engine"
	"github.com/Felyne/crawler/distributed/serializer"
)

type CrawlService struct{}

//要注册的rpc服务
//传入序列化的参数，得到序列化的结果
func (CrawlService) Process(req serializer.Request, result *serializer.ParseResult) error {
	engineReq, err := serializer.DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = serializer.SerializeResult(engineResult)
	return nil
}
