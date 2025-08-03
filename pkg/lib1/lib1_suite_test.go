package lib1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLib1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lib1 Suite")
}
