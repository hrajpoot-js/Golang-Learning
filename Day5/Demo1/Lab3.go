//REST API (get method) examples

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Emp struct {
	Empno  int     `json:"empno"`
	Ename  string  `json:"ename"`
	Salary float64 `json:"salary"`
}

type EmpManager struct {
	emparr []Emp
}

func (empmgr *EmpManager) add(e1 Emp) {
	empmgr.emparr = append(empmgr.emparr, e1)
}

func main() {

	empmgr := new(EmpManager)
	emp := Emp{1, "One", 111}

	//data1
	empmgr.add(emp)

	//data2
	empmgr.add(Emp{2, "Two", 222})

	//data3
	empmgr.add(Emp{3, "THREE", 333})

	fmt.Println("EmpMgr ", empmgr)

	gethandler := func(w http.ResponseWriter, req *http.Request) {
		//WriteString always takes string as argument. So first Go struct is converted into byte array
		//and then byte array converted to string.
		bytearr, _ := json.Marshal(empmgr.emparr)
		io.WriteString(w, string(bytearr))
	}

	http.HandleFunc("/emp", gethandler)

	fmt.Printf("sever starting on 8080")
	http.ListenAndServe(":8080", nil)
}
