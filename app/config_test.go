package app_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ildarusmanov/go-up/app"
)

var _ = Describe("Config", func() {
	var (
		c *Config
	)

	BeforeEach(func() {
		c = NewConfig()
	})

	It("should create new config instance", func() {
		Expect(c).NotTo(BeNil())
	})

	Describe("methods", func() {
		Describe(".Get(), .Set()", func() {
			var (
				key   string
				value string
			)

			Context("with empty config", func() {

				It(".Get() should not return a value", func() {
					_, ok := c.Get(key)

					Expect(ok).To(BeFalse())
				})

				It(".Set() should set a value", func() {
					c.Set(key, value)
					v, ok := c.Get(key)

					Expect(ok).To(BeTrue())
					Expect(value).To(Equal(v))
				})
			})
		})
	})
})
