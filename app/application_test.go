package app_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

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
})
