package main

import (
	"fmt"
)

// errcheck case
func caseErrcheck() {
	fn := func() error {
		fmt.Println("case errorcheck do something")
		return fmt.Errorf("test error")
	}

	// wrong
	// fn()

	// correct
	err := fn()
	fmt.Println("got err", err)
}

func main() {
	caseErrcheck()
}
