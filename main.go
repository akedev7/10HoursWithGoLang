package main

import (
	"akedev7/10hourswithGo/controllers"
	"akedev7/10hourswithGo/models"
	"errors"
	"fmt"
	"net/http"
)

const (
	Red = iota
	Blue
	Green
)

func main() {
	basicGo()

	//Use struct from another package
	u := models.User{1, "Hello", "World"}
	fmt.Println(u)

	//Function
	port := 3000
	port, err := startWebServer(port)
	fmt.Println(port, err)

	_, err2 := startWebServer(port) //If the first param will not be used then use _ instead.
	fmt.Println(err2)

	//Build HTTP controller
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil) //nill = DefaulatServeMux

}

func startWebServer(port int) (int, error) { //Function can return multiple returned value
	fmt.Println("Starting server....")
	fmt.Println("Server started on port", port)
	return port, errors.New("Something went wrong")
}

func basicGo() {
	fmt.Println("Hello from a basic func")

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

	fmt.Println(Red, Blue, Green)

	//Array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//Slice
	slice := arr[:]
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice) //Same for arr and slice because, when use slice with arrays, it's refering to the array memory.

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)
	slice2 = append(slice2, 4, 5, 6)
	fmt.Println(slice2)

	s1 := slice2[1:]
	s2 := slice2[:2]
	s3 := slice2[1:2]
	fmt.Println(s1, s2, s3)

	//Map
	m := map[string]int{"foo": 2}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println(m)

	//Struct   // Struct is not class, it only contains properties
	type user struct { // can be moved out of main to package level
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Ake"
	u.LastName = "NaKrub"
	fmt.Println(u.FirstName)

	u2 := user{2,
		"You",
		"Too", // can leave , here
	}
	fmt.Println(u2)

}
