// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "abc"
	strPointer := reflect.ValueOf(&str)
	fmt.Printf("strPointer type %v\n", strPointer.Type())

}
