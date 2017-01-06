package main

import (
	"fmt"
	"reflect"
)

func main() {
	//name := "Aman Kapoor"
	age := 21
	//date := "18th December, 2016 3.37am"
	agep := &age

	fmt.Println("Memory Address of age is", &age, "and value of age is", *agep)
	fmt.Println("Type of agep is", reflect.TypeOf(agep))
	fmt.Println("Memory Address of agep is", agep)
}
