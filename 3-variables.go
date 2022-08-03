package main
import ("fmt")

func main() {
        //One way to declare a variable 
	var age int =  10;
	fmt.Println("My age is: ", age);

	//another way to declare variable 
	name := "hardev singh";
	fmt.Println("My name is ", name);

	//multiple variable declaration of same type.
	var a, b, c, d int = 10,20,30,40;
	fmt.Println("Multiple variables declaration and initialization"); 
	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);
	fmt.Println(d);

	//not initialized variables
	fmt.Println("Default values of variables"); 
	var myage int;
	var myname string;
	var check bool;
	fmt.Println(myage);
	fmt.Println(myname);
	fmt.Println(check);

	fmt.Println("Multiple variable of multiple types declare and initialized");
	var yourname, yourage, adult="Ram", 35, true;
	fmt.Println(yourname, " and age ", yourage, "is an adult: ", adult);

	//variable declaration in a block

}
