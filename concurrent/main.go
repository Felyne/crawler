package main

import (
	"github.com/Felyne/crawler/concurrent/engine"
	"github.com/Felyne/crawler/concurrent/persist"
	"github.com/Felyne/crawler/concurrent/scheduler"
	"github.com/Felyne/crawler/concurrent/zhenai/parser"
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
