package util

import (
	"github.com/go-redis/redis"
)

const maxCount = 100

type handlerFunc func(key string, values []interface{}) error

// redis的lpush操作不能一次写太多数据
func LPushMany(c *redis.Client, key string, data []interface{}) error {
	fn := func(key string, values []interface{}) error {
		return c.LPush(key, values...).Err()
	}
	return segment(fn, key, data)
}

func RPushMany(c *redis.Client, key string, data []interface{}) error {
	fn := func(key string, values []interface{}) error {
		return c.RPush(key, values...).Err()
	}
	return segment(fn, key, data)
}

func segment(fn handlerFunc, key string, data []interface{}) error {
	l := len(data)
	x, y := l/maxCount, l%maxCount
	if x > 0 {
		for i := 0; i < x; i++ {
			start := i * maxCount
			end := start + maxCount - 1
			if err := fn(key, data[start:end]); err != nil {
				return err
			}
		}
	}
	if y > 0 {
		if err := fn(key, data[x*maxCount:]); err != nil {
			return err
		}
	}
	return nil
}

func DelMany(c *redis.Client, keys []string) error {
	l := len(keys)
	x, y := l/maxCount, l%maxCount
	if x > 0 {
		for i := 0; i < x; i++ {
			start := i * maxCount
			end := start + maxCount - 1
			if err := c.Del(keys[start:end]...).Err(); err != nil {
				return err
			}
		}
	}
	if y > 0 {
		if err := c.Del(keys[x*maxCount:]...).Err(); err != nil {
			return err
		}
	}
	return nil
}
