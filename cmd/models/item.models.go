package models

type StoreItem interface {
	Get()					(interface{}, error)
	GetUsage()				uint32
}
