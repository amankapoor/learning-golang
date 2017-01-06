package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	go func() {
		defer waitgroup.Done()
		time.Sleep(5*time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitgroup.Done()
		fmt.Println("World")
	}()
	waitgroup.Wait()
}