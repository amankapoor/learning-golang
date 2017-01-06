package main

import (
	"fmt"
	//"strings"
)

func main() {
	uID := displayUID(1,2,3,4,5,6)
	fmt.Println(uID)
}

func displayUID(userids ...int) int {
	
	userid := userids[0]
	for _, i := range userids {
		if i>userid {
			userid = i
		}
	}
	
	return userid
} 