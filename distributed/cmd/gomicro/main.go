package main

import (
	"github.com/Felyne/crawler/distributed/config"
	"github.com/Felyne/crawler/distributed/engine"
	"github.com/Felyne/crawler/distributed/scheduler"
	itemsaver "github.com/Felyne/crawler/distributed/service/gomicro/persist/client"
	worker "github.com/Felyne/crawler/distributed/service/gomicro/worker/client"
	"github.com/Felyne/crawler/distributed/service/gomicro/worker/pb"
	"github.com/Felyne/crawler/distributed/zhenai/parser"

	"flag"

	"github.com/go-redis/redis"
)

func main() {
	flag.Parse()
	etcdAddrs := []string{config.EtcdAddr}
	itemChan, err := itemsaver.ItemSaver(etcdAddrs)
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	if err := redisClient.Ping().Err(); err != nil {
		panic(err)
	}
	defer redisClient.Close()
	//pool := createClientPool(etcdAddrs, 1)
	//processor := worker.CreateProcessor(pool)
	processor := worker.CreateProcessor2(etcdAddrs)
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
		RedisClient:      redisClient,
	}

	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	Parser: engine.NewFuncParser(parser.ParseCityList,
	//		config.ParseCityList),
	//})

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity,
			config.ParseCity),
	})
}

//爬虫连接池
//根据启用的rpc服务数去生成相应的客户端数
//workerCount个数的goroutine去抢client调用相应服务
//rpc服务端并发处理
func createClientPool(etcdAddrs []string, n int) chan pb.CrawlerService {
	var clients []pb.CrawlerService
	for i := 0; i < n; i++ {
		clients = append(clients, worker.GetClient(etcdAddrs))
	}
	out := make(chan pb.CrawlerService)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
