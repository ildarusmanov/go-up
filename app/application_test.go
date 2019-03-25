package app_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ildarusmanov/go-up/app"
	. "github.com/ildarusmanov/go-up/app"
)

var _ = Describe("Application", func() {
	var (
		a *Application
	)

	BeforeEach(func() {
		a = NewApplication(context.Background(), nil, nil)
	})

	It("should create new application", func() {
		Expect(a).NotTo(BeNil())
	})

	Describe("methods", func() {
		Describe(".AddServiceFactory(), .GetService()", func() {
			var (
				name    string
				factory func(context.Context) (app.Service, error)
			)

			BeforeEach(func() {
				name = "service_name"
				factory = func(ctx context.Context) (app.Service, error) {
					return func(v string) string {
						return v
					}, nil
				}
			})

			It("should add service", func() {
				err := a.AddServiceFactory(name, factory)
				v := "123123123"
				s, getErr := a.GetService(name)

				Expect(err).To(BeNil())
				Expect(getErr).To(BeNil())
				Expect(s.(func(string) string)(v)).To(Equal(v))
			})
		})

		Describe(".SetConfig(), .GetConfig()", func() {
			var (
				key, val string
			)

			BeforeEach(func() {
				key = "key123"
				val = "val44234323"
			})

			Context("with value", func() {
				It("should set and get value", func() {
					a.SetConfig(key, val)
					retV, retOk := a.GetConfig(key)

					Expect(retV).To(Equal(val))
					Expect(retOk).To(BeTrue())
				})
			})

			Context("without value", func() {
				It("should not found non-existing value", func() {
					retV, retOk := a.GetConfig(key)

					Expect(retOk).To(BeFalse())
					Expect(retV).To(Equal(""))
				})
			})
		})
	})
})
