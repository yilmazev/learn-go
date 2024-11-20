package main

import "fmt"

func main() {
	total := add(10, 20)
	fmt.Println(total)

	numbers := []int{30, 30}

	fmt.Println((sum(numbers)))

	fmt.Println(sum2(2, 5, 9, 12, 24, 1))
}

func add(x int, y int) int {
	return x + y
}

func sum(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func sum2(numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}
