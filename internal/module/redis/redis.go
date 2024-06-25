package redis

import (
	"affiliate/internal/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/logrusorgru/aurora"
)

var store *redis.Client

// InitRedis ...
func InitRedis() {
	ctx := context.Background()
	cfg := config.GetENV().Redis

	store = redis.NewClient(&redis.Options{
		Addr:     cfg.URI,
		Password: cfg.Password,
		DB:       0, // use default DB
	})

	// Test
	_, err := store.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Cannot connect to redis", cfg.URI, err)
	}

	fmt.Println(aurora.Green("*** CONNECTED TO REDIS: " + cfg.URI))
}

// GetClient ...
func GetClient() *redis.Client {
	return store
}

// SetKeyValue ...
func SetKeyValue(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	storeByte, _ := json.Marshal(value)
	r := store.Set(ctx, key, storeByte, expiration)
	return r.Err()
}

// GetValueByKey ...
func GetValueByKey(key string) string {
	ctx := context.Background()
	value, _ := store.Get(ctx, key).Result()
	return value
}

// GetJSON ...
func GetJSON(key string, result interface{}) (ok bool) {
	v := GetValueByKey(key)
	if v == "" {
		return false
	}
	if err := json.Unmarshal([]byte(v), result); err != nil {
		return false
	}
	return true
}

// DelKey ...
func DelKey(key string) error {
	ctx := context.Background()
	return store.Del(ctx, key).Err()
}

// DelAllKeyByPattern ...
func DelAllKeyByPattern(pattern string) error {
	ctx := context.Background()
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = store.Scan(ctx, cursor, pattern, 100000).Result()
		if err != nil {
			return err
		}

		// If found any keys, delete
		if len(keys) > 0 {
			if err := store.Del(ctx, keys...).Err(); err != nil {
				return err
			}
		}

		if cursor == 0 {
			break
		}
	}
	return nil
}

// NewMutex ...
func NewMutex(name string) *redsync.Mutex {
	pool := goredis.NewPool(store)

	rs := redsync.New(pool)

	opts := []redsync.Option{
		redsync.WithRetryDelay(time.Millisecond * 100), // retry after 100 ms
		redsync.WithTries(10),                          // try 10 times
	}

	return rs.NewMutex(name, opts...)
}
