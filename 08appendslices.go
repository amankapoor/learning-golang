package main 

import "fmt"

func main() {

	oldSlice := []int{1,2,3,4}
	fmt.Println("Old Slice is", oldSlice)
	newSlice := []int{5,6,7,8}
	fmt.Println("New Slice is", newSlice)

	oldSlice = append(oldSlice, newSlice...)
	fmt.Println("After Appending:", oldSlice)

}