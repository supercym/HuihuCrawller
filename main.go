package main

import (
	"huihuCrawler02/engine"
	"huihuCrawler02/parser/huihu"
	"huihuCrawler02/persist"
	"huihuCrawler02/scheduler"
)

func main() {
	itemChan, err := persist.ItemSaver("huihu")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		//Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:10,
		ItemChan:itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.hhkaoyan.com/teacher/major.html",
		ParserFunc: huihu.ParseCollegeList,
	})

}
