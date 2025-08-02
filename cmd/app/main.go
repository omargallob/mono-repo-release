package app

import (
	"fmt"

	"github.com/omargallob/mono-repo-release/pkg/lib1"
	"github.com/omargallob/mono-repo-release/pkg/lib2"
)

func main() {
	fmt.Println(lib1.Greet())
	fmt.Println(lib2.Farewell())
}
