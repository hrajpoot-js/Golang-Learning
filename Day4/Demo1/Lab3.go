//channel concept explained

package main

import (
	"fmt"
	"strconv"
	"time"
)

func writer(str string, c chan string) {
	for i := 1; i <= 5; i++ {
		fmt.Println("#########in count ", i)
		c <- str + strconv.Itoa(i)
		time.Sleep(time.Millisecond * 100)
	}
}
func reader(c chan string) {
	for msg := range c {
		fmt.Println("in reader1 ", msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//here 1 is buffer size of channel
	ch := make(chan string, 1)
	go writer("S", ch)
	go reader(ch)
	time.Sleep(time.Millisecond * 5000)
}
