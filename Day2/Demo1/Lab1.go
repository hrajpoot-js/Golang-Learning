package main
import (
	"fmt"
	"common"
)
func main() {
   fmt.Println("In Lab1.main")
   sum := common.Add(100,4000);
   fmt.Println("Sum = ", sum)
   fmt.Println("Divide = ", common.Divide(1000,10))
}