package main

import (
	"fmt"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Current data in ", primes)
	fmt.Println("Length of Prime = ", len(primes))
	s1 := primes[1:3]
	fmt.Println("S1 ", s1, "len = ", len(s1))
	primes[2] = 333
	s1 = primes[2:4]
	fmt.Println("S1 ", s1, "len = ", len(s1))
	s1[4] = 9
	fmt.Println("Current data in ", primes)

}
