package main

import "fmt"

func main() {

	var a = [5]int{1, 2, 3, 4, 5}
	var b [5]int
	fmt.Println(a)
	fmt.Println("emp: ", b)
	b[0] = 1
	fmt.Println("empn: ", b)

	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}
