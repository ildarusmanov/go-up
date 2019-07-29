package goup_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ildarusmanov/go-up/goup"
)

var (
	appLog = ""
)

func StartApplication(ctx context.Context) StopApplicationHandler {
	appLog = "started"

	return func() {
		appLog = "finished"
	}
}

var _ = Describe("Application", func() {
	var (
		ctx context.Context
	)

	BeforeEach(func() {
		ctx = context.WithValue(context.Background(), "k", "v")
	})

	It("should start and stop application", func() {
		Expect(appLog).To(Equal(""))

		stop := StartApplication(ctx)

		Expect(appLog).To(Equal("started"))

		stop()

		Expect(appLog).To(Equal("finished"))
	})
})
