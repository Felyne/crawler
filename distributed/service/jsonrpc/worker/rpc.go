package worker

import (
	"crawler/distributed/engine"
	"crawler/distributed/service/common"
)

type CrawlService struct{}

//要注册的rpc服务
//传入序列化的参数，得到序列化的结果
func (CrawlService) Process(req common.Request, result *common.ParseResult) error {
	engineReq, err := common.DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = common.SerializeResult(engineResult)
	return nil
}
