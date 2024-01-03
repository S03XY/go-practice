package main

import "fmt"

func collection() {
	fmt.Println("from collection practice")
	// type conversion can be done
	// a := 5
	// b := &a // pointer which stores reference to get the access you can change it using derefering methods *
	// *b = 6
	// fmt.Println(b)

	// array are fixed sized
	// array := [...]int64{1, 2, 3, 4}

	// slice are dynamic
	// array := []int64{1, 2, 3, 4}
	// array = append(array, 1)
	// fmt.Println("array", array[2:])

	// if else switch nare normal

	a := 1

	switch a {

	case 1:
		fmt.Println("case 1")
		// fallthrough this is used to skip run the next case as well
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("default")

	}

	// there are three type of for loops
	// classic for loop
	// for count := 0; count < 5; count++ {
	// 	fmt.Println(count)
	// }

	//  while loop
	// count := 0
	// for count < 5 {
	// 	fmt.Println(count)
	// 	count++
	// }

	// for range loops
	// collection := []int64{1, 2, 2, 3, 4}
	// for i, val := range collection {
	// 	fmt.Printf("%d and %d \n", i, val)
	// }

	// var lam = func() bool {
	// 	return false
	// }
	// lam()

	// lam := func() bool { return false }
	// lam()

	// var c, b = multiple_return()
	// var c, b string = multiple_return()

	// var (
	// 	some1 int
	// 	some2 int
	// 	some3 int
	// )

	// _, b := multiple_return()

	// _ is used to ignore the values
	// fmt.Printf("%s \n", b)

	// value := return_ref()
	// fmt.Printf("%s \n", *value)

	// deffered function execute the the end of the scope and order of execution is accorging to stack first in last out
	// defer func() {
	// 	fmt.Println("deferred function 1")
	// }()

	// defer func() {
	// 	fmt.Println("deferred function 2")
	// }()

	// fmt.Println("normal call")

}

// func multiple_return() (string, string) {
// 	return "shashank", "pandey"
// }

// in go we can return a reference in a function
// func return_ref() *string {
// 	a := "Shashank"
// 	return &a
// }
