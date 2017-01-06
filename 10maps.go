package main 

import "fmt"

func main() {

	newmap := map[string]int {"A":1, "B":2, "C":3} 

	for i, j := range newmap {
		fmt.Println("Key is", i, "and Value is", j)
	}
}