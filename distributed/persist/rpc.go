package persist

import (
	"gorobot/engine"
	"gorobot/persist"
	"log"

	"github.com/go-redis/redis"
)

// ItemSaverService struct
type ItemSaverService struct {
	Client *redis.Client
	Index  string
}

var itemCount int64

// Save item
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	itemCount += 0
	err := persist.Save(itemCount, item, s.Client)
	//log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v %v", item, err)
	}
	return err
}
