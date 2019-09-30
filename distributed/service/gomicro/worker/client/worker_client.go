package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Felyne/crawler/distributed/engine"
	"github.com/Felyne/crawler/distributed/serializer"
	"github.com/Felyne/crawler/distributed/service/gomicro/microsupport"
	"github.com/Felyne/crawler/distributed/service/gomicro/worker/pb_crawler"
	"github.com/micro/go-micro/client"
)

func GetClient(etcdAddrs []string) pb_crawler.CrawlerService {
	cli := microsupport.NewClient(etcdAddrs,
		func(c client.Client) interface{} {
			return pb_crawler.NewCrawlerService(pb_crawler.ServiceName_name[0], c)
		})
	return cli.(pb_crawler.CrawlerService)
}

//worker从池子里拿到client去调用rpc服务
func CreateProcessor(clientChan chan pb_crawler.CrawlerService) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		pbReq, err := getPbReq(req)
		if err != nil {
			return engine.ParseResult{}, err
		}
		c := <-clientChan
		fmt.Println("process() start")
		pbRes, err := c.Process(context.TODO(), pbReq)
		if err != nil {
			log.Printf("Process() rpc error: %s\n", err)
			return engine.ParseResult{}, err
		}
		fmt.Println(len(pbRes.Requests), len(pbRes.Items))
		fmt.Println("process() end")

		return getEngineResult(pbRes)
	}
}

func getPbReq(req engine.Request) (*pb_crawler.Request, error) {
	sReq := serializer.SerializeRequest(req)
	args, err := json.Marshal(sReq.SerializedParser.Args)
	if err != nil {
		return nil, err
	}
	return &pb_crawler.Request{
		Url: sReq.Url,
		SerializedParser: &pb_crawler.SerializedParser{
			Name: sReq.SerializedParser.Name,
			Args: args,
		},
	}, nil
}

func getEngineResult(pbRes *pb_crawler.ParseResult) (engine.ParseResult, error) {
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
	//fmt.Println("req count:", len(requests), "item count:", len(items))
	return engine.ParseResult{
		Items:    items,
		Requests: requests,
	}, nil
}
