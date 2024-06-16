package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/patos-ufscar/duckis-server/models"
)

var _ = Describe("StoreItemStdImpl", func() {

	Describe("Get", func() {
		Context("when item has not expired", func() {
			It("should return the stored value", func() {
				vals := []interface{}{
					"string",
					int32(-42),
					uint8(42),
				}
				stores := []models.StoreItem{}

				for _, v := range vals {
					stores = append(stores, models.NewStoreItemStdImpl(v))
				}

				for i, v := range stores {
					Ω(v.Get()).Should(Equal(vals[i]))
				}
			})
		})
	})
})