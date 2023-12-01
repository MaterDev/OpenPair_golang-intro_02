package main

import "fmt"

/*
    Key Concepts Covered:

    1. Arrays: Fixed-size, ordered collections of elements of a single type. The script shows how to declare, initialize, and manipulate arrays, including multi-dimensional arrays.
	
    2. Slices: Dynamic and flexible segments of an array. The script explores slice creation, appending elements, and creating sub-slices, illustrating their dynamic nature compared to arrays.

    3. Maps: Key-value pairs that provide a versatile way to store and manage data. The demonstration includes map creation, adding and removing elements, and iterating over maps.

    4. Structs: Custom data types that allow the grouping of multiple fields under a single unit. The script illustrates defining structs, creating instances, and modifying their fields.
*/


// Arrays in Go
func demonstrateArrays() {
    // Declaring an array of 10 integers. In Go, arrays are of fixed size.
    // The type of elements (in this case 'int') and length (10) are part of the array's type.
    var numbers [10]int 
    for i := 0; i < len(numbers); i++ {
        // Initializing each element of the array.
        // Here, each element is assigned twice its index.
        numbers[i] = i * 2 
    }
    // Printing the entire array to show the initialized values.
    fmt.Println("Numbers array:", numbers)

    // Demonstrating a multi-dimensional array, specifically a 3x3 matrix.
    // Multi-dimensional arrays are essentially arrays of arrays.
    var matrix [3][3]int
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            // Assigning a value based on the sum of row and column indices.
            matrix[i][j] = i + j
        }
    }
    // Printing the multi-dimensional array to show its structure.
    fmt.Println("Matrix:", matrix)
}

// Slices in Go
func demonstrateSlices() {
    // Slices are similar to arrays but are resizable and more flexible.
    // Here, we're creating a slice of strings with initial elements.
    colors := []string{"Red", "Green", "Blue"}
    // Printing the initial slice.
    fmt.Println("Initial colors slice:", colors)

    // Demonstrating the dynamic nature of slices by appending more elements.
    moreColors := []string{"Yellow", "Purple", "Orange"}
    // The append function returns a new slice containing the existing and new elements.
    colors = append(colors, moreColors...)
    // Printing the extended slice to show the added elements.
    fmt.Println("Extended colors slice:", colors)

    // Creating a sub-slice from the existing slice. Slices support the slice[low:high] syntax.
    // This selects a range starting at 'low' index and ending just before the 'high' index.
    subsetColors := colors[1:4] // Selects elements at index 1, 2, and 3.
    // Printing the sub-slice to demonstrate slicing.
    fmt.Println("Subset of colors:", subsetColors)
}

// Maps in Go
func demonstrateMaps() {
    // Maps are key-value pairs. They are similar to dictionaries in other languages.
    // Here, we define a map with strings as keys and ints as values.
    userAges := map[string]int{
        "Alice": 28,
        "Bob":   34,
    }

    // Adding a new key-value pair to the map.
    // Maps are dynamic, and new pairs can be added at any time.
    userAges["Charlie"] = 40
    // Printing the map to show the newly added pair.
    fmt.Println("User ages with Charlie added:", userAges)

    // Deleting a key-value pair from the map using the 'delete' built-in function.
    delete(userAges, "Bob")
    // Printing the map to show the state after deletion.
    fmt.Println("User ages after deleting Bob:", userAges)
}

// Structs in Go
func demonstrateStructs() {
    // A struct is a composite type that groups together variables under a named unit.
    // Here, we define a 'Person' struct with 'Name' and 'Age' fields.
    type Person struct {
        Name string
        Age  int
    }

    // Creating a slice of 'Person' structs with initial values.
    people := []Person{
        {Name: "Alice", Age: 28},
        {Name: "Bob", Age: 34},
    }

    // Adding a new 'Person' to our slice of structs.
    // Structs can be added to slices like any other data type.
    people = append(people, Person{Name: "Charlie", Age: 40})
    // Printing the slice of structs to show all the persons.
    fmt.Println("People slice:", people)

    // Modifying a field in one of our structs.
    // Here, we're updating Alice's age. Struct fields are accessed using a dot (.)
    people[0].Age = 29
    // Printing the updated slice to show the changed age.
    fmt.Println("Updated people slice:", people)
}

func main() {
    // Calling our demonstration functions to show different concepts in action.
    // Each function will print its results to the console.
    demonstrateArrays()   // Shows fixed-size arrays and multi-dimensional arrays.
    demonstrateSlices()   // Demonstrates dynamic slices, append function, and sub-slicing.
    demonstrateMaps()     // Covers maps with key-value pairs, adding, and deleting elements.
    demonstrateStructs()  // Explains struct definition and manipulation in a slice of structs.
}
