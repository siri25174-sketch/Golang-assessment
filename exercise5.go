//Map Exercises
//Create a map to store student names as keys and scores as values:

package main

import "fmt"

func main() {
	//Print all students with their scores
	students := map[string]int{"A": 85, "C": 78, "B": 92, "D": 60}
	fmt.Println("Students and their scores:", students)

	//b.Print all students with their scores
	for name, score := range students {
		fmt.Printf("%s: %d\n", name, score)
	}
	//Find the student with the highest score

	//Here ot decide the highest score i am using for loop with range and two variables .
	highestScore := 0
	topStudent := "" //As we dont kkow the name of the top student we are initializing it with empty string
	for name, score := range students {
		if score > highestScore {
			highestScore = score //If yes it will update the highest score
			topStudent = name    //And it will update the name of the student
		}
	}
	fmt.Println("Top student is ", topStudent, "with score", highestScore)

	//c.Add a new student "E" with score 88
	var student string = "E"
	var score int = 88
	students[student] = score //Here students is the map which has the key as student and value as score
	fmt.Println("Added new student:", students)

	//d.Check if "C" exists using the ok idiom
	//from class 19 maps Retrieving Values from a Map

	makeMap := map[string]int{"A": 85, "C": 78, "B": 92, "D": 60, "E": 88}
	value, ok := makeMap["C"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("C exists in the map with score:", value)
	} else {
		fmt.Println("C does not exist in the map")
	}

}
