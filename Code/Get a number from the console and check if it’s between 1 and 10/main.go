// Get a number from the console and check if it’s between 1 and 10.
package main

import "fmt"

func main(){
	var input int

	fmt.Println("Enter Your Number: ")
	fmt.Scan(&input)

	if input <= 10 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}