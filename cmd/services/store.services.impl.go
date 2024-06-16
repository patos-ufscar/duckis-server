package services

import (
	"sync"
	"time"

	"github.com/patos-ufscar/duckis-server/models"
	"github.com/patos-ufscar/duckis-server/utils"
)

type StoreServiceImpl struct {
	cache					map[string]models.StoreItem
	mu						sync.RWMutex
}

func NewStoreServiceImpl() StoreService {
	return &StoreServiceImpl{
		cache: make(map[string]models.StoreItem),
		mu: sync.RWMutex{},
	}
}

func (s *StoreServiceImpl) Set(key string, val interface{}) {
	s.mu.Lock()
	s.cache[key] = models.NewStoreItemStdImpl(val)
	s.mu.Unlock()
}

func (s *StoreServiceImpl) SetEx(key string, val interface{}, ttl time.Duration) {
	s.mu.Lock()
	s.cache[key] = models.NewStoreItemExImpl(val, ttl)
	s.mu.Unlock()
}

func (s *StoreServiceImpl) Get(key string) (*interface{}, error) {
	s.mu.RLock()
	storeItem, ok := s.cache[key]
	s.mu.RUnlock()

	if !ok {
		return nil, utils.ErrKeyNotPresent
	}

	val, err := storeItem.Get()
	if err != nil {
		if err == utils.ErrValueTimedOut {
			delete(s.cache, key)
		}

		return nil, err
	}

	return &val, nil
}