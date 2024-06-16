package services_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/patos-ufscar/duckis-server/services"
)

var _ = Describe("StoreServiceImpl", func() {
	var storeService services.StoreService

	BeforeEach(func() {
		storeService = services.NewStoreServiceImpl()
	})

	Describe("Set->Get", func() {
		It("should return correctly", func() {
			keys := []string{
				"first",
				"second",
				"third",
			}
			vals := []interface{}{
				"string",
				int32(-42),
				uint8(42),
			}
			for i, v := range keys {
				storeService.Set(v, vals[i])
			}
			for i, v := range keys {
				got, err := storeService.Get(v)
				Ω(err).NotTo(HaveOccurred())
				Ω(*got).To(Equal(vals[i]))
			}
		})

		It("should err on empty vals", func() {
			keys := []string{
				"first",
				"second",
				"third",
			}
			for _, v := range keys {
				_, err := storeService.Get(v)
				Ω(err).Should(MatchError(services.ErrKeyNotPresent))
			}
		})
	})

	Describe("SetEx->Get", func() {
		It("should return correctly when inside ttl", func() {
			keys := []string{
				"first",
				"second",
				"third",
			}
			vals := []interface{}{
				"string",
				int32(-42),
				uint8(42),
			}
			ttl := 1 * time.Second
			for i, v := range keys {
				storeService.SetEx(v, vals[i], ttl)
			}
			for i, v := range keys {
				got, err := storeService.Get(v)
				Ω(err).NotTo(HaveOccurred())
				Ω(*got).To(Equal(vals[i]))
			}
		})

		It("should not be present when outside ttl", func() {
			keys := []string{
				"first",
				"second",
				"third",
			}
			vals := []interface{}{
				"string",
				int32(-42),
				uint8(42),
			}
			ttl := 1 * time.Second
			for i, v := range keys {
				storeService.SetEx(v, vals[i], ttl)
			}
			time.Sleep(2 * ttl)
			for _, v := range keys {
				_, err := storeService.Get(v)
				Ω(err).To(MatchError(services.ErrKeyNotPresent))
			}
		})

		It("should err on empty vals", func() {
			keys := []string{
				"first",
				"second",
				"third",
			}
			for _, v := range keys {
				_, err := storeService.Get(v)
				Ω(err).Should(MatchError(services.ErrKeyNotPresent))
			}
		})
	})
})