package test

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

// RegisterAndRunSpecs sets up Ginkgo and Gomega for a test suite.
func RegisterAndRunSpecs(t *testing.T, suiteName string) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, suiteName)
}
