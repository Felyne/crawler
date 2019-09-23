package engine

import (
	"log"

	"github.com/go-redis/redis"
)

func isDuplicate(c *redis.Client, key string, value string) bool {
	v, err := c.Get(key).Result()
	if v == "" || err == redis.Nil {
		err = c.Set(key, value, 0).Err()
		if err != nil {
			log.Println("redis set failed:", err)
		}
		return false
	}
	if err != nil {
		log.Println("redis get failed:", err)
		return false
	}
	return true
}
