package main

import (
	"github.com/Felyne/crawler/distributed/config"
	"github.com/Felyne/crawler/distributed/engine"
	"github.com/Felyne/crawler/distributed/persist"
	"github.com/Felyne/crawler/distributed/scheduler"
	"github.com/Felyne/crawler/distributed/zhenai/parser"

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
	//	Url:    "http://www.zhenai.com/zhenghun",
	//	Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	//})

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})

}
