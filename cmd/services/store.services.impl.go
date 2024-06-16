package services

import (
	"regexp"
	"sync"
	"time"

	"github.com/patos-ufscar/duckis-server/models"
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
	defer s.mu.Unlock()
	
	s.cache[key] = models.NewStoreItemStdImpl(val)
}

func (s *StoreServiceImpl) SetEx(key string, val interface{}, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.cache[key] = models.NewStoreItemExImpl(val, ttl)
}

func (s *StoreServiceImpl) Get(key string) (*interface{}, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	storeItem, ok := s.cache[key]

	if !ok {
		return nil, ErrKeyNotPresent
	}

	val, err := storeItem.Get()
	if err != nil {
		if err == models.ErrValueTimedOut {
			delete(s.cache, key)
			return nil, ErrKeyNotPresent
		}

		return nil, err
	}

	return &val, nil
}

func (s *StoreServiceImpl) Search(reg *regexp.Regexp) *[]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	vals := []interface{}{}
	for k, v := range s.cache {
		if !reg.MatchString(k) {
			continue
		}

		val, err := v.Get()
		if err != nil {
			if err == models.ErrValueTimedOut {
				delete(s.cache, k)
			}
			continue
		}

		vals = append(vals, val)
	}

	return &vals
}