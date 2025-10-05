// Basics and For Loops
// Problem: Print all numbers from 1 to 20 using a for loop.
package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ { // condition for Forloop
		fmt.Printf("Result: %d \n", i)
	}
	//Also print whether each number is even or odd.
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, ":even")
		} else {
			fmt.Println(i, ":odd")
		}
	}
}
