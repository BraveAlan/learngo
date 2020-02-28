package main

import (
	"imooc.com/sb/learngo/crawler/persist"
	"imooc.com/sb/learngo/crawler/scheduler"
	"imooc.com/sb/learngo/crawler/zhenai/parser"

	"imooc.com/sb/learngo/crawler/engine"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 5,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
