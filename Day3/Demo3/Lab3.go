package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	ch := "1"
	fmt.Println("Enter a number 1..12")
	fmt.Scanln(&ch)
	//url := "https://reqres.in/api/users/" + ch
	url := "https://reqres.in/api/users?page=" + ch
	fmt.Println("Fetching " + url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error = ", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		fmt.Println("Wrong user id, ...1..12")
		return
	}
	//fmt.Println("Server Returned  = ", resp)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Body content   = ", string(body))

}
