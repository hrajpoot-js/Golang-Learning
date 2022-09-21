package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	hname, herr := os.Hostname()
	if herr == nil {
		fmt.Println("Hostname: ", hname)
	} else {
		fmt.Println("Hostname Error: ", herr)
	}

	//Show current time
	timenow := time.Now()
	fmt.Println("Current Time: ", timenow)
}
