package main

import (
	"fmt"
)

type Emp struct {
	empno  int
	ename  string
	salary int
}

func main() {

	mymap := make(map[int]Emp)

	mymap[1] = Emp{1, "dev", 2000}
	mymap[2] = Emp{2, "raj", 1000}

	for v := range mymap {
		fmt.Println(mymap[v])
	}
}
