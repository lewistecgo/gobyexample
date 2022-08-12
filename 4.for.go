package main

import "fmt"

func main() {

	i := 1
	for i <= 10 {
		fmt.Println(i)
		//i = i + 1
		i++
	}

	for j := 7; j > 0; j-- {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
