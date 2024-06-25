package redis

import (
	"errors"
	"fmt"

	"github.com/go-redsync/redsync/v4"
)

// RedisInterface ...
type RedisInterface interface {
	LockFlowTopup(key string) (*redsync.Mutex, error)
}

type redisImpl struct {
}

func (r redisImpl) LockFlowTopup(key string) (*redsync.Mutex, error) {
	mu := NewMutex(fmt.Sprintf("topup_flow_%s", key))
	if err := mu.Lock(); err != nil {
		if errors.Is(err, redsync.ErrFailed) {
			return r.LockFlowTopup(key)
		}
		return nil, err
	}
	return mu, nil
}

// Redis ...
func Redis() RedisInterface {
	return redisImpl{}
}
