package main

//分步式先启动 gorobot/distributed/persist/server/main.go

import (
	"gorobot/distributed/config"
	"gorobot/distributed/persist"
	"gorobot/distributed/rpcsupport"

	"github.com/go-redis/redis"
)

func main() {
	err := serverRpc(config.ItemServerPort, config.RedisIndex)
	if err != nil {
		panic(err)
	}
}

// serverRpc(host, index string)
func serverRpc(host, index string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 5,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	return rpcsupport.ServerRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
