package main

import (
	"fmt"
	"time"
)

func main() {
	//example1()
	//example2()
	example3()
}

/**
 * 无缓冲通道, 发送方发送操作会被阻塞
 */
func example1() {
	var ch1 chan int
	ch1 = make(chan int)
	go func() {
		ch1 <- 1
		fmt.Println("发送完毕")
	}()

	time.Sleep(time.Second)
	fmt.Println("main end")
}

/**
 * 无缓冲, 发送,接收
 */
func example2() {
	var ch1 chan int
	ch1 = make(chan int)
	go func() {
		ch1 <- 1
		fmt.Println("发送完毕")
	}()
	go func() {
		// 这里代码也能写在主线程上(主线程实际上也是一个goroutine在执行)
		/*a :=<-ch1
		fmt.Println("接收完毕",a)*/
		fmt.Println("接收完毕", <-ch1)
	}()
	time.Sleep(time.Second)
	fmt.Println("main end")
}

/**
 * 2个goroutine, 进行发送, 1个接收goroutine
 */
func example3() {
	var ch1 chan int
	ch1 = make(chan int)
	go func() {
		ch1 <- 1
		fmt.Println("发送完毕1")
	}()
	go func() {
		ch1 <- 2
		fmt.Println("发送完毕2")
	}()
	go func() {
		// 这里代码也能写在主线程上(主线程实际上也是一个goroutine在执行)
		/*a :=<-ch1
		fmt.Println("接收完毕",a)*/
		fmt.Println("接收完毕", <-ch1, len(ch1))
	}()
	time.Sleep(time.Second)
	fmt.Println("main end")
}

/**
 * 2个goroutine, 进行发送, 2个接收goroutine
 */
func example4() {
	var ch1 chan int
	ch1 = make(chan int)
	go func() {
		ch1 <- 1
		fmt.Println("发送完毕1")
	}()
	go func() {
		ch1 <- 2
		fmt.Println("发送完毕2")
	}()
	go func() {
		// 这里代码也能写在主线程上(主线程实际上也是一个goroutine在执行)
		/*a :=<-ch1
		fmt.Println("接收完毕",a)*/
		fmt.Println("接收完毕", <-ch1, len(ch1))
	}()
	go func() {
		// 这里代码也能写在主线程上(主线程实际上也是一个goroutine在执行)
		/*a :=<-ch1
		fmt.Println("接收完毕",a)*/
		fmt.Println("接收完毕", <-ch1, len(ch1))
	}()
	time.Sleep(time.Second)
	fmt.Println("main end")
}
