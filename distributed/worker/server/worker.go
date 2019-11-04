package main

import (
	"fmt"
	"gorobot/distributed/config"
	"gorobot/distributed/rpcsupport"
	"gorobot/distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServerRpc(
		fmt.Sprintf(":%d", config.WorkerPort0),
		//":8888",
		worker.RobotService{}))
}
