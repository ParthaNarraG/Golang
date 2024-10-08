package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Declaring a struct
type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var additionSum int = add(5, 10)
	fmt.Println("additionSum", additionSum)
	additionSum = add2(10, 20)
	fmt.Println("additionSum", additionSum)
	printMessage()
	optionalParameters(person{firstName: "Partha", lastName: "Narra"})
	optionalParameters(person{age: 28, lastName: "Narra"})
	varadicParameters(2, 3, 4, 5, 6, 89, 10)
	varadicParameters([]int{10, 20}...)
	//If you have multiple return values you have to assign each value to a variable
	division, modulodivision, err := returnMultipleValues(2, 2)
	if err != nil {
		fmt.Println(err)
		// The program will exit from here and below statements won't execute
		os.Exit(1)
	}
	fmt.Println("Multiple Return Type", division, modulodivision)
	//If you don't want to use a returned value use _ (underscore)
	division1, modulodivision1, _ := returnMultipleValues(2, 1)
	fmt.Println("Ignore variable", division1, modulodivision1)

	division2, modmodulodivision2, _ := returnNamedValues(2, 1)
	fmt.Println("Multiple Named Return Values", division2, modmodulodivision2)

	division3, modmodulodivision3, _ := blankReturns(2, 1)
	fmt.Println("Blank Returns", division3, modmodulodivision3)

	// Declaring the function as a variable
	// var functionVariable func(string) int
	// Binding the function to this varible
	var functionVariable func(string) int = printRandomMessage
	var result int = functionVariable("Hello")
	fmt.Println("Function Variable", result)

	// By default the value of this function is "nil"
	var funcVar func(string)
	fmt.Println("When no function were assigned to the variable", funcVar)

	//We are binding the functions to the map
	var vals = map[string]func(int, int) int{
		"+": addvalue,
		"-": subvalue,
		"*": multiplyvalue,
		"/": dividevalue,
	}

	//This expressions consist of map of map of strings
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression Syntax", expression)
			continue
		}
		// Convert the first character
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Checking if the value exists in the map
		_, ok := vals[expression[1]]
		if !ok {
			fmt.Println("This is an invalid operator", expression)
			continue
		}
		// Convert the second character
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Calling the function by passing the parameters
		fmt.Println(expression[1], vals[expression[1]](p1, p2))
	}

	// We can also declare the function using type
	// This practice is good for readability and documentation
	type opFunc func(int, int) int
	var newMaps = map[string]opFunc{}
	fmt.Println("newMaps", newMaps)

	anonymousFunctions()
	printDummyMsgForClosure()
	// Pass function as parameter
	sorting()
	// Return function as output
	retFunc := returnFunctions
	fmt.Println(retFunc()())

	deferFunc()
}

// Function with return type
func add(a int, b int) int {
	return int(a + b)
}

// If two or more parameter have same data type
func add2(a, b int) int {
	return int(a + b)
}

// Function with no parameters and no return type
func printMessage() {
	fmt.Println("You are just a dummy function")
}

// Named and optional parameters were not supported in go, however we can use structs to define which value we need to pass, and the
// remaining parameter will be assigned a default value
func optionalParameters(user person) {
	fmt.Println("OptionalParameters", user)
}

// Just like arguments in Javascript function
// Varadic Input Parameters
// Varadic parameters will take all the input parameters into a single slice
func varadicParameters(vals ...int) []int {
	var newSlice = make([]int, len(vals))
	for i, v := range newSlice {
		newSlice[i] = v + vals[i]
	}
	fmt.Println("varadicParameters", newSlice)
	return newSlice
}

// This function describes which return multiple values
func returnMultipleValues(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return numerator, denominator, errors.New("cannot divide by zero")
	}
	//If there are no errors, you generally return nil
	return numerator / denominator, numerator % denominator, nil
}

// Named return values
// By default these named return values were assigned values at compile time
// You can name the return values for better readabiity, but it's not good practice in all the cases
// Reasons
// a)These variables take up some space, and also not useful outside the function, cause there scope is limited
// b)You can shadow them meaning reassign
func returnNamedValues(numerator int, denominator int) (result int, remainder int, err error) {
	result = numerator / denominator
	remainder = numerator % denominator
	return result, remainder, err
}

// In this example even though we are not returning anything, the compiler returns result, remainder, and error cause they are assigned default values
// This is not a good practice, and most of the time we need to avoid blank returns and named return values
func blankReturns(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 || numerator < denominator {
		return
	}
	return
}

func printRandomMessage(msg string) int {
	return len(msg)
}

// Calculator functions
func addvalue(a, b int) int {
	return a + b
}

func subvalue(a, b int) int {
	return a - b
}

func multiplyvalue(a, b int) int {
	return a * b
}

func dividevalue(a, b int) int {
	return a / b
}

func anonymousFunctions() {
	// Inner functions are called anonymous functions
	newFunc := func(i int, j int) int { return i + j }
	for i := 0; i < 5; i++ {
		fmt.Println("Anonymous", newFunc(i, i+1))
	}
	anonymousFunctions1()
}

func anonymousFunctions1() {
	for i := 0; i < 5; i++ {
		// We haven't declared any name and this works just like IIFE in javascript
		// Not the ideal case, but just showing how it works
		func(j int) {
			fmt.Println("AnonymousFunctions1", i)
		}(i)
	}
}

// Closures the inner function has the ability to modify the resources of the outer function
func printDummyMsgForClosure() {
	var msg string = "Hello World!"
	func() {
		msg = "No, this is an old message!"
	}()
	fmt.Println("Message", msg)
}

// We can pass functions as a parameter to another functions
func sorting() {

	type person struct {
		name     string
		sur_name string
		age      int
	}

	var newSlice = make([]person, 0, 10)

	newSlice = append(newSlice, person{
		name:     "Partha",
		sur_name: "narra",
		age:      28,
	})

	newSlice = append(newSlice, person{
		name:     "chinnu",
		sur_name: "narra",
	})

	// Passing a function as a parameter to another functions
	// This can also be called higher order functions
	sort.Slice(newSlice, func(i, j int) bool {
		return newSlice[i].name > newSlice[j].name
	})

	fmt.Println("Sorting", newSlice)

}

// Similarly we can return the functions from a function
func returnFunctions() func() bool {
	return func() bool {
		return true
	}
}

// Defer statements
// Ḍefer statements is mostly used for resource cleanup like closing the files,network connections
// Defer statements execute at the last stage of the program consider it as not putting await in javascript
// Defer will follow stack principle, the last inserted defer statement will execute first

func deferFunc() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	deferLoop()
}

func deferLoop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	var newSlice = []int{1, 2, 3, 4, 5}
	var slices = newSlice[:2]
	fmt.Println(slices)
	newLoop(slices)
	fmt.Println("newSlice", newSlice)
}

// Go is call by value
// When you pass a variable for parameter to a function, go will create a copy of that variable and whatever changes done to that variable will not
// impact the original variable
// But maps & slices were different, cause they deal in pointers
func newLoop(slices []int) []int {
	slices[0] = 99
	return slices
}
