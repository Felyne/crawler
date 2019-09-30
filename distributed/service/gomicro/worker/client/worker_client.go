package client

import (
	"context"
	"encoding/json"

	"github.com/Felyne/crawler/distributed/engine"
	"github.com/Felyne/crawler/distributed/serializer"
	"github.com/Felyne/crawler/distributed/service/gomicro/microsupport"
	"github.com/Felyne/crawler/distributed/service/gomicro/worker/pb"

	"github.com/micro/go-micro/client"
)

func GetClient(etcdAddrs []string) pb.CrawlerService {
	cli := microsupport.NewClient(etcdAddrs,
		func(c client.Client) interface{} {
			return pb.NewCrawlerService(pb.ServiceName_name[0], c)
		})
	return cli.(pb.CrawlerService)
}

//测试
func CreateProcessor2(etcdAddrs []string) engine.Processor {
	cli := GetClient(etcdAddrs)
	return func(req engine.Request) (engine.ParseResult, error) {
		pbReq, err := getPbReq(req)
		if err != nil {
			return engine.ParseResult{}, err
		}
		pbRes, err := cli.Process(context.TODO(), pbReq)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return getEngineResult(pbRes)
	}
}

//worker从池子里拿到client去调用rpc服务
func CreateProcessor(clientChan chan pb.CrawlerService) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		pbReq, err := getPbReq(req)
		if err != nil {
			return engine.ParseResult{}, err
		}
		c := <-clientChan
		pbRes, err := c.Process(context.TODO(), pbReq)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return getEngineResult(pbRes)
	}
}

func getPbReq(req engine.Request) (*pb.Request, error) {
	sReq := serializer.SerializeRequest(req)
	args, err := json.Marshal(sReq.SerializedParser.Args)
	if err != nil {
		return nil, err
	}
	return &pb.Request{
		Url: sReq.Url,
		SerializedParser: &pb.SerializedParser{
			Name: sReq.SerializedParser.Name,
			Args: args,
		},
	}, nil
}

func getEngineResult(pbRes *pb.ParseResult) (engine.ParseResult, error) {
	var items []engine.Item
	for _, t := range pbRes.Items {
		var payload interface{}
		if err := json.Unmarshal(t.Payload, &payload); err != nil {
			return engine.ParseResult{}, err
		}
		item := engine.Item{
			Url:     t.Url,
			Type:    t.Type,
			Id:      t.Id,
			Payload: payload,
		}
		items = append(items, item)
	}
	var requests []engine.Request
	for _, r := range pbRes.Requests {
		var args interface{}
		if err := json.Unmarshal(r.SerializedParser.Args, &args); err != nil {
			return engine.ParseResult{}, err
		}
		sReq := serializer.Request{
			Url: r.Url,
			SerializedParser: serializer.SerializedParser{
				Name: r.SerializedParser.Name,
				Args: args,
			},
		}
		engineReq, err := serializer.DeserializeRequest(sReq)
		if err != nil {
			return engine.ParseResult{}, err
		}
		requests = append(requests, engineReq)
	}

	return engine.ParseResult{
		Items:    items,
		Requests: requests,
	}, nil
}
