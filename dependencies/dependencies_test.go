package dependencies_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ildarusmanov/go-up/dependencies"
)

var _ = Describe("Dependencies", func() {
	var (
		d *Dependencies
	)

	BeforeEach(func() {
		d = NewDependencies()
	})

	It("should create new dependencies instance", func() {
		Expect(d).NotTo(BeNil())
	})

	Describe("methods", func() {
		Describe(".Add(), .Get()", func() {
			var (
				name    string
				service func(string) string
			)

			BeforeEach(func() {
				name = "hello"
				service = func(v string) string {
					return v
				}
			})

			It("should add new service", func() {
				d.Add(name, service)
				s, err := d.Get(name)

				Expect(err).To(BeNil())
				Expect(s.(func(string) string)("hello")).To(Equal("hello"))
			})
		})
	})
})
