package models_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/patos-ufscar/duckis-server/models"
	"github.com/patos-ufscar/duckis-server/utils"
)

var _ = Describe("StoreItemExImpl", func() {

	Describe("Get", func() {
		Context("when item has not expired", func() {
			It("should return the stored value", func() {
				ttl := 1 * time.Second
				vals := []interface{}{
					"string",
					int32(-42),
					uint8(42),
				}
				stores := []models.StoreItem{}

				for _, v := range vals {
					stores = append(stores, models.NewStoreItemExImpl(v, ttl))
				}

				time.Sleep(2 * ttl)

				for i, v := range stores {
					Ω(v.Get()).Should(Equal(vals[i]))
				}
			})
		})

		Context("when item has expired", func() {
			It("should return the expired error", func() {
				ttl := 1 * time.Second
				vals := []interface{}{
					"string",
					int32(-42),
					uint8(42),
				}
				stores := []models.StoreItem{}

				for _, v := range vals {
					stores = append(stores, models.NewStoreItemExImpl(v, ttl))
				}

				time.Sleep(2 * ttl)

				for _, v := range stores {
					Ω(v.Get()).Should(MatchError(utils.ErrValueTimedOut))
				}
			})
		})
	})
})