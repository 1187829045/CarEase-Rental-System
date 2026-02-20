package Init

import (
	"context"
	"time"

	"car.rental/global"
	"github.com/redis/go-redis/v9"
)

func RedisInit() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	if err := global.RedisClient.Ping(ctx).Err(); err != nil {
		panic(err)
	}
}
