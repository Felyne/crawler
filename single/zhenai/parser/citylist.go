package parser

import (
	"crawler/single/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
const limit = 5

//根据城市列表页面解析各个城市的url
func ParseCityList(contents []byte) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	n := limit
	//m[1]是城市对应的url，m[2]是城市名
	for _, m := range matches {
		result.Items = append(result.Items, "City: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity, //该Url对应的的解析器
		})
		n--
		if n == 0 {
			break
		}
	}

	return result
}
