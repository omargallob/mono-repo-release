package lib1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Greet", func() {
	It("returns the correct greeting message", func() {
		Expect(Greet()).To(Equal("Hello from lib1"))
	})
})
