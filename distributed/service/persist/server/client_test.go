package main

import (
	"crawler/distributed/config"
	"crawler/distributed/engine"
	"crawler/distributed/model"
	"crawler/distributed/service/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const addr = ":1234"
	go serveRpc(addr, config.ElasticURL, config.ElasticIndex)
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(addr)
	if err != nil {
		panic(err)
	}
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1314495053",
		Type: "zhenai",
		Id:   "1314495053",
		Payload: model.Profile{
			Name:       "风中的蒲公英",
			Gender:     "女",
			Age:        41,
			Height:     158,
			Weight:     48,
			Income:     "3001-5000元",
			Marriage:   "离异",
			Education:  "中专",
			Occupation: "公务员",
			Hokou:      "四川阿坝",
			Xinzuo:     "处女座",
			House:      "已购房",
			Car:        "未购车",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
