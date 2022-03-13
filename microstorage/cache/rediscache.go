package cache

import (
	"context"
	"github.com/go-redis/redis"
	"microstorage"
	"strconv"
	"time"
)

func RedisCacheSetup(config microstorage.TCache, cache *SCache) {
	db, _ := strconv.ParseInt(config.DB, 10, 16)
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Pass,
		DB:       int(db),
	})
	var ctx = context.Background()
	var ttl = time.Duration(int64(time.Millisecond) * config.TTL)
	cache.Get = func(key string) (result CachedObject, error error) {
		val, err := rdb.Get(ctx, key).Bytes()
		if err != nil {
			return result, err
		}
		rdb.Expire(ctx, key, ttl)
		return transformToJson(val), nil
	}

	cache.Put = func(key string, data CachedObject) {
		rdb.Set(ctx, key, transformFromJson(data), ttl)
	}

	cache.Exists = func(key string) bool {
		i, e := rdb.Exists(ctx, key).Result()
		return (i > 0) && (e == nil)
	}

	cache.Delete = func(key string) {

	}

}
