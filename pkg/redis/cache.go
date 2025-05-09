package redis

import (
	"context"
)

type Provider interface {
	Set(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
	Del(ctx context.Context, key string) error
}

type ObjectCacher[T any] struct {
	provider Provider
}

const KeyPrefix = "USER_PHONE_NUMBER"

func CreateKey(k string) string {
	return KeyPrefix + "." + k
}

func NewObjectCacher[T any](p Provider) *ObjectCacher[T] {
	return &ObjectCacher[T]{
		provider: p,
	}
}

func (c *ObjectCacher[T]) Set(ctx context.Context, key string, v T) error {
	data, err := c.Marshal(v)
	if err != nil {
		return err
	}
	return c.provider.Set(ctx, CreateKey(key), data)
}

func (c *ObjectCacher[T]) Get(ctx context.Context, key string) (T, error) {
	var t T
	data, err := c.provider.Get(ctx, CreateKey(key))
	if err != nil {
		// if errors.Is(err, ErrRedisNotFound) {
		// 	return t, nil
		// }
		return t, err
	}
	return t, c.Unmarshal(data, &t)
}

func (c *ObjectCacher[T]) Del(ctx context.Context, key string) error {
	return c.provider.Del(ctx, CreateKey(key))
}
