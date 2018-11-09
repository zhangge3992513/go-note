package main

import "fmt"

func main() {
	var naturals = make(chan int)
	var squares = make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

/**
 * out: 出去的通道, 发送 x*x 结果
 * in: 进入的通道, 接收 x
 */
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
