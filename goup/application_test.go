package goup_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"
	. "github.com/ildarusmanov/go-up/goup"
	"github.com/ildarusmanov/go-up/test/mock_goup"
)

var _ = Describe("Application", func() {
	var (
		a            *Application
		mockCtrl     *gomock.Controller
		config       *mock_goup.MockConfigManager
		dependencies *mock_goup.MockDependenciesManager
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		config = mock_goup.NewMockConfigManager(mockCtrl)
		dependencies = mock_goup.NewMockDependenciesManager(mockCtrl)

		a = NewApplication(context.Background(), dependencies, config)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("should create new application", func() {
		Expect(a).NotTo(BeNil())
	})

	Describe("methods", func() {
		Describe(".AddServiceFactory(), .GetService()", func() {
			var (
				name    string
				factory func(context.Context) (interface{}, error)
				service string
			)

			BeforeEach(func() {
				name = "service_name"

				service = "service 123"

				factory = func(ctx context.Context) (interface{}, error) {
					return service, nil
				}
			})

			It("should add service", func() {
				dependencies.EXPECT().Add(name, service).AnyTimes()
				dependencies.EXPECT().Get(name).Return(service, nil).AnyTimes()

				err := a.AddServiceFactory(name, factory)
				s, getErr := a.GetService(name)

				Expect(err).To(BeNil())
				Expect(getErr).To(BeNil())
				Expect(s.(string)).To(Equal(service))
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
					config.EXPECT().Set(key, val).AnyTimes()
					config.EXPECT().Get(key).AnyTimes().Return(val, true)

					a.SetConfig(key, val)
					retV, retOk := a.GetConfig(key)

					Expect(retV.(string)).To(Equal(val))
					Expect(retOk).To(BeTrue())
				})
			})

			Context("without value", func() {
				It("should not found non-existing value", func() {
					config.EXPECT().Get(key).AnyTimes().Return(nil, false)

					_, retOk := a.GetConfig(key)

					Expect(retOk).To(BeFalse())
				})
			})
		})
	})
})
