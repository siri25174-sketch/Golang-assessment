//Exercise 2: Ranges and Slices
//Create a slice of integers: [10, 23, 45, 12, 34, 56, 78]

package main

import "fmt"

func main() {

	s := []int{10, 23, 45, 12, 34, 56, 78}
	fmt.Println("Slice of given integer:", s)

	//Use range to:
	//Print all elements
	//Calculate the sum of elements
	//Find the maximum and minimum values

	for _, value := range s {
		fmt.Println("value:", value)
	}
	sum := 0 //here im calculating the sum of all the elemnts given
	for _, value := range s {
		sum += value
	}
	fmt.Println("Sum of numbers:", sum)

	//max and min values
}
