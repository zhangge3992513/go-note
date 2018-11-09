package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	//squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()
	for {
		temp, ok := <-squares
		fmt.Println(temp, ok)
		time.Sleep(time.Second)
	}
}
