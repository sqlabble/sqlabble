package sqlabble

import (
	"fmt"
	"os"
)

func main() {
	if err := Generate(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
