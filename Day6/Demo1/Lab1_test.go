package main

import (
	"fmt"
	"testing"
)

// func hello(str string) string {
func TestHello(t *testing.T) {
	got := Hello("fands")
	if got != "hello, fands" {
		t.Error("hello invocation got ", got, " want 'Hello, fands'")
	}
}

// func Add(no1 int, no2 int)int  {
func TestAdd(t *testing.T) {
	got := Add(10, 400)
	if got != 410 {
		t.Error("add invocation with 10, 400 got ", got, " but wanted 410")
	}
}

// func Divide(s1 string, s2 string) int {
func TestDivide1(t *testing.T) {
	got := Divide("100", "4")
	if got != 25 {
		t.Error("divide invocation with 100, 4 got ", got, " but wanted 25")
	}
}

func TestDivide2(t *testing.T) {
	defer func() {
		fmt.Println("before recover")
		r := recover()
		fmt.Println("after recover")
		if r != nil {
			//  fmt.Println("got ", got)
			fmt.Println("Recovered for Divide Invocation with 100 and a", r)
		} else {
			t.Error("Exception not raised ")
		}
	}()
	got := Divide("100", "a")
	fmt.Println("Not Problem, got", got)
}
