/*
	Fibonacci Series printing
*/

package main 

import (
	"fmt"
	"os"
	"strconv"
)

func printFibonacciSeries(num int) {
	x := 0
	y := 1
	z := 0
	fmt.Printf("Fibonacci Series is: %d %d", x, y)
	for y < num{
	   z = y
	   y = x + y
	   x = z
	   fmt.Printf(" %d", y)
	}

	fmt.Println()
}

func main() {

	cmdArgs := os.Args[1]

	userInput,_ := strconv.Atoi(cmdArgs)
	printFibonacciSeries(userInput)

}