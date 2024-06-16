package models

import (
	"unsafe"
)

type StoreItemStdImpl struct {
	value 			interface{}
}

func NewStoreItemStdImpl(value interface{}) StoreItem {
	return &StoreItemStdImpl{
		value: value,
	}
}

func (s *StoreItemStdImpl) Get() (interface{}, error) {
	return s.value, nil
}

func (s *StoreItemStdImpl) GetUsage() uint32 {
	return uint32(unsafe.Sizeof(*s))
}
