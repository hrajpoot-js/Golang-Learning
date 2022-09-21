//Interface explained

package main

import (
	"fmt"
	"strconv"
)

type display interface {
	Print()
}

type Emp struct {
	empId     int
	empName   string
	empSalary int
}

type Product struct {
	name  string
	price int
}

// interface function implementation
func (empObj Emp) Print() {
	str := "Name " + empObj.empName + " and salary " + strconv.Itoa(empObj.empSalary) + " his id " + strconv.Itoa(empObj.empId)

	fmt.Println(str)
}

func (prd Product) Print() {
	str := prd.name + "cost is " + strconv.Itoa(prd.price)
	fmt.Println(str)
}

func main() {
	ch := "e"
	fmt.Println(ch)

	fmt.Println("Enter your choice e for employee and p for product")
	fmt.Scanln(&ch)

	var obj display

	if ch == "e" {
		obj = Emp{empName: "Hardev", empId: 22, empSalary: 2000}
	}

	if ch == "p" {
		obj = Product{name: "shoe", price: 3000}
	}

	obj.Print()
}
