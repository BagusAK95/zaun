package common

import (
	"encoding/json"

	"github.com/BagusAK95/zaun/config"
	"github.com/go-redis/redis"
)

//Cache : cache configuration struct
type Cache struct {
	Client        *redis.Client
	Configuration *config.Configuration
}

//NewCache : initialized new cache
func NewCache(cl *redis.Client, co *config.Configuration) Cache {
	return Cache{
		Client:        cl,
		Configuration: co,
	}
}

//Get : get cache data
func (c *Cache) Get(key string) (string, error) {
	val, err := c.Client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

//Set : set cache data
func (c *Cache) Set(key string, value interface{}) error {
	obj, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.Client.Set(key, string(obj), c.Configuration.Redis.TTL).Err()
}
