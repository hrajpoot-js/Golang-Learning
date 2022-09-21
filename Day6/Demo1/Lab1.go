//Unit testing in go lang explained in this demo

// run this to see the coverage - $go tool cover -html=tmp.txt -o tmp.html or $go test Lab1_test.go Lab1.go -v -coverprofile tmp.txt

package main

import (
	"fmt"
	"strconv"
)

func Hello(str string) string {
	fmt.Println("hello invoked with ", str)
	return "hello, " + str
}
func Add(no1 int, no2 int) int {
	fmt.Println("add invoked with ", no1, ",", no2)
	return no1 + no2
}

/*function orientation says if even error happnes tells that in fnction
prototye. so return -1 is not what is recommended. that is old style coding.

below error could be handled this way. see Lab2 done this way

func Divide(s1 string, s2 string) (int, error) {

	//function

	   num1, err := strconv.Atoi(s1)
	   num2, err := strconv.Atoi(s2)
	   if err != nil {
	       fmt.Println("invalid string")
	       return 0,err
	   }
}


	So many if this is actully bloating code of just one line num1/num2
	if else bloat your code 400 times.. try catch bloat your code by 100 times
	but it is better than if else . in golang we don not have try catch
	so use recover . see Lab2.go

*/

func Divide(s1 string, s2 string) int {
	// two types of errors
	//	1. invalid string where strconv gives errors
	//	2. second argument is zero
	fmt.Println("divide invoked with ", s1, ", ", s2)
	num1, _ := strconv.Atoi(s1)
	num2, _ := strconv.Atoi(s2)
	return num1 / num2
	/*
	   if err != nil {
	       fmt.Println("invalid string")
	       return -1
	   }

	   if num2 == 0 {
	       fmt.Println("second argument is zero")
	       return -1
	   }
	*/

}
func main() {
	fmt.Println("hello with fands returned ", Hello("fands"))
	fmt.Println("add with 10,40 returned ", Add(10, 40))
	fmt.Println("divide with 40,10 returned ", Divide("40", "10"))
	fmt.Println("divide with 40,a returned ", Divide("40", "a"))
	fmt.Println("divide with 40,0 returned ", Divide("40", "0"))
}
