/*
	Program to show multiple return values
*/

package main 

import (
	"fmt"
)

/*
//this is called un-nammed variable
func addSub(x int, y int) (int, int) {
	return x+y, x-y
}
*/

//this is called nammed return.
func addSub(x , y int) (a, b int) {
	a, b = x+y, x-y
	return
}

func main() {
	add, diff := addSub(50,25)
	fmt.Println("addSub returns: ", add, " & ", diff)
}