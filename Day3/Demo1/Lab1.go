//map example

package main

import (
	"fmt"
)

func main() {
	var mymap map[string]int
	mymap = make(map[string]int)
	fmt.Println("mymap details ", len(mymap))
	mymap["a"] = 100
	fmt.Println("mymap details ", len(mymap))
	mymap["b"] = 200
	fmt.Println("mymap details ", len(mymap))
	mymap["c"] = 300
	fmt.Println("mymap details ", len(mymap))
	mymap["a"] = 10
	fmt.Println("mymap details ", len(mymap))
	mymap["b"] = 20
	fmt.Println("mymap details ", len(mymap))

	fmt.Println(mymap)
}
