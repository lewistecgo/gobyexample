package main

import "fmt"

func plus(a int, b int) (result int) {
	result = a + b
	return
}

func vals() (int, int) {
	return 10, 20
}

func main() {
	result := plus(3, 4)
	fmt.Println(result)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
}
