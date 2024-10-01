// This package contains the code that starts the go program
package main

import (
	"fmt"
	"maps"
	"slices"
)

// If you create a variable with first letter capital the variable can accessed outside of the package
const Environment string = "Development" // public

// x: = 32; No var style will not work for global variables

func main() {
	variableCreation()
	defaultAliases()
	typeConversions()
	workingWithArrays()
	workingWithSlices()
	workingWithMaps()
	workingWithStructs()
	exercises3()
}

// This method consists of the ways we can create a variable
func variableCreation() {
	//Creating a typed variable with var keyword
	//Typed - We are enforcing the go compiler to use the data type int
	var number int = 10
	fmt.Println("number", number)

	//Creating an untyped variable with var keyword
	//If we do not enforce the type the compiler enforce the type at the compiling time
	var users = 0
	fmt.Println("users", users)

	// Tip: Once a type was assigned to a variable you cannot change the type again
	// Creating a variable with no var style
	licenses := 20 //Internally it will work as var licenses int
	fmt.Println("licenses", licenses)

	// If you declare a variable but you didn't assign a value the go compiler will assign default value based on the type you have specified
	var accountsSize int
	fmt.Println("accountsSize", accountsSize)

	// Assigning multiple values in a single line
	x, y, z := 1, "Sam", 17.40
	fmt.Println("x,y,z", x, y, z)

	var p1, p2, p3 = 1, false, ""
	fmt.Println("p1, p2, p3", p1, p2, p3)

	var (
		id     = 2
		name   = "Partha"
		height = 179.64
	)
	fmt.Println("id, name, height", id, name, height)

	// Creating a variable with const keyword, once assigned you cannot reassign the value
	const isLogged bool = false
	fmt.Println("isLogged", isLogged)

	// Tip: You cannot assign a var value to a const variable, if we want to assign they should be same declared type
	// numbers := 43
	// const dataSet int = numbers
	// fmt.Println("dataSet", dataSet)

	const numbers = 43
	const dataSet = numbers
	fmt.Println("dataSet", dataSet)
}

// This method shows how go compiler assigns default data types based on the values assigned to the variables
func defaultAliases() {
	age := 30
	fmt.Printf("%T\n", age)
	name := "Partha"
	fmt.Printf("%T\n", name)
	isLogged := true
	fmt.Printf("%T\n", isLogged)
	height := 179.89
	fmt.Printf("%T\n", height)
}

// This method shows you how to convert the typed values
func typeConversions() {
	// Ny default the type is int
	kilometers := 1000
	//If we assign the value to a float variable the compiler throws an error
	//---- var kms float32 = kilometers

	// The Correct way
	var kms float32 = float32(kilometers)
	fmt.Println("kms", kms)
}

func workingWithArrays() {
	//Creating an array variable
	//All the elements in an array has the same type, if we don't assign, the compiler assigns default values
	var usersList [3]int
	fmt.Println("usersList", usersList)

	//Assigning values to an array
	var dauList = [5]int{1, 3, 4, 5, 6}
	dauList[0] = 20
	fmt.Println("dauList", dauList)

	//Sparse array where most of the elements were zero and you define certain indexes not to be zero
	//in this case you are enforcing index 7 and 9 not to be zero
	var inactiveList = [10]int{1, 2, 7: 100, 9: 1000}
	fmt.Println("inactiveList", inactiveList)

	//We can also assign values without specifying the index
	var guestUsers = [...]int{1, 2, 3, 4, 5}
	fmt.Println("guestUsers", guestUsers)

	//You can compare two arrays if their array size and values were equal
	newUsers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("newUsers", newUsers)
	fmt.Println("Give truthy value", guestUsers == newUsers)

	// Finding the length of an array
	fmt.Println("Length of an array", len(newUsers))

	// Tip:Go considers the size of the array to be part of the type of the array.
	// This also means that you cannot use a variable to specify the size of an array, because types must be resolved at compile time, not at runtime.
}

