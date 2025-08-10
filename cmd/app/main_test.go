package main

import (
	"bytes"
	"os"

	"github.com/omargallob/mono-repo-release/pkg/lib1"
	"github.com/omargallob/mono-repo-release/pkg/lib2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("main", func() {
	It("prints greet and farewell messages to stdout", func() {
		r, w, _ := os.Pipe()
		old := os.Stdout
		os.Stdout = w
		main()
		w.Close()
		os.Stdout = old
		var buf bytes.Buffer
		buf.ReadFrom(r)
		output := buf.String()
		Expect(output).To(ContainSubstring(lib1.Greet()))
		Expect(output).To(ContainSubstring(lib2.Farewell()))
	})

	It("returns non-empty greet and farewell messages", func() {
		Expect(lib1.Greet()).NotTo(BeEmpty())
		Expect(lib2.Farewell()).NotTo(BeEmpty())
	})
})
