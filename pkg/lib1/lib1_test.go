package lib1_test

import (
	"github.com/omargallob/mono-repo-release/pkg/lib1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Greet", func() {
	It("returns the correct greeting message", func() {
		Expect(lib1.Greet()).To(Equal("Hello from lib1"))
	})
})
