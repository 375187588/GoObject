package client

import (
	"gorobot/distributed/config"
	"gorobot/distributed/rpcsupport"
	"gorobot/engine"
)

//ItemSaver chan interface
func ItemSaver(host, index string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		var itemCount int64
		itemCount = 0
		for {
			item := <-out
			//log.Printf("Item Saver: got item "+"#%d : %v", itemCount, item)
			itemCount++
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err == nil {
				itemCount++
			}
		}
	}()
	return out, nil
}
