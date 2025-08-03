package lib2

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLib2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lib2 Suite")
}
