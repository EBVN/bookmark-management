package repository

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

const expTime = 3600 * time.Second

type URLStorage interface {
	StoreURL(ctx context.Context, key, url string) error
	GetURL(ctx context.Context, key string) (string, error)
}

type urlStorage struct {
	c *redis.Client
}

func NewURLStorage(c *redis.Client) URLStorage {
	return &urlStorage{c: c}
}

func (s *urlStorage) StoreURL(ctx context.Context, key, url string) error {
	return s.c.Set(ctx, key, url, expTime).Err()
}

var ErrKeyNotFound = errors.New("key not found")

func (s *urlStorage) GetURL(ctx context.Context, key string) (string, error) {
	url, err := s.c.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrKeyNotFound
		}

		return "", err
	}

	return url, nil
}
