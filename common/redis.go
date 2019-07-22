package common

import (
	"github.com/BagusAK95/zaun/config"
	"github.com/go-redis/redis"
)

//NewRedisConnection : initialized new redis connection
func NewRedisConnection(c *config.Configuration) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password, // no password set
		DB:       c.Redis.Database, // use default DB
	})
	_, err := client.Ping().Result()

	if err != nil {
		return nil, err
	}

	return client, nil
}
