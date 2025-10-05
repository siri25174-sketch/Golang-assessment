//Real-Time Scenario
//Did not push yet

package main

import "fmt"

func main() {
	//Create a small inventory system using slices and maps:
	//Slice of product names: ["Laptop", "Mouse", "Keyboard", "Monitor"]
	products := []string{"Laptop", "Mouse", "Keyboard", "Monitor"}
	fmt.Println("Products available:", products)

	//b.Map to store stock quantity: map[string]int
	storestock := map[string]int{"Laptop": 20, "Mouse": 30, "Keyboard": 16, "Monitor": 20}
	fmt.Println("Stock available:", storestock)

	//c.Allow the user to:
	//Add stock for a product
	var productName string
	var quantity int
	//we are doing this with the user input
	fmt.Print("Enter product name to add stock: ")
	fmt.Scanln(&productName)

	fmt.Print("Enter quantity to add: ")
	fmt.Scanln(&quantity)

	storestock[productName] += quantity
	fmt.Println("Updated stock after addition:")
	for _, product := range products { //I dont need the index so keeping it blank
		fmt.Printf("%s : %d\n", product, storestock[product]) //Here as i dont want to specify the index i wanted to print the product name and its quantity
	}

	//d.Remove stock for a product
	fmt.Print("Enter product name to remove stock: ")
	fmt.Scanln(&productName)

	fmt.Print("Enter quantity to remove: ")
	fmt.Scanln(&quantity)

	storestock[productName] -= quantity
	fmt.Println("Updated stock after removal:", storestock)

	for product, quantity := range storestock {
		fmt.Printf("%s : %d\n", product, quantity)

		//e.Print the final inventory
		fmt.Println("Updated stock after removal:")
		for _, product := range products {
			fmt.Printf("%s : %d\n", product, storestock[product])
		}

		//Display products with stock less than 5 for alert purposes
		fmt.Println("Products that are less than 5:")
		for _, product := range products {
			if storestock[product] < 5 {
				fmt.Printf("%s : %d\n", product, storestock[product])
			}
		}
	}
}
