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
	max := s[0] //for max value
	min := s[0] //for min value
	for _, value := range s {
		if value > max {
			max = value //if the value is greater than max it will update the max value
		}
		if value < min {
			min = value //if the value is less than min it will update the min value
		}
	}

	fmt.Printf("Maximum value: %d\n", max)
	fmt.Printf("Minimum value: %d\n", min)
}
