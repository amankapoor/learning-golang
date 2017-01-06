package main

import (
	"fmt"
	"strings"
)

func main() {
	welcome := "wElcOmE"
	fullname := "aman kapoor"

	fmt.Println(beautify(welcome, fullname))
}

func beautify (welcome, fullname string) (string, string) {
	welcome = strings.Title(welcome[:1])+strings.ToLower(welcome[1:])
	fullname = strings.Title(fullname[:1])+strings.ToLower(fullname[1:])

	return welcome, fullname
}