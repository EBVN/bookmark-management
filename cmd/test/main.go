package main

import (
	"context"
	"time"

	"github.com/EBVN/bookmark-management/pkg/redis"
)

func main() {
	ctx := context.Background()

	rdb, _ := redis.NewClient("")

	rdb.Set(ctx, "123", "google.com", 1*time.Hour)

	cacheRedisDB, _ := redis.NewClient("CACHE")
	cacheRedisDB.Set(ctx, "hjk", "oiyasxciozuxcio", 1*time.Hour)
}
