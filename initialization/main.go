// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

var variableTime = time.Now()
var variableInt = 1

func init() {
	fmt.Printf("init time: %v\n", time.Now())
	fmt.Printf("varible int: %v\n", variableInt)
}

func main() {
	fmt.Printf("variable time: %v\n", variableTime)
	fmt.Println("now time:", time.Now())
}
