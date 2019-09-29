package persist

import (
	"context"
	"crawler/distributed/service/gomicro/persist/pb"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/astaxie/beego/config"
	"gopkg.in/olivere/elastic.v5"
)

var ErrConfig = errors.New("config error")

type ItemSaverService struct {
	client *elastic.Client
	index  string //类似数据库名
}

func NewItemSaverService(cfg config.Configer) (*ItemSaverService, error) {
	elasticURL := cfg.String("ElasticURL") //读取配置，初始化服务
	index := cfg.String("ElasticIndex")
	if elasticURL == "" || index == "" {
		return nil, ErrConfig
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(elasticURL))
	if err != nil {
		return nil, err
	}
	return &ItemSaverService{
		client: client,
		index:  index,
	}, nil
}

func (s *ItemSaverService) Save(ctx context.Context, item *pb.Item, resp *pb.Resp) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	var payload interface{}
	err := json.Unmarshal(item.Payload, &payload)
	if err != nil {
		fmt.Println("fuck 1111")
		return err
	}
	newItem := struct {
		Url     string
		Type    string
		Id      string
		Payload interface{}
	}{
		Url:     item.Url,
		Type:    item.Type,
		Id:      item.Id,
		Payload: payload,
	}
	indexService := s.client.Index().
		Index(s.index).
		Type(newItem.Type).
		BodyJson(newItem)
	if newItem.Id != "" {
		indexService.Id(newItem.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		fmt.Println("fuck 2222")
		return err
	}
	fmt.Println("success:", item.Id)
	resp.Result = "ok"
	//fmt.Printf("%+v", resp)
	return nil
}