func workingWithSlices() {
	//Creating a slice similar to an creating an array except you don't specify the size
	var usersList = []int{1, 2, 3}
	fmt.Println("usersList", usersList)

	var noElements []int
	fmt.Println("noElements", noElements)
	fmt.Println("Comparing with Nil", noElements == nil)

	//Sparse slice where most of the elements were zero and you define certain indexes not to be zero
	//in this case you are enforcing index 7 and 9 not to be zero
	var inactiveList = []int{1, 2, 7: 100, 9: 1000}
	fmt.Println("inactiveList", inactiveList)

	// You can compare two slices with the following function
	fmt.Println("Slices Equality", slices.Equal(noElements, inactiveList))
	fmt.Println("Length of slice", len(inactiveList))

	// You can append two slices and assign the result slice to a variable
	noElements = append(noElements, inactiveList...)
	noElements = append(noElements, 22, 78, 330, 454, 2023)
	fmt.Println("noElements after append", noElements)

	// Tip:There is a fixed capacity(Memory allocation) for every slice
	// When the slice size is getting increased by append function, the go runtime increases the capacity of the slice by
	// doubling it whenever it reaches it's peak(Max capacity)

	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 60)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 60)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 60)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 60)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 60)
	fmt.Println(x, len(x), cap(x))

	//If we want to decrease the burden on go run time, we can use make functions to enforce the size and capacity of a slice
	var makeSlice = make([]int, 5, 10)
	makeSlice = append(makeSlice, 10)
	makeSlice = append(makeSlice, 10)
	makeSlice = append(makeSlice, 10)
	makeSlice = append(makeSlice, 10)
	makeSlice = append(makeSlice, 10)
	makeSlice = append(makeSlice, 10)

	fmt.Println("makeSlice", makeSlice == nil, cap(makeSlice))

	//One of the best ways to use slice declare length as zero and capacity with some value
	var newMakeSlice = make([]int, 0, 10)
	fmt.Println("newMakeSlice", len(newMakeSlice), newMakeSlice == nil)

	//Tip:Never use capacity size smaller than array size the compiler throws an error

	//Emptying a slice, this will only clear the values to default, but the length still remains same
	clear(makeSlice)
	fmt.Println("makeSlice after clearing", makeSlice, len(makeSlice), cap(makeSlice))

	//Slicing of slices (Note:It will create a shallow copy)
	slicingArray := []string{"a", "b", "c", "d", "e", "f"}
	//Consider indexes less than given number
	slice1 := slicingArray[:2]
	fmt.Println("slice1", slice1)
	//Consider indexes from given number
	slice2 := slicingArray[2:]
	fmt.Println("slice1", slice2)
	//Consider indexes from given number 2 and 5
	slice3 := slicingArray[2:5]
	fmt.Println("slice1", slice3)
	//Consider all elements
	slice4 := slicingArray[:]
	fmt.Println("slice1", slice4)

	//If we try to change an element, it will reflect in every slice because of shared memory location
	slice4[0] = "red"
	slice4[3] = "yellow"
	fmt.Println("slicingArray", slicingArray)
	fmt.Println("slice1", slice1)
	fmt.Println("slice4", slice4)

	//If we want to copy a slice without affecting other slice
	var source = []int{1, 2, 3, 4, 5}
	var destination = make([]int, 2)
	//Copy function will return the number of elements that were copied from source to destination
	var copiedSlice = copy(destination, source)
	fmt.Println("destination", destination)
	fmt.Println("copiedSlice", copiedSlice)

	copy(destination, source[:2:3])
	fmt.Println("destination", destination)

	destination[0] = 44
	destination = append(destination, 45)
	fmt.Println("destination", destination, cap(destination))
	fmt.Println("source", source, cap(source))

	//Converting arrays to slices
	var xArray = [4]int{1, 2, 3, 4}
	var xSlice = xArray[:]
	xArray[0] = 100
	//Since slicing makes the elements the share location, whatever changes you make to an array reflect on slice
	fmt.Println("xArray", xArray)
	fmt.Println("xSlice", xSlice)

	//If you append before reassigning the value, the result might be different
	xSlice = append(xSlice, 23)
	fmt.Println("xSlice", xSlice)

	//Converting slices to arrays
	var ySlice = []int{1, 2, 3, 4}
	var yArray = [4]int(ySlice)
	fmt.Println("ySlice", ySlice)
	fmt.Println("yArray", yArray)

	//Since we are using type conversion instead of slicing the elements will not share memory allocation
	ySlice[0] = 100
	fmt.Println("ySlice", ySlice)
	fmt.Println("yArray", yArray)

	//Tip:The size of the array can be smaller, but it cannot be bigger than slice
}

