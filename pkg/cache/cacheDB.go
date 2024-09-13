package cache

import "github.com/golang/groupcache/lru"

const (
	defaultMaxEntries = 3
)

type CacheDB interface {
	Get(key interface{}) (interface{}, error, bool)
	Add(key interface{}, value interface{}) error
}

type lruCacheDB struct {
	lru *lru.Cache
}

func NewLRUCacheDB() CacheDB {
	return &lruCacheDB{
		lru: lru.New(defaultMaxEntries),
	}
}

func (lc *lruCacheDB) Get(key interface{}) (i interface{}, err error, ok bool) {
	i, ok = lc.lru.Get(key)
	return
}

func (lc *lruCacheDB) Add(key interface{}, value interface{}) error {
	lc.lru.Add(key, value)
	return nil
}

type mockCacheDB struct{}

func NewMockCacheDB() CacheDB {
	return nil
}

func (mc *mockCacheDB) Get(key interface{}) (i interface{}, err error, ok bool) {
	return
}

func (mc *mockCacheDB) Add(key interface{}, value interface{}) error {
	return nil
}
