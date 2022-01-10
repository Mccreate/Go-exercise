package main

import (
	"fmt"
	"strings"
)

//Go needs Data Type in below.
/*
   func funcName(argv1 argv1's Datatype, argv2 argv2's Datatype, ...) return's DataType {
	   return blah
   }
*/
func multiply(a int, b int) int {
	return a * b
}

//Go can return many argvs.
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	fmt.Println(multiply(2, 2))
	fmt.Println(lenAndUpper("mccreate"))
}
