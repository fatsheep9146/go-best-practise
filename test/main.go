package main

import (
	"fmt"
)

func main() {
	min := 10
	max := 19
	step := 10
	for index := min; index <= max; index += step {
		end := index + step - 1
		if end > max {
			end = max
		}
		fmt.Printf("range %d - %d\n", index, end)
	}
}
