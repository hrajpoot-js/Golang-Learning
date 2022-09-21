package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	indexhandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>Index Page</h1><h2><a href='p1'>Page1</a></h2><h2><a href='p2'>Page2</a></h2>")
	}
	page1handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>Page1 Page</h1>")
	}
	page2handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>Page2 Page</h1>")
	}
	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/p1", page1handler)
	http.HandleFunc("/p2", page2handler)

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
