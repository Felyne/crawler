package parser

import (
	"regexp"

	"github.com/Felyne/crawler/distributed/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	//cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

//根据城市的第一页来匹配页面中的用户
func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	//m[1]是url，m[2]人名

	for _, m := range matches {
		//result.Items = append(result.Items, "User: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	//页面底部的相关链接
	//matches = cityUrlRe.FindAllSubmatch(contents, -1)
	//for _, m := range matches {
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url:        string(m[1]),
	//		ParserFunc: ParseCity,
	//	})
	//}

	return result
}
