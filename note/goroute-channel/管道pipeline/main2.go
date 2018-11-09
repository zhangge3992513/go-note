package main

import (
	"fmt"
	"time"
)

func main() {
	example1()
	/*	go example1()
		time.Sleep(time.Second * 5)*/

}

/**
 * close, 捕获错误
 */
func example1() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				errMsg := fmt.Errorf("%v", err)
				fmt.Println(errMsg)
			}
		}()
		for x := 0; ; x++ {
			naturals <- x
		}

	}()

	//squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
	}()
	for {
		temp := <-squares
		if temp > 2 {
			close(naturals)
			close(squares)
			//break
		}
		fmt.Println(temp)
		time.Sleep(time.Second)
	}
}
