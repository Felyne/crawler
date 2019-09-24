package worker

import "crawler/distributed/engine"

type CrawlService struct{}

//要注册的rpc服务
//传入序列化的参数，得到序列化的结果
func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
