package cache

import (
	"context"
	"golang.org/x/sync/singleflight"
	"log"
	"sync"
)

type Cache interface {
	GetWithFn(ctx context.Context, key string, fn func() (interface{}, error)) (i interface{}, err error)
}

type cache struct {
	sf singleflight.Group

	sync.Mutex
	db CacheDB
}

func New() Cache {
	return &cache{
		db: NewLRUCacheDB(),
	}
}

func (c *cache) GetWithFn(ctx context.Context, key string, fn func() (interface{}, error)) (i interface{}, err error) {
	if c == nil {
		i, err = fn()
		return
	}

	c.Lock()
	result, dbErr, ok := c.db.Get(key)
	c.Unlock()
	if ok {
		return result, nil
	}
	if dbErr != nil {
		err = dbErr
		return
	}

	sharded := false
	i, err, sharded = c.sf.Do(key, func() (interface{}, error) {
		v, err := fn()
		if err == nil {
			c.Lock()
			_ = c.db.Add(key, v)
			c.Unlock()
		}
		return v, nil
	})
	if sharded {
		log.Printf("key=%s is duplicated", key)
	}

	return
}
