package persist

import (
	"context"
	"crawler/distributed/engine"
	"crawler/distributed/model"
	"encoding/json"
	"testing"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	const elasticURL = "http://127.0.0.1:9200"
	const index = "dating_test"
	//here using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(elasticURL))
	if err != nil {
		panic(err)
	}

	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%v", resp.Source)
	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}
