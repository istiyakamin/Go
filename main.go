package main

import "fmt"

func main()  {

	// var name string
	// name = "Istiyak Amin"
	// fmt.Println("Hello", name);

	// var a int = 10;
	// var b int = 20;

	var c, d int
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)

	sum_result := sum(c, d)
	fmt.Println(sum_result);

	serial_print(sum_result);

}

func sum(a int, b int) int {
	return a+b;
}

func serial_print(num int) int {
	for i := 1; i<=num; i++ {
		fmt.Println(i)
	}

	return 0;
}

// FIRST COMMIT 
// SECOND COMMIT 
// THIRD COMMIT
