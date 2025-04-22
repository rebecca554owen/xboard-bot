package redis

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"xboard-bot/config"

	redis "github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	once   sync.Once
)

const Nil = redis.Nil

type Pipeliner interface {
	redis.Pipeliner
}

func GetRedisClient(cfg *config.Config) *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
			Password:     cfg.Redis.Password,
			DB:           cfg.Redis.Database,
			MaxRetries:   3,
			DialTimeout:  time.Second * 5,
			ReadTimeout:  time.Second * 20,
			WriteTimeout: time.Second * 20,
			PoolSize:     50,
			MinIdleConns: 2,
			PoolTimeout:  time.Minute,
		})

		if err := client.Ping(context.Background()).Err(); err != nil {
			panic(err)
		}
		log.Println("Redis连接成功")
	})

	return client
}