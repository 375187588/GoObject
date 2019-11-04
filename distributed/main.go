package main

//https://github.com/375187588/GoObject
//先启动 itemsaver.go
//在启动 worker.go
//最后启动 main.go

import (
	"gorobot/distributed/config"
	itemsaver "gorobot/distributed/persist/client"
	worker "gorobot/distributed/worker/client"
	"gorobot/engine"
	"gorobot/scheduler"
	"gorobot/zhenai/parser"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	itemChan, err := itemsaver.ItemSaver(":1234", "dating_profile")
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		GUrl: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
