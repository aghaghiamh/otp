package implementation

import (
	"context"
	"fmt"
	"otp/src/pkg/config"
	"otp/src/repo"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisOTPRepository struct {
    client *redis.Client
}

func NewRedisOTPRepository(client *redis.Client) repo.OTPManagement {
    return &redisOTPRepository{client: client}
}

func (r *redisOTPRepository) Store(ctx context.Context, mobileNumber, otpHash string) error {
	cnf := config.GetAppConfigInstance()
	key := genKey(mobileNumber)
	expirationTime := time.Duration(cnf.DefaultExpirationInMinute) * time.Minute

	return r.client.Set(ctx, key, otpHash, expirationTime).Err()
}

func (r *redisOTPRepository) Get(ctx context.Context, mobileNumber string) (string, error) {
	key := genKey(mobileNumber)
	return r.client.Get(ctx, key).Result()
}

func genKey(mobileNumber string) string {
	cnf := config.GetAppConfigInstance()
	return fmt.Sprintf("%s:%s",cnf.OTPCacheReservedKey, mobileNumber)
}