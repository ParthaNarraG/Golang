package main

import (
	"errors"
	"fmt"
	"os"
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
