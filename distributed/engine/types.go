package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult //Url页面内容所对应的解析器
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

//存储的一条记录
type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
