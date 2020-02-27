package main

import (
	"imooc.com/sb/learngo/crawler/scheduler"
	"imooc.com/sb/learngo/crawler/zhenai/parser"

	"imooc.com/sb/learngo/crawler/engine"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 5,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
