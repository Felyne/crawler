package persist

import (
	"context"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			//_, err := save(item)
			//if err != nil {
			//	log.Printf("Item Saver: error saving item %v: %v", item, err)
			//}
		}
	}()
	return out
}

func save(item interface{}) (string, error) {
	//must turn off sniff in docker
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	//Index()存数据
	resp, err := client.Index().
		Index("db_test").
		Type("tb_test").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	//log.Printf("%+v", resp)
	return resp.Id, nil
}
