package main

import (
	"github.com/Felyne/crawler/single/engine"
	"github.com/Felyne/crawler/single/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
