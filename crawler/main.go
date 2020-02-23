package main

import (
	"imooc.com/sb/learngo/crawler/zhenai/parser"

	"imooc.com/sb/learngo/crawler/engine"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
