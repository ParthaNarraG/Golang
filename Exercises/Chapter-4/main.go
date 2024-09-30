package main

import (
	"fmt"
	"math/rand"
)

func main() {
	problem1()
	problem3()
}

// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
func problem1() {
	var newSlice = make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		newSlice = append(newSlice, rand.Intn(100))
	}
	fmt.Println("newSlice", newSlice, cap(newSlice))
	problem2(newSlice)
}

// Loop over the slice you created in exercise 1.
// For each value in the slice, apply the following rules:
// a)If the value is divisible by 2, print “Two!”
// b)If the value is divisible by 3, print “Three!”
// c)If the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
// d)Otherwise, print “Never mind”.

func problem2(newSlice []int) {
	for _, v := range newSlice {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six")
		case v%2 == 0:
			fmt.Println("Divisible by two")
		case v%3 == 0:
			fmt.Println("Divisible by three")
		default:
			fmt.Println("Never mind")
		}
	}
}

// Start a new program. In main, declare an int variable called total.
// Write a for loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive).
// The body of the for loop should be as follows: total := total + i fmt.Println( total)
// After the for loop, print out the value of total. What is printed out? What is the likely bug in this code?

func problem3() {
	var total int
	for i := 0; i < 10; i++ {
		total := +1
		fmt.Println("Total", total)
	}
	fmt.Println("Total", total)
}
