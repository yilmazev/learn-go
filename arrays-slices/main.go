package main

import (
	"fmt"
)

func main() {
	accounts := [2]string{"array 1", "array 2"} // sabit array

	socialMedia := []string{"instagram", "x", "twitter"} // dinamik array
	socialMedia = append(socialMedia, "facebook")        // array'e yeni veri ekleme

	fmt.Println(accounts)
	fmt.Println(socialMedia)
	fmt.Println(socialMedia[0:2]) // slices
}
