//thirdparty module import which are from non- standard library

package main

import (
	"fmt"

	"github.com/magiconair/properties"
)

func main() {
	p := properties.MustLoadFile("myprop.properties", properties.UTF8)

	fmt.Println(p)

	fmt.Println(p.MustGetString("hostname"))
}
