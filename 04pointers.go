package main

import (
"fmt"
"reflect"
)

func main() {
	age := 21

	fmt.Println("Age is:", age, "type of age is", reflect.TypeOf(age))

	changeAge(&age)

	fmt.Println("New age is:", age)
}

func changeAge(a *int) int {
	*a = *a+10
	return *a
}