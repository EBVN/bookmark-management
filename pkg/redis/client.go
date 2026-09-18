package redis

import "github.com/redis/go-redis/v9"

// NewClient create new redis client
func NewClient(envPrefix string) (*redis.Client, error) {
	cfg, err := newConfig(envPrefix)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	return rdb, nil
}
