package main

import (
	"crawler/distributed/config"
	"crawler/distributed/engine"
	"crawler/distributed/persist"
	"crawler/distributed/scheduler"
	"crawler/distributed/zhenai/parser"

	"github.com/go-redis/redis"
)

func main() {
	itemChan, err := persist.ItemSaver(config.ElasticURL, config.ElasticIndex)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	if err := client.Ping().Err(); err != nil {
		panic(err)
	}
	defer client.Close()

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
		RedisClient:      client,
	}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})

}
