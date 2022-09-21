//Read file and send it to http server

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	//read file
	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	//write file data
	IndexFilehandler := func(w http.ResponseWriter, req *http.Request) {
		//file data is byte array so converting it into string
		io.WriteString(w, string(data))
	}

	//handle the function
	http.HandleFunc("/index", IndexFilehandler)

	fmt.Printf("sever starting on 8080")
	http.ListenAndServe(":8080", nil)
}
