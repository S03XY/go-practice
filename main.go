package main

import "fmt"

func main() {
	var take int = 2
	var arg2 string = "shashank"
	var arg3 bool = true

	fmt.Printf("value %d %s %t \n", take, arg2, arg3)
}

/*
	go mod init
	touch main.go

	every file starts with a package name
	you can create multiple file from same package
	go uses first letter as capital letter to export something from a package

	To import other packages within your folder all you need to do it to import with path of current <name_of_current_package>/<path_to_other_package>
*/

/*
we can declare a variable in 3 ways either using var <varibale_name> type  = value or just var <variable_name> = value or using walrus operator
we use tye when we are specific about the spaces in the memory
<varibale_name> := value

and you dont need to put ; lexer does that automatically

*/
