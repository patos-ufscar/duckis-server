package services

import (
	"regexp"
	"time"
)

type StoreService interface {
	// Set sets the value for key, overwrites any existing previous val :: O(1)
	Set(key string, val interface{})

	// SetEx sets the value for key with the ttl, after ttl the key will not be returned :: O(1)
	SetEx(key string, val interface{}, ttl time.Duration)

	// Get gets the value previously set in the key :: O(1)
	Get(key string)											(*interface{}, error)
	
	// Search searches the complete db for any value that matches the regEx pattern :: O(n)
	Search(reg *regexp.Regexp)								*[]interface{}
}
