//Slice Operations
//Given a slice of strings: ["apple", "banana", "orange", "grapes", "mango"]
//Perform the following:
//Append a new fruit "kiwi"

//From class 18 slices and functions

package main

import "fmt"

func main() {
	s := []string{"apple", "banana", "orange", "grapes", "mango"}
	sub := []string{"kiwi"}
	s = append(s, sub...) // With this we are adding the kiwi as a new wlwmet to the string and continuing the output
	fmt.Println("Appended a new fruit:", s)

	//b.Remove "orange"
	s1 := []string{"apple", "banana", "orange", "grapes", "mango"}
	i := 2
	s1 = append(s1[:i], s1[i+1:]...) //here we are removing the orange from the string
	//:i is bascially representing the values before the index value 2
	fmt.Println("Removed orange:", s1)
	fmt.Println(s1)

	//c.Print the final slice
	fmt.Println("The final slice is :", s1)
}
