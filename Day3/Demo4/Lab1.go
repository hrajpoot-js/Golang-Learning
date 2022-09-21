package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type UserStruct struct {
	Data struct {
		Id        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

func main() {
	ch := "1"
	fmt.Println("Enter a number 1..12")
	fmt.Scanln(&ch)
	url := "https://reqres.in/api/users/" + ch
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
	//fmt.Println("Body content   = ", string(body))
	res := UserStruct{}
	json.Unmarshal(body, &res)
	fmt.Println("User Email Name " + res.Data.Email)
	fmt.Println("User First Name " + res.Data.FirstName)

	//f, err := os.OpenFile("testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	f, err := os.OpenFile("demo.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("This is a test log entry")
	log.Println("User Email Name " + res.Data.Email)
	log.Println("User First Name " + res.Data.FirstName)
}
