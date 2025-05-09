package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mojtabamovahedi/otp/config"
	rdis "github.com/redis/go-redis/v9"
)

var (
	// ErrRedisNotFound is returned when the OTP is not found in Redis
	ErrRedisNotFound = errors.New("OTP not found in Redis")
)

type redisConnection struct {
	client *rdis.Client
}

func NewRedisConnection(cfg config.RedisConfig) Provider {
	client := rdis.NewClient(&rdis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: "", // no password set
		DB:       cfg.DB,
	})

	return &redisConnection{
		client: client,
	}
}

func (r *redisConnection) Close() error {
	return r.client.Close()
}

func (r *redisConnection) Set(ctx context.Context, key string, value []byte) error {
	return r.client.Set(ctx, key, value, 5*time.Minute).Err()
}

func (r *redisConnection) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, rdis.Nil) {
			return nil, ErrRedisNotFound
		}
		return nil, err
	}
	return val, nil
}

func (r *redisConnection) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
