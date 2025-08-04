package main

import (
	"testing"

	"github.com/omargallob/mono-repo-release/test"
)

func TestApp(t *testing.T) {
	test.RegisterAndRunSpecs(t, "App Suite")
}
