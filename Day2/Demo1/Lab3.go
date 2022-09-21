/*
Lab 3 - Create emp struct
              fields =   empno int, ename string, salary int
              func - incrSalary(percent), print -> fmt.println with all details
*/

package main

import (
	"fmt"
)

type Emp struct {
	empno  int
	ename  string
	salary int
}

func (empObj *Emp) incrSalary(percent int) {
	increasedSalary := (empObj.salary*percent)/100 + empObj.salary

	fmt.Println(empObj.ename, "employee id", empObj.empno, "having increased salary", increasedSalary)
}

func main() {
	empObj := Emp{empno: 22, ename: "Hardev", salary: 10000}

	empObj.incrSalary(10)
}
