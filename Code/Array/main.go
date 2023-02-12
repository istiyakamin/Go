package main

import "fmt"

func main() {

	var a [5]int
	b := [5]int{1, 2, 3, 4, 5}

	a[2] = 3
	a[0] = 1
	a[1] = a[2]-a[0]
	a[3] = 4
	a[4] = 5

	var two_dim [5][3]int

	fmt.Println(a)
	fmt.Println(b)

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			two_dim[i][j] = j
		}
	}

	fmt.Println(two_dim)

}