package main

import (
	"parallelCrawler/huihu/parser"
	"parallelCrawler/engine"
	"parallelCrawler/persist"
	"parallelCrawler/scheduler"
)



func main() {
	itemChan, err := persist.ItemSaver("college_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan:itemChan,
		RequestProcessor:engine.Worker,
	}
	e.Run(engine.Request{
		Url:		"http://www.hhkaoyan.com/teacher/major.html",
		Parser:engine.NewFuncParser(parser.ParseCollege, "parseCollege"),
	})

	//e.Run(engine.Request{
	//	Url:		"http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc:	parser.ParseCity,
	//})
}
