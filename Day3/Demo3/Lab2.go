//Interface example

package main

import (
	"fmt"
	"strings"
)

type Emp struct {
	empno  int
	ename  string
	salary int
}

type display interface {
	toUpper()
}

func (empObj Emp) toUpper() {
	fmt.Println("Employee Name in upper case is : ", strings.ToUpper(empObj.ename))
}

func main() {

	var obj display
	obj = Emp{ename: "hardev"}

	obj.toUpper()
}
