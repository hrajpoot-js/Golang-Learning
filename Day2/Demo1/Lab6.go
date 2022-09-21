package main
import (
   "fmt"

)
func main() {
  a := make([]int, 5)  
  fmt.Println("for a currently len is " , len(a) ," and capacity is " , cap(a), " and contents ", a )
  a[0]=11
  a[1]=22
  a[2]=33
  a[3]=44
  a[4]=55
  fmt.Println("for a currently len is " , len(a) ," and capacity is " , cap(a), " and contents ", a)
  //a[5]=66
  a =	append(a,66)
  fmt.Println("for a currently len is " , len(a) ," and capacity is " , cap(a), " and contents ", a)
  a =	append(a,77)
  fmt.Println("for a currently len is " , len(a) ," and capacity is " , cap(a), " and contents ", a)
  a =	append(a,88)
  a =	append(a,99)
  a =	append(a,1)
  fmt.Println("for a currently len is " , len(a) ," and capacity is " , cap(a), " and contents ", a)
  a =	append(a,1)
  fmt.Println("for a currently len is " , len(a) ," and capacity is " , cap(a), " and contents ", a)
  for i, v := range a {
	fmt.Println("i = ",  i, "Value = ", v)
  }
  for i:=0; i< len(a); i++{
	fmt.Println("i = ",  i, "Value = ", a[i])
  }
}
