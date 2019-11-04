package main

import (
	"fmt"
	"gorobot/distributed/config"
	"gorobot/distributed/rpcsupport"
	"gorobot/distributed/worker"
	"testing"
	"time"
)

func TestRobotService(t *testing.T) {
	const host = ":8888"
	go rpcsupport.ServerRpc(host, worker.RobotService{})
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "https://album.zhenai.com/u/1792597478",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "Amy8月的雨",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.RobotServiceRpc, req, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
