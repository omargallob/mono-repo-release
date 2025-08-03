package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/omargallob/mono-repo-release/pkg/lib1"
	"github.com/omargallob/mono-repo-release/pkg/lib2"
)

// add a test for the runapp function
func TestMain(t *testing.T) {
	// Capture stdout
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// run the main entry point
	main()

	w.Close()
	os.Stdout = old
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, lib1.Greet()) {
		t.Errorf("Expected output to contain greet message: %q", lib1.Greet())
	}
	if !strings.Contains(output, lib2.Farewell()) {
		t.Errorf("Expected output to contain farewell message: %q", lib2.Farewell())
	}
}
