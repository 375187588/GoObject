package main
//https://github.com/375187588/GoObject
import (
	"gorobot/engine"
	"gorobot/persist"
	"gorobot/scheduler"
	"gorobot/zhenai/parser"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	itemChan, err := persist.ItemSaver("data_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		GUrl:       "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
