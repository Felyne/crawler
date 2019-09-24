package worker

import (
	"crawler/distributed/config"
	"crawler/distributed/engine"
	"crawler/distributed/zhenai/parser"
	"errors"
	"fmt"
	"log"
)

//把Parser序列化用于传输
type SerializedParser struct {
	Name string
	Args interface{}
}

// {"ParseCityList", nil}, {"ProfileParser", userName}

//用来传输的Request
type Request struct {
	Url              string
	SerializedParser SerializedParser
}

//用来传输的ParseResult
type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

//序列化engine.Request
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		SerializedParser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

//序列化engine.ParseResult
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

//解序列化Request
func DeserializeRequest(r Request) (engine.Request, error) {
	p, err := deserializeParser(r.SerializedParser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: p,
	}, nil
}

//解序列化ParseResult
func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializeRequest request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

//解序列化Parser
//一种是把Parser的名字注册到一个全局的map去,通过map找到对应的函数
//这里使用的是第二种switch
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(
			parser.ParseCity,
			config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid args: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}
