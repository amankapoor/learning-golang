package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(rand.Seed(10))
}