package main

import (
	"gorobot/distributed/rpcsupport"
	"gorobot/engine"
	"gorobot/module"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serverRpc(host, "test1")

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	//Call Item
	item := engine.Item{
		URL:  "https://album.zhenai.com/u/1792597478",
		Type: "zhenai",
		ID:   "1792597478",
		Payload: module.Profile{
			Name:       "Amy8月的雨",
			Gender:     "女",
			Height:     160,
			Age:        30,
			Weight:     49,
			Income:     "8001-12000元",
			Marriage:   "未婚",
			Education:  "大专",
			Occupation: "人事/行政",
			Hokou:      "四川广元",
			Xinzuo:     "天蝎座(10.23-11.21)",
			House:      "已购房",
			Car:        "未买车",
			WorkAddr:   "成都青羊区",
			WorkCity:   "成都",
		},
	}

	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "" {
		t.Errorf("result:%s; err: %s", result, err)
	}

}
