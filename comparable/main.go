package main

import (
	"fmt"
)

/*
  test about the comparability of different types.
*/

func main() {
	var b1, b2 = true, true
	fmt.Printf("the compare between [Boolean] b1 == b2: %v\n", b1 == b2)

	var ar1, ar2 = [2]string{"e1", "e2"}, [2]string{"e1", "e2"}
	fmt.Printf("the compare between [StringArray] b1 == b2: %v\n", ar1 == ar2)

	type testArrayStruct struct {
		name  string
		books [2]string
	}
	var ts1, ts2 = testArrayStruct{name: "n1", books: [2]string{"e1", "e2"}}, testArrayStruct{name: "n1", books: [2]string{"e1", "e2"}}
	fmt.Printf("the compare between [testArrayStruct] ts1 == ts2: %v\n", ts1 == ts2)

	type testSliceStruct struct {
		name  string
		books *[]string
	}
	var books []string = []string{"e1", "e2"}
	var tss1, tss2 = testSliceStruct{name: "n1", books: &books}, testSliceStruct{name: "n1", books: &books}
	fmt.Printf("the compare between [testArrayStruct] tss1 == tss2: %v\n", tss1 == tss2)

	type testInterfaceStruct struct {
		name  string
		books interface{}
	}
	var ti1, ti2 = testInterfaceStruct{name: "n1", books: &books}, testInterfaceStruct{name: "n1", books: &books}
	fmt.Printf("the compare between [testArrayStruct] ti1 == ti2: %v\n", ti1 == ti2)
}
