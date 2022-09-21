package main
      import (
         "fmt"

      )
      func main() {
        var a [2]string
        a[0] = "aaaa"
        a[1]="bbbb"
        fmt.Println(a)
        fmt.Println(len(a))

        primes := [6]int{2, 3, 5, 7, 11, 13}
        fmt.Println(primes)
        fmt.Println(len(primes))
     }