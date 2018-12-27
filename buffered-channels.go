package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	// c <- 3  缓冲区写满会异常退出
	fmt.Println(<-c)
	fmt.Println(<-c)
}
