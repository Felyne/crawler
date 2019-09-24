package persist

import (
	"crawler/distributed/engine"
	"crawler/distributed/persist"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	client *elastic.Client
	index  string //类似数据库名
}

func NewItemSaverService(elasticURL, index string) (*ItemSaverService, error) {
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

//要注册的rpc服务
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.client, s.index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
