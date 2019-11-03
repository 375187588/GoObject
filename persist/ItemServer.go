package persist

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorobot/engine"
	"time"

	"github.com/go-redis/redis"
)

func createRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:               "127.0.0.1:6379",
		DB:                 0,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           10,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        time.Minute,
		IdleCheckFrequency: 100 * time.Millisecond,
	})

	_, err := rdb.Ping().Result()

	if err != nil {
		return nil, err
	}

	return rdb, nil
}

//ItemSaver chan interface
func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	rdb, err := createRedis()
	if err != nil {
		return nil, err
	}
	go func() {
		var itemCount int64
		itemCount = 0
		for {
			item := <-out

			err = Save(itemCount, item, rdb)
			if err == nil {
				//log.Printf("Item saver:got item "+"#%d: %v", itemCount, item)
				itemCount++
				fmt.Println(itemCount)
			}
		}
	}()
	return out, nil
}

// Save(index )
func Save(index int64, item engine.Item, rdb *redis.Client) error {
	data, err := json.Marshal(item)
	if err != nil {
		fmt.Println("json marshal fail")
		return err
	}

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	if item.ID == "" {
		return errors.New("must supply ID")
	}

	//fmt.Printf("2:%v\n", string(data))
	rdb.HSet(item.Type, item.ID, string(data))
	return nil
}
