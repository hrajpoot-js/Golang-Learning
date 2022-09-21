//read and write into the file

package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	inputFile := os.Args[1]
	message := os.Args[2]

	//write user message into the input file
	err := os.WriteFile(inputFile, []byte(message), 0666)
	if err != nil && os.IsExist(err) {
		log.Fatal(err)
	}

	//read file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
	
	fmt.Println()

}