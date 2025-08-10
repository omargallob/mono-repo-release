package lib2_test

import (
	"github.com/omargallob/mono-repo-release/pkg/lib2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Farewell", func() {
	It("returns the correct farewell message", func() {
		Expect(lib2.Farewell()).To(Equal("Good bye from lib2"))
	})
})
