package main 

import "fmt"

func main() {

	oldSlice := make([]int, 4, 8)
	fmt.Println("Old Slice is", oldSlice)
	newSlice := make([]int, 4, 8)
	
	for i:=0; i<4; i++ {
		newSlice[i] = i+1
	}
	
	oldSlice = append(oldSlice, newSlice...)
	fmt.Println("After Appending:", oldSlice)

}