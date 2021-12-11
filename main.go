package main

import "fmt"

func main() {
	fmt.Println("Hello from a module")

	//3 Way to declare variable
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	//Pointer
	var firstNamePointer *string = new(string) // Intilize pointer by assign an memory address of string
	*firstNamePointer = "Arthur"               //Refer the pointer variable and assign text into that memory address
	fmt.Println(*firstNamePointer)             //Dereference also use * to get value from the memory address

	ptr := &firstName //Address operator
	fmt.Println(ptr, *ptr)
	firstName = "Tricia"
	fmt.Println(ptr, *ptr)

	//Constant
	const pi = 3.415 //can be implicit type
	const c int = 3  // or explicite type
	fmt.Println(pi, c)
}
