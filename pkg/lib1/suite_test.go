package lib1_test

import (
	"testing"

	"github.com/omargallob/mono-repo-release/test"
)

func TestLib1(t *testing.T) {
	test.RegisterAndRunSpecs(t, "Lib1 Suite")
}
