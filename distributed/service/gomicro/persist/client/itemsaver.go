package client

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Felyne/crawler/distributed/engine"
	"github.com/Felyne/crawler/distributed/service/gomicro/common"
	"github.com/Felyne/crawler/distributed/service/gomicro/microsupport"
	"github.com/Felyne/crawler/distributed/service/gomicro/persist/pb"

	"github.com/micro/go-micro/client"
)

func GetClient(etcdAddrs []string) pb.ItemSaverService {
	cli := microsupport.NewClient(etcdAddrs,
		func(c client.Client) interface{} {
			return pb.NewItemSaverService(pb.ServiceName_name[0], c)
		})
	return cli.(pb.ItemSaverService)
}

//ItemSaver客户端
func ItemSaver(etcdAddrs []string) (chan engine.Item, error) {
	cli := GetClient(etcdAddrs)
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item  #%d: %v", itemCount, item)
			itemCount++
			//Call RPC to save item
			payload, err := json.Marshal(item.Payload)
			if err != nil {
				if err != nil {
					log.Printf("Item Saver: error saving item %v: %v", item, err)
				}
			}
			pbItem := &common.Item{
				Url:     item.Url,
				Type:    item.Type,
				Id:      item.Id,
				Payload: payload,
			}
			_, err = cli.Save(context.TODO(), pbItem)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
