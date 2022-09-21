package main 

import (
	"fmt"
)

type Emp struct {
	empno int
	ename string
	salary int 
} 

type EmpManger  struct {
	employeeArray [10]Emp
}

func (empMngr *EmpManger) add(position int, e Emp) {
	if position+1 < 10 {
		empMngr.employeeArray[position-1] = e
	} else {
		fmt.Println("No more position left to fill")
	}
}

func (empMngr *EmpManger) print() {

	for _,v := range empMngr.employeeArray {
		fmt.Println(v)
	}
}

func main() {
	
	employee1 := Emp{ empno : 1, ename : "dev1", salary : 2000 }
	employee2 := Emp{ empno : 2, ename : "dev2", salary : 3000 }
	employee3 := Emp{ empno : 3, ename : "dev3", salary : 4000 }
	employee4 := Emp{ empno : 4, ename : "dev4", salary : 5000 }
	
	empMngrObj := EmpManger{}

	empMngrObj.add(1, employee1)
	empMngrObj.add(2, employee2)
	empMngrObj.add(3, employee3)
	empMngrObj.add(4, employee4)

	empMngrObj.print()
}