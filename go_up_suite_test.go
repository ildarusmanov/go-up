package go_up_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoUp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoUp Suite")
}
