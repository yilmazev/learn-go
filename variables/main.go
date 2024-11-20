package main

import (
	"fmt"
	"reflect"
)

func main() {
	var age int = 19 // Büyük verileri tutmak için (int64)

	var isHasAccount bool = true

	var surname = "Yılmaz"

	name := "Yılmaz"

	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(surname)
	fmt.Println(age)
	fmt.Println(isHasAccount)
}
