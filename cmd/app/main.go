package main

import (
	"fmt"

	"github.com/omargallob/mono-repo-release/pkg/lib1"
	"github.com/omargallob/mono-repo-release/pkg/lib2"
)

// version is set at build time via ldflags. Default to "dev" for local builds.
var version = "dev"

// main is the entry point for the application.
func main() {
	fmt.Println(lib1.Greet())
	fmt.Println(lib2.Farewell())
}
