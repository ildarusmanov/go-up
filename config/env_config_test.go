package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ildarusmanov/go-up/config"
)

var _ = Describe("Config", func() {
	var (
		c *EnvConfig
	)

	BeforeEach(func() {
		c = NewEnvConfig()
	})

	It("should create new config instance", func() {
		Expect(c).NotTo(BeNil())
	})

	Describe("methods", func() {
		Describe(".RequireKeys()", func() {
			var (
				key   string
				value string
			)

			BeforeEach(func() {
				key = "somekey"
				value = "somevalue"

				c.Set(key, value)
			})

			AfterEach(func() {
				c.Unset(key)
			})

			It("should find existing key", func() {
				Expect(c.RequireKeys([]string{key})).To(BeNil())
				Expect(c.RequireKeys([]string{key, key + "1"})).NotTo(BeNil())
				Expect(c.RequireKeys([]string{key + "1"})).NotTo(BeNil())
			})
		})

		Describe(".Get(), .Set()", func() {
			var (
				key   string
				value string
			)

			Context("with key, value config", func() {
				BeforeEach(func() {
					key = "somekey"
					value = "somevalue"

					c.Set(key, value)
				})

				AfterEach(func() {
					c.Unset(key)
				})

				It(".Get() should not return a value", func() {
					v, ok := c.Get(key)

					Expect(ok).To(BeTrue())
					Expect(v).To(Equal(value))
				})

				It(".GetString() should not return a value", func() {
					v, ok := c.GetString(key)

					Expect(ok).To(BeTrue())
					Expect(v).To(Equal(value))
				})

				It(".Set() should set a value", func() {
					c.Set(key, value)
					v, ok := c.Get(key)

					Expect(ok).To(BeTrue())
					Expect(v).To(Equal(value))
				})
			})

			Context("with empty config", func() {
				It(".Get() should not return a value", func() {
					_, ok := c.Get(key)

					Expect(ok).To(BeFalse())
				})

				It(".GetString() should not return a value", func() {
					_, ok := c.Get(key)

					Expect(ok).To(BeFalse())
				})

				It(".Set() should set a value", func() {
					c.Set(key, value)
					v, ok := c.Get(key)

					Expect(ok).To(BeTrue())
					Expect(v).To(Equal(value))
				})
			})
		})
	})
})
