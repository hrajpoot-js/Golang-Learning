package main

import (
	"fmt"
	"time"
)

func add(i, j int, ch chan int) {
	fmt.Println("add invoke with ", i, " ", j)
	time.Sleep(2 * time.Second)
	fmt.Println("add 2 sec sleep over")

	//sending result via channel
	ch <- (i + j)
	fmt.Println("in add after ch send 1")
}

func main() {
	//go sum := add(5, 10) //wrong
	//sum := go add(5, 10) //worng

	ch := make(chan int)
	go add(5, 10, ch)
	time.Sleep(5 * time.Second)
	fmt.Println("main 5 sec sleep over")

	//receiving the channel result in sum variable
	sum := <-ch

	fmt.Println("sum = ", sum)

	fmt.Println("main 5 sec sleep started again")
	time.Sleep(5 * time.Second)

}
