package persist

import (
	"context"
	"crawler/distributed/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/olivere/elastic.v5"
)

func TestItemSaver(t *testing.T) {
	expected := model.Profile{
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
	}
	var err error
	id, err := save(expected)
	assert.Equal(t, err, nil)
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	assert.Equal(t, err, nil)
	resp, err := client.Get().
		Index("db_test").
		Type("tb_test").
		Id(id).
		Do(context.Background())
	assert.Equal(t, err, nil)
	actual, err := model.FromJsonObj(*resp.Source)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, actual)

	_, _ = client.Delete().Index("db_test").Type("tb_test").Id(id).Do(context.Background())
}
