package main

// Single Line Comment
/*
	Block Comment
*/

import (
	"fmt"
	"os"
)

func main() {
	// All code will basically be a consequence of what happens inside of here!
	fmt.Println("Welcome to Golang Introduction! Woohoo!") 
	fmt.Println("---------------------------------------") 

	fmt.Println("\n1. Variables and Data Types")

	// In GoLang, you declare variables using the `var` keyword.
		// When you declare with `var`, the value can change!
	var name string // declaring a variable of type `string`, named `name`
	name = "John" // Assign the value to `name`
	fmt.Println("Name:", name)

	// You can declare and assign in one line.
	age := 30
	fmt.Println("Age:", age)

	var isStudent bool = false
	fmt.Println("Is Student:", isStudent)

	// Constants
		// Constant variables cannot change their value.
	fmt.Println("\n2. Constants")
	const pi float64 = 3.14159265359
	fmt.Println("Value of Pi:", pi)

	// Control Structures
	fmt.Println("\n3. Control Structures")
	// Conditionals (if-else)
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	// Loops (for)
	fmt.Println("\n4. Loops")
	// Loop has 3 parts:
		// Intialization of the index (i)
		// Condition to control how long the loop goes
		// Increment the loop so that it continues forward according to some formula
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Functions
	fmt.Println("\n5. Functions")
	// Function that will take in a name and return a sentence, which is assigned to a variable.
	greeting := greet("TacoCat")
	fmt.Println(greeting)

	// Arrays and Slices
	fmt.Println("\n6. Arrays and Slices")
	numbers := [5]int{1,2,3,4,5} // Declare an array with 5 integers
	fmt.Println("Array:", numbers)
	slice := numbers[0:5] // Slice will only include everything infront of the end index. (5, will get indexes including up to 4.)
	fmt.Println("Slice", slice)

	// Maps
	fmt.Println("\n7. Maps")
	// Maps store key-value pairs.

	// To declare an empty map, use make()
	// fruitMap := make(map[string]int)

	// This is how you create a literal map, with assigned KV pairs upon creation.
	fruitMap := map[string]int{
		"apple": 5,
		"orange": 10,
	}

	// Create a new element in the map
	fruitMap["banana"] = 7
	// Update an existing element in the map
	fruitMap["apple"] =  6

	// Loop over all key-value pairs in a map
	for fruit, quantity := range fruitMap {
		fmt.Printf("%s: %d\n",fruit, quantity)
	} 

	// To check to see if something exists inside of a map, for condition. (Good for validation.)
	if qty, ok := fruitMap["apple"]; ok {
		fmt.Println("Apple Quantity:", qty, ok)
	}

	fmt.Println("About to delete 'orange'", fruitMap)
	// Delete an element from a map
	delete(fruitMap, "orange")
	fmt.Println(fruitMap)

	// Structs
	fmt.Println("\n8. Structs")
	// Structs allow you to create your own custom data types.
	type Person struct {
		Name string
		Age int
		IsHuman bool
	}

	person := Person{Name: "Charlie", Age: 25, IsHuman: false}
	fmt.Println("Person", person)

	// Error Handling
	fmt.Println("\n9. Error Handling")
	// Go uses errors for handling exceptional situations
	result, err := divide(10, 2)
	
	if err != nil {
		// If err is `nil` that means there is an error.
		fmt.Println("Error:",err)
	} else {
		// If err is `nil` then there i no error, which means the function ran without an issue.
		fmt.Println("Result:", result)
	}

	// Defer and Panic
	fmt.Println("\n10. Defer and Panic")

	// Panic will indicate a critical error
	// Uncomment the lines below:
		// value := -5
		// if value < 0 {
		// 	fmt.Println("Recieved a negative value. This is a critical error! ⚠️")
		// 	panic("Critical Error Occured!")
		// }


	// Defer is used to schedule a functuion call to run after the surrounding function returns.
	defer fmt.Println("Defered statement")

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Defer is used for clean-up orperations like closing a file.
	defer file.Close() // This will close the file when main() exists
	// This defer will happen regardless of whether there is an error or not.

	fmt.Println("Regular statement")


	fmt.Println("\n-----------------")
	fmt.Println("End of GoLang Intro")
	fmt.Println()	
}

/*
This function takes in a name and will return a sentence with the incoming name as a greeting.
args:
	name string
return: string
*/
func greet(name string) string {
	return "Hello, " + name + "!"
}

/*
This function will perform division with two numbers

args:
	a any
	b float64
returns: float64, error
*/
func divide(a, b float64) (float64, error) {
	// Check if the divisor (b) is zero.
	if b == 0 {
		// If it is, return an error with a descriptive message
		// The `fmt.Errorf` function will create an error with a specific given message
		return 0, fmt.Errorf("division by zero")
	}

	// If divisor is not zero, then perform the division and return the result. (Error will be nil)
	return a/b, nil
}
