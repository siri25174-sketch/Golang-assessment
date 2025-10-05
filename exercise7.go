// Algorithm Implementation
// Write a linear search algorithm to search an element in a slice:
// Slice of integers [12, 34, 56, 78, 90, 11]
package main

import "fmt"

func main() {
	s := []int{12, 34, 56, 78, 90, 11}
	val := 78
	found := false
	for i, v := range s {
		if v == val { //Output: index if found, else print "Not Found"
			fmt.Printf("Value %d found at index %d\n", val, i)
			found = true
			break

		}
	}
	if !found {
		fmt.Printf("Value %d not found in the slice\n", val)
	}
	//d.Discuss time complexity in a comment (From class 19)
	//In the worst case, the program checks every element in the slice to find the value.

	//For example, if we search for a number that does not exist (like 4), it goes through the entire slice.

	//so,the time complexity is O(n) — the time increases as the number of elements increases to n th times.

	//The space complexity is O(1) — only a few extra variables are used, no matter how large the slice is.

}
