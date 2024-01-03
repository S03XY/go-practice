package main

import (
	"fmt"

	"github.com/S03XY/practice/package/home"
	"github.com/S03XY/practice/package/utils"
)

func main() {
	fmt.Println("hello world")
	home.HomeFunction()
	utils.UtilsFunction()
}

/*
	go mod init
	touch main.go

	every file starts with a package name
	you can create multiple file from same package
	go uses first letter as capital letter to export something from a package

	To import other packages within your folder all you need to do it to import with path of current <name_of_current_package>/<path_to_other_package>
*/
