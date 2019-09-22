package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult //Url页面内容所对应的解析器
}

type ParseResult struct {
	Requests []Request
	Items []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}