package main

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/persist"
	"crawler/concurrent/scheduler"
	"crawler/concurrent/zhenai/parser"
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
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
