//REST API using single handler. Use Request struct
//example of POST
//marshling of data & parsing json object from cleint in POST

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Emp struct {
	Empno  int    `json:"empno"`
	Ename  string `json:"ename"`
	Salary int    `json:"salary"`
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

	handler := func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("Request came for method", req.Method)

		//Get Method
		if req.Method == "GET" {
			fmt.Println("GET Request: ", req)
			bytearr, _ := json.Marshal(empmgr.emparr)
			//WriteString always takes string as argument. So first Go struct is converted into byte array
			//and then byte array converted to string.
			io.WriteString(w, string(bytearr))
		}

		//POST Method
		if req.Method == "POST" {

			/*
				emp1 := Emp{}
				emp1.Empno, _ = strconv.Atoi(req.FormValue("empno"))
				emp1.Ename = req.FormValue("ename")
				emp1.Salary, _ = strconv.Atoi(req.FormValue("salary"))

				empmgr.add(emp1)
			*/

			//In case if you are sending data as json object from postman, unmarshling will be needed

			//check decoder in docs, here is an example how to parse json object coming from client
			emp2 := Emp{}
			decoder := json.NewDecoder(req.Body)
			decoder.Decode(&emp2)

			empmgr.add(emp2)

			io.WriteString(w, "Hello From POST Method")
		}
	}

	http.HandleFunc("/emp", handler)

	fmt.Printf("sever starting on 8080")
	http.ListenAndServe(":8080", nil)
}
