package main

import (
	"fmt"
	"os"
)

func errExist(str string) {
	fmt.Println(str)
	os.Exit(1)
}

func testStringValue() {

	res := EmptyResource()
	res.Set("key1", "value1")

	value1, err := res.Get("key1").String()
	if err != nil {
		errExist(fmt.Sprintf("res.Get().String() failed: %v", err))
	}
	fmt.Println("key1:", value1)

}

func testMap() {

	res := EmptyResource()

	value2 := make(map[string]interface{})
	value2["subkey1"] = "subvalue1"

	res.Set("key2", value2)

	v2, err := res.Get("key2").Map()
	if err != nil {
		errExist(fmt.Sprintf("res.Get().Map() failed: %v", err))
	}
	fmt.Println("key2:", v2)
}

func testArray() {
	var err error

	res := EmptyResource()

	v := make([]interface{}, 0)
	v = append(v, "value1")

	res.Set("key1", v)

	v, err = res.Get("key1").Array()
	if err != nil {
		errExist(fmt.Sprintf("res.Get().Array() failed: %v", err))
	}
	fmt.Println("key2:", v)
}

func main() {
	testStringValue()
	testMap()
	testArray()
}
