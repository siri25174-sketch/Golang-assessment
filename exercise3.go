//Exercise 3: Array Manipulation

package main

import "fmt"

func main() {
	//Create an array of 10 integers.

	arr := [10]int{1, 2, 3, 5, 7, 8, 1, 2, 4, 19}
	fmt.Println(arr)
	for i = 0; i <= 5; i++ {
		fmt.Print("enter the number:")
		fmt.Scan(&arr[i])
		fmt.Println("array:", i)
	}
	//

}
