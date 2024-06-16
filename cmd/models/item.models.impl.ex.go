package models

import (
	"time"
	"unsafe"

	"github.com/patos-ufscar/duckis-server/utils"
)

type StoreItemExImpl struct {
	value 			interface{}
	ex				time.Time
}

func NewStoreItemExImpl(value interface{}, ttl time.Duration) StoreItem {
	return &StoreItemExImpl{
		value: value,
		ex: time.Now().Add(ttl),
	}
}

func (s *StoreItemExImpl) Get() (interface{}, error) {

	if s.ex.Before(time.Now()) {
		return nil, utils.ErrValueTimedOut
	}

	return s.value, nil
}

func (s *StoreItemExImpl) GetUsage() uint32 {
	return uint32(unsafe.Sizeof(*s))
}
