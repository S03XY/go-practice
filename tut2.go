package main

import (
	"fmt"
)

type Employee struct {
	first_name string
	last_name  string
}

func (e Employee) get_details() string {
	return "first name: " + e.first_name + " last name: " + e.last_name
}

func (e *Employee) update_employee(new_first_name string, new_last_name string) {
	e.first_name = new_first_name
	e.last_name = new_last_name
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

// inherittance

type Animal struct {
	Breed string
}

func (a Animal) GetBreed() string {
	return a.Breed
}

// struct can store struct but then the access should be different

type Dog struct {
	Animal
	Legs int
}

func (d Dog) GetBreed() string {
	return d.Breed
}

func (d Dog) GetLegs() int {
	return d.Legs
}

func Tut2() {
	fmt.Println("from tut2")

	value1 := map[int]string{}
	value1[1] = ""

	slice1 := []string{}
	slice1 = append(slice1, "Shashank")
	fmt.Println(slice1)

	// make is used for creating a slice, map or a channel
	// value2 := make(map[int]string)

	// slice2 := make([]string, 0)
	// slice2 = append(slice2, "appended")
	// fmt.Println("slice 2", slice2)

	// employee := Employee{"", ""}
	// employee := Employee{first_name: "Shashank", last_name: "Pandey"}
	// employee.first_name = "rohan"
	// update_employee(&employee, "subject1")
	// fmt.Println("first_name", employee.first_name)

	// new value will allocate some memory in heap and return a reference to it
	employee := new(Employee)
	employee.first_name = "shashank"
	employee.last_name = "pandey"
	// fmt.Println((*employee).first_name)
	// fmt.Println(employee.first_name)
	// fmt.Println(employee.get_details())
	// employee.update_employee("rohan", "tiwari")
	// fmt.Println(employee.get_details())

	var reactangle Shape = Rectangle{height: 10, width: 10}

	fmt.Println("area", reactangle.Area())
	fmt.Println("perimeter", reactangle.Perimeter())

	a := Animal{Breed: "dog"}
	d := Dog{Legs: 4, Animal: a}
	fmt.Println(d.GetBreed())

}

// func update_employee(employee *Employee, new_name string) {
// 	employee.first_name = new_name
// }
