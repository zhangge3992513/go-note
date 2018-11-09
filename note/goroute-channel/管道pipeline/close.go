package main

import (
	"fmt"
	"time"
)

func main() {
	example3()
}

func example3() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x < 5; x++ {
			naturals <- x
		}
		//close(naturals)
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