func workingWithMaps() {
	//Map is a key value pair
	//We have to specify key is what type and value is of what type
	//Declaring a variable

	// In the below example we have specified key type is string and value type is int
	// Key type won't accept slices/maps
	var nilMap = map[string]int{}
	fmt.Println("nilmap", nilMap)
	fmt.Println("length of nilmap", len(nilMap))

	var intMap = map[string]int{
		"one": 1,
	}

	// In the below example we have specified key type is string and value type is slice of strings
	// Key type won't accept slices/maps
	var teams = map[string][]string{
		"CSK": []string{"Dhoni", "Raina", "Bravo"},
		"RCB": make([]string, 0, 10),
	}
	fmt.Println("teams", teams)
	fmt.Println("length of teams", len(teams))

	//If we have to specify the exact length an type of map, we can use make functions
	var newTeams = make(map[string][]string, 10)
	fmt.Println("Length of map", len(newTeams))

	//In many ways slices and maps are alike
	//The answer would be zero, cause even though you have specified the length, the go compiler will not create default key value pairs.
	//This is one of the differences between slice and map
	//For maps, the size parameter is just a hint for initial memory allocation. It doesn't create any key-value pairs
	//The size hint is used to potentially improve performance by reducing the number of internal resizing operations, but it doesn't affect the logical contents of the map.

	teams["RCB"] = append(teams["RCB"], "24")
	fmt.Println("After appending", teams["RCB"])
	//If it didn't find the specified key it will return the default value of value type, in this case empty slice
	fmt.Println("After appending", teams["KKR"])

	//Increments the value
	intMap["one"]++
	fmt.Println("intMap", intMap)

	//The comma ok idiom
	//If you want to know whether the key actually present in the map, we can use comma ok idiom
	//It has two variables v->value and ok, value is the value of the key and ok represents whether key actually exists in map
	//Tip:You can use any variable to declare, most common is ok, so we are going with ok
	s, ok := intMap["two"]
	fmt.Println(s, ok)

	//Deleting a key from a map
	delete(intMap, "one")
	fmt.Println("intMap after delete", intMap)

	//Clearing an entire map
	//Unlike slices, here the clear function sets the length of the map to zero
	clear(teams)
	fmt.Println("newTeams", newTeams)

	//If we want compare two maps, we can use same method of slice
	fmt.Println("Equal", maps.Equal(nilMap, intMap))

	//Using maps as sets
	var intset = map[int]bool{}
	var vals = []int{1, 2, 3, 4, 5, 6, 6, 6, 7, 8, 8, 9}
	for _, v := range vals {
		intset[v] = true
	}
	//Maps do not accept duplicate keys
	//Finding an element is easier in sets compared to slices
	fmt.Println("intset", intset)
}

func workingWithStructs() {
	//Structs are similar to maps, but they accept heterogenous elements
	//Declaring a struct
	type person struct {
		name string
		age  int
	}
	var user person
	fmt.Println("User", user)
	//Populate values to this struct
	user.name = "Partha"
	user.age = 20
	fmt.Println("User", user)

	//Another way of assigning the values
	var info = person{
		name: "Anil",
		age:  34,
	}
	fmt.Println("info", info)

	//Anonymous struct
	//It would be useful when the struct having only one or two properties
	var data = struct {
		name string
		age  int
	}{
		"Anonymous",
		45,
	}
	fmt.Println("data", data)
}

func exercises3() {

	// Write a program that defines a variable named greetings of type slice of strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
	// Create a subslice containing the first two values; a second subslice with the second, third, and fourth values;
	// and a third subslice with the fourth and fifth values. Print out all four slices.

	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	var firstSubSlice = greetings[:2]
	var secondSubSlice = greetings[1:4]
	var thirdSubSlice = greetings[3:5]

	// Write a program that defines a string variable called message with the value "Hi and " and prints the
	// fourth rune in it as a character, not a number.

	fmt.Println("firstSubSlice", firstSubSlice)
	fmt.Println("secondSubSlice", secondSubSlice)
	fmt.Println("thirdSubSlice", thirdSubSlice)

	var message string = "Hi 😀 and 😀"
	fourthRune := []rune(message)[3]
	fmt.Printf("fourthRune%c", fourthRune)

	// Write a program that defines a struct called Employee with three fields: firstName, lastName, and id.
	// The first two fields are of type string, and the last field (id) is of type int. Create three instances of this struct using whatever values you’d like.
	// Initialize the first one using the struct literal style without names, the second using the struct literal style with names, and the third with a var declaration.
	//  Use dot notation to populate the fields in the third struct. Print out all three structs.

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	// var firstInstance Employee(
	// 	"Partha",
	// 	"Narra",
	// 	"508"
	// )

	var secondInstance = Employee{
		firstName: "Partha",
		lastName:  "Narra",
		id:        508,
	}

	fmt.Println("secondInstance", secondInstance)

	var thirdInstance Employee

	thirdInstance.firstName = "Chinnu"
	thirdInstance.lastName = "Narra"
	thirdInstance.id = 740

	fmt.Println("thirdInstance", thirdInstance)

}
