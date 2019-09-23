package main

import (
	"crawler/distributed/engine"
	"crawler/distributed/persist"
	"crawler/distributed/scheduler"
	"crawler/distributed/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         persist.ItemSaver(),
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
