package persist

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Felyne/crawler/distributed/service/gomicro/common"

	"github.com/Felyne/crawler/distributed/engine"

	"github.com/Felyne/crawler/distributed/service/gomicro/persist/pb"

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

func (s *ItemSaverService) Save(ctx context.Context, item *common.Item, resp *pb.Resp) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	var payload interface{}
	err := json.Unmarshal(item.Payload, &payload)
	if err != nil {
		return err
	}
	newItem := engine.Item{
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
		return err
	}
	resp.Result = "ok"
	//fmt.Printf("%+v", resp)
	return nil
}
