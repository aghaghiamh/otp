package adaptor

import (
	"context"
	"otp/src/pkg/config"
	"otp/src/pkg/log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// TODO: for testing purpose better to implement a Cache interface with required methods (GET/SET/...),
// For the sake of time, I just ignore it.
func CreateRedisClient() *redis.Client {
	cnf := config.GetAppConfigInstance()

	cache := redis.NewClient(&redis.Options{
		Addr:                  cnf.Cache.Host + ":" + strconv.Itoa(cnf.Cache.Port),
		Password:              cnf.Cache.Password,
		DB:                    cnf.Cache.DB,
		ContextTimeoutEnabled: true,
	})

	if _, err := cache.Ping(context.Background()).Result(); err != nil {
		log.GetLoggerInstance().Fatalf("Could not connect to Redis: %v", err)
	}

	return cache
}
