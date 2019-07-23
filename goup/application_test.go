package goup_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"
	. "github.com/ildarusmanov/go-up/goup"
	"github.com/ildarusmanov/go-up/test/mock_goup"
)

var _ = Describe("Application", func() {
	var (
		a        *Application
		ctx      context.Context
		mockCtrl *gomock.Controller
		config   *mock_goup.MockConfigManager
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		config = mock_goup.NewMockConfigManager(mockCtrl)
		ctx = context.WithValue(context.Background(), "k", "v")
		a = NewApplication().WithConfig(config).WithContext(ctx)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("should create new application", func() {
		Expect(a).NotTo(BeNil())
	})

	Describe("methods", func() {
		Describe(".RequireConfig(), .SetConfig(), .UnsetConfg(), .GetConfig(), .GetConfigString()", func() {
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
					config.EXPECT().GetString(key).AnyTimes().Return(val, true)
					config.EXPECT().Unset(key).AnyTimes().Return(nil)
					config.EXPECT().RequireKeys([]string{key}).AnyTimes().Return(nil)

					a.SetConfig(key, val)
					retV, retOk := a.GetConfig(key)
					retStrV, retStrOk := a.GetConfigString(key)
					reqErr := a.RequireConfig([]string{key})
					unsetErr := a.UnsetConfig(key)

					Expect(retV.(string)).To(Equal(val))
					Expect(retOk).To(BeTrue())
					Expect(retStrV).To(Equal(val))
					Expect(retStrOk).To(BeTrue())
					Expect(reqErr).To(BeNil())
					Expect(unsetErr).To(BeNil())
				})
			})

			Context("without value", func() {
				It("should not found non-existing value", func() {
					config.EXPECT().Get(key).AnyTimes().Return(nil, false)
					config.EXPECT().GetString(key).AnyTimes().Return("", false)
					config.EXPECT().Unset(key).AnyTimes().Return(errors.New("err"))
					config.EXPECT().RequireKeys([]string{key}).AnyTimes().Return(errors.New("err"))

					_, retOk := a.GetConfig(key)
					retStrV, retStrOk := a.GetConfigString(key)
					reqErr := a.RequireConfig([]string{key})
					unsetErr := a.UnsetConfig(key)

					Expect(retOk).To(BeFalse())
					Expect(retStrV).To(Equal(""))
					Expect(retStrOk).To(BeFalse())
					Expect(reqErr).NotTo(BeNil())
					Expect(unsetErr).NotTo(BeNil())
				})
			})
		})

		Describe(".WithContext(), .Context()", func() {
			It("should change context", func() {
				v := a.WithContext(context.WithValue(ctx, "k1", "v1")).Context().Value("k1").(string)

				Expect(v).To(Equal("v1"))
			})
		})
	})
})
