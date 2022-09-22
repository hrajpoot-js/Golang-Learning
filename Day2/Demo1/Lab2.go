/*
	struct in Go Lang
	struct passing as function argument

*/

package main 

import ("fmt")

type Point2D struct {
	x int //x is small mean local use if it is X it can be globally used. Same for y
	y int
}

func shift(pt Point2D) {
	fmt.Println("Shift Line 1 - ", pt)

	pt.x = pt.x + 10
	pt.y = pt.y + 20

	fmt.Println("Shift Line 2 - ", pt)
}

//receiver function. see the diff from above. now this function can not be called directly 
//It will always be called by Point2D object.
func (pt Point2D) shift() {
	fmt.Println("Shift Line 1 - ", pt)

	pt.x = pt.x + 10
	pt.y = pt.y + 20

	fmt.Println("Shift Line 2 - ", pt)
}


func (pt *Point2D) shift1() {
	fmt.Println("Shift Line 1 - ", pt)

	pt.x = pt.x + 10
	pt.y = pt.y + 20

	fmt.Println("Shift Line 2 - ", pt)
}


func main() {
	p1 := Point2D {}

	fmt.Println("x = ", p1.x, " and y = ", p1.y)

	p2 := Point2D {x : 10, y : 20}

	fmt.Println("x = ", p2.x, " and y = ", p2.y)

	//call of receiveable function. no param passed but called on p2 (Point2D) object
	p2.shift()

	fmt.Println(p2)

	p2.shift1()

	fmt.Println(p2)
}
