package lib2

import (
	"testing"

	"github.com/omargallob/mono-repo-release/test"
)

func TestLib2(t *testing.T) {
	test.RegisterAndRunSpecs(t, "Lib2 Suite")
}
