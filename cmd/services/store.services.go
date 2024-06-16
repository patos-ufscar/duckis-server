package services

import "time"

type StoreService interface {
	Set(key string, val interface{})
	SetEx(key string, val interface{}, ttl time.Duration)
	Get(key string)											(*interface{}, error)
}
