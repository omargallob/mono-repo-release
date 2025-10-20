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

	// New test for AsciiArt function
	It("returns the correct ASCII art", func() {
		expectedArt := `
   _____
  /     \
 | () () |
  \  ^  /
   |||||
   |||||
`
		Expect(lib2.AsciiArt()).To(Equal(expectedArt))
	})

	// New test for FarewellWithArt function
	It("returns the correct farewell message with ASCII art", func() {
		expectedMessage := "Good bye from lib2\n" + `
   _____
  /     \
 | () () |
  \  ^  /
   |||||
   |||||
`
		Expect(lib2.FarewellWithArt()).To(Equal(expectedMessage))
	})
	// New test for Header function
	It("returns the correct header message", func() {
		expectedHeader := lib2.Header()
		Expect(expectedHeader).To(ContainSubstring("Welcome to lib2!"))
	})

	// New test for FarewellWithHeaderAndArt function
	It("returns the correct farewell message with header and ASCII art", func() {
		expectedMessage := lib2.Header() + "\n" + "Good bye from lib2\n" + `
   _____
  /     \
 | () () |
  \  ^  /
   |||||
   |||||
`
		Expect(lib2.FarewellWithHeaderAndArt()).To(Equal(expectedMessage))
	})
})
