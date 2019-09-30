package worker

import (
	"context"
	"encoding/json"

	"github.com/Felyne/crawler/distributed/service/gomicro/common"

	"github.com/Felyne/crawler/distributed/engine"

	"github.com/Felyne/crawler/distributed/serializer"
	"github.com/Felyne/crawler/distributed/service/gomicro/worker/pb"
)

type CrawlerService struct {
}

func (c CrawlerService) Process(ctx context.Context, pbReq *pb.Request, pbRes *pb.ParseResult) error {
	engineReq, err := getEngineReq(pbReq)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	pbRes, err = getPbResult(engineResult)
	if err != nil {
		return err
	}
	return nil
}

func getEngineReq(pbReq *pb.Request) (engine.Request, error) {
	var args interface{}
	err := json.Unmarshal(pbReq.SerializedParser.Args, &args)
	if err != nil {
		return engine.Request{}, err
	}
	sReq := serializer.Request{
		Url: pbReq.Url,
		SerializedParser: serializer.SerializedParser{
			Name: pbReq.SerializedParser.Name,
			Args: args,
		},
	}
	return serializer.DeserializeRequest(sReq)
}

func getPbResult(result engine.ParseResult) (*pb.ParseResult, error) {
	var items []*common.Item
	for _, t := range result.Items {
		payload, err := json.Marshal(t.Payload)
		if err != nil {
			return nil, err
		}
		items = append(items, &common.Item{
			Url:     t.Url,
			Type:    t.Type,
			Id:      t.Id,
			Payload: payload,
		})
	}
	var requests []*pb.Request
	for _, r := range result.Requests {
		sReq := serializer.SerializeRequest(r)
		args, err := json.Marshal(sReq.SerializedParser.Args)
		if err != nil {
			return nil, err
		}
		requests = append(requests, &pb.Request{
			Url: sReq.Url,
			SerializedParser: &pb.SerializedParser{
				Name: sReq.SerializedParser.Name,
				Args: args,
			},
		})
	}
	return &pb.ParseResult{
		Items:    items,
		Requests: requests,
	}, nil

}
