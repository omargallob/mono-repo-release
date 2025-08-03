package main

import (
	"testing"

	"github.com/omargallob/mono-repo-release/pkg/lib1"
	"github.com/omargallob/mono-repo-release/pkg/lib2"
)

func TestMainOutput(t *testing.T) {
	greet := lib1.Greet()
	farewell := lib2.Farewell()

	if greet == "" {
		t.Errorf("Expected non-empty greet message")
	}
	if farewell == "" {
		t.Errorf("Expected non-empty farewell message")
	}
}
