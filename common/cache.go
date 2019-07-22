package common

import (
	"github.com/BagusAK95/zaun/config"
	"github.com/go-redis/redis"
)

//Cache : cache configuration struct
type Cache struct {
	Client        *redis.Client
	Configuration *config.Configuration
}

//Fetch : fetch key
type Fetch func(key string) interface{}

//NewCache : initialized new cache
func NewCache(cl *redis.Client, co *config.Configuration) Cache {
	return Cache{
		Client:        cl,
		Configuration: co,
	}
}

//Get : get cache data
func (c *Cache) Get(key string, fn Fetch) interface{} {
	val, err := c.Client.Get(key).Result()
	if err != nil {
		return fn(key)
	}

	return val
}

//Set : set cache data
func (c *Cache) Set(key string, value interface{}) {
	c.Client.Set(key, value, c.Configuration.Redis.TTL)
}
