package main

//"github.com/go-redis/redis/v7/internal"
//go get gopkg.in/redis.v4
//go get gopkg.in/redis.v7
import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

func createRedis() *redis.Client {
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
		panic(err)
	}

	return rdb
}

// redis.v7 的连接池管理
func connectPool(client *redis.Client) {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 100000; j++ {
				client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				client.Get(fmt.Sprintf("name%d", j)).Result()
			}

			fmt.Printf("PoolStats, TotalConns: %d, FreeConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().IdleConns)
		}()
	}

	wg.Wait()
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	rdb := createRedis()
	//onnectPool(rdb)
	mapInterface := make(map[string]interface{})
	mapInterface["name"] = "test1"
	mapInterface["age"] = 23
	rdb.HMSet("10001", mapInterface)
	fields, err := rdb.HMGet("10001", "name", "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fields1 in user_test: ", fields)

	mapInterface["name"] = "test2"
	mapInterface["age"] = 25
	rdb.HMSet("10002", mapInterface)
	fields, err = rdb.HMGet("10002", "name", "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fields2 in user_test: ", fields)

	s := rdb.PoolStats()
	fmt.Println("Hits:", s.Hits)
	fmt.Println("Misses:", s.Misses)
	fmt.Println("Timeouts:", s.Timeouts)
	fmt.Println("TotalConns:", s.TotalConns)
	fmt.Println("IdleConns:", s.IdleConns)
	fmt.Println("StaleConns:", s.StaleConns)

	stu2 := Student{
		Name: "test2",
		Age:  25,
	}

	data, err := json.Marshal(stu2)
	if err != nil {
		fmt.Println("json marshal fail")
	}
	fmt.Printf("1:%v\n", string(data))
	rdb.HSet("1", "100001", string(data))

	stu3 := Student{
		Name: "test3",
		Age:  22,
	}
	data, err = json.Marshal(stu3)
	if err != nil {
		fmt.Println("json marshal fail")
	}
	fmt.Printf("2:%v\n", string(data))
	rdb.HSet("1", "100002", string(data))
	fmt.Println("End")

	//rdb.HGet("1","100001")

	/*
		err := rdb.Set("key", "value", 0).Err()

		if err != nil {
			panic(err)
		}

		val, err := rdb.Get("key").Result()

		if err != nil {
			panic(err)
		}

		fmt.Println("key", val)

		val2, err := rdb.Get("key").Result()

		if err == redis.Nil {
			fmt.Println("missing_key not exist")
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println(val2)
		}

		//也可以这样进行设置
		rdb.Append("key3", "val3")
	*/
}
