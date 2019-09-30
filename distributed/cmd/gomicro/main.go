package main

import (
	"github.com/Felyne/crawler/distributed/config"
	"github.com/Felyne/crawler/distributed/engine"
	"github.com/Felyne/crawler/distributed/scheduler"
	itemsaver "github.com/Felyne/crawler/distributed/service/gomicro/persist/client"
	"github.com/Felyne/crawler/distributed/service/jsonrpc/rpcsupport"
	worker "github.com/Felyne/crawler/distributed/service/jsonrpc/worker/client"
	"github.com/Felyne/crawler/distributed/zhenai/parser"

	"flag"
	"log"
	"net/rpc"
	"strings"

	"github.com/go-redis/redis"
)

var workerHosts = ":9000,:9001"

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
	pool := createClientPool(strings.Split(workerHosts, ","))
	processor := worker.CreateProcessor(pool)
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
func createClientPool(addrs []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range addrs {
		c, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, c)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
