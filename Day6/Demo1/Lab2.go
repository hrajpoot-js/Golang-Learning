package main

import (
	"fmt"
	"strconv"
)

func Divide(s1 string, s2 string) (int, error) {
	defer func() {
		r := recover()
		fmt.Println("---------------recover ", r)
		if r != nil {
			fmt.Println("Recovered in f", r)
			//return 0, err
		}
	}()
	fmt.Println("divide invoked with ", s1, ", ", s2)
	num1, err := strconv.Atoi(s1)

	if err != nil {
		panic("Conversion Error")
	}
	num2, err := strconv.Atoi(s2)
	if err != nil {
		panic("Conversion Error")
	}
	ans := num1 / num2
	return ans, nil
}

func main() {
	ans1, err := Divide("40", "10")
	ans2, err := Divide("40", "a")
	ans3, err := Divide("40", "0")
	fmt.Println("divide with 40,10 returned ", ans1, " Error: ", err)
	fmt.Println("divide with 40,a returned ", ans2, " Error: ", err)
	fmt.Println("divide with 40,0 returned ", ans3, " Error: ", err)
}
