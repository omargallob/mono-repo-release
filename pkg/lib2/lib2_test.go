package lib2

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Farewell", func() {
	It("returns the correct farewell message", func() {
		Expect(Farewell()).To(Equal("Good bye from lib2"))
	})
})
