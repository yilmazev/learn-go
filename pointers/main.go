package main

import "fmt"

func main() {
	/*a := 10
	pointer := &a

	fmt.Println(a) // saf veri

	*pointer = 20 // veriyi güncelle

	fmt.Println(pointer) // pointer
	fmt.Println(a)       // yenilenmiş veri
	*/

	a := 10

	add12(&a)
	fmt.Println(a)

	// Notion not:
	// a := 10 // Değişken

	// pointer := &a // Veriyi Ram'den al

	// *pointer = 20 // Veriyi güncelle

	// fmt.Println(a) // Veriyi yaz
}

func add12(x *int) {
	*x = *x + 12
}
