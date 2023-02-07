package main

import "fmt"

func main()  {

	// var name string
	// name = "Istiyak Amin"
	// fmt.Println("Hello", name);

	var a int = 10;
	var b int = 20;

	sum_result := sum(a, b)
	fmt.Println(sum_result);

}

func sum(a int, b int) int {
	return a+b;
}
