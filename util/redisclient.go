package util

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var pool *redis.Pool

func InitRedis() {

	pool = NewRedisPool(viper.GetString("redis.host"),
		viper.GetString("redis.password"),
		viper.GetInt("redis.port"),
		0)
}

func NewRedisPool(redis_host, password string, redis_port, db int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     viper.GetInt("maxIdle"),
		MaxActive:   viper.GetInt("maxActive"),
		IdleTimeout: 480 * time.Second,
		Dial: func() (redis.Conn, error) {
			timeout := time.Duration(2) * time.Second
			c, err := redis.DialTimeout("tcp", fmt.Sprintf("%s:%d", redis_host, redis_port), timeout, 0, 0)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if db > 0 && db < 16 {
				if _, err := c.Do("SELECT", db); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
}
func Get() redis.Conn {
	return pool.Get()
}
