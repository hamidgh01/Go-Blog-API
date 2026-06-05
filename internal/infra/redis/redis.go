package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hamidgh01/Go-Blog-API/config"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func InitRedis(cfg *config.RedisConf) (*redis.Client, error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}

	log.Println("[INFO] Redis connection established successfully.") // log.Info("[INFO] [Data] [Msg]")
	return client, nil
}
