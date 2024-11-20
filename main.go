package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Variable
	greeting := "Hello world"

	fmt.Println(greeting, reflect.TypeOf(greeting))
}
