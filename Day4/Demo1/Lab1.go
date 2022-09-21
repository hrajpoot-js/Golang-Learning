//go routine or thread explaination

package main

import (
	"fmt"
	"time"
)

func Print(str string) {
	for i := 0; i < 50; i++ {
		fmt.Print(str)
	}
}

func main() {
	go Print(".")
	go Print("x")

	fmt.Println("main function done")

	// sleep 2 sec
	//accept some details from user
	//code which takes more time

	time.Sleep(2 * time.Second)

}
