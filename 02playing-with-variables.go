package main

import (
	"fmt"
	"reflect" //this package is for seeing internal data
)

func main() {
	var a = 10.0000000
	var b = 3

	fmt.Println("a is of type", reflect.TypeOf(a), "and b is of type", reflect.TypeOf(b))

	//var c = a+b //this will not work because of type mismatch
	var c = int(a)+b

	fmt.Println("c is of type", reflect.TypeOf(c), "and its value is", c)
}