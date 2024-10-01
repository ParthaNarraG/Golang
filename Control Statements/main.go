package main

import (
	"Math/rand"
	"fmt"
)

// Control statements with if, for and switch statements
func main() {
	//This will generate a random interger from o,n
	fmt.Println("random", rand.Intn(2))
	var userName string = "Partha"
	if userName == "Partha" {
		fmt.Println("This user name was Partha")
	} else {
		fmt.Println("This user was a different user")
	}

	//This variable works only on the if scope
	if x := rand.Intn(10); x == 2 {
		fmt.Println("This is a random integer", x)
	} else {
		fmt.Println("Not the required output", x)
	}

	//for can be declared in many ways
	for i := 0; i < 10; i++ {
		fmt.Println("i", i)
	}

	//we can ignore the initalization
	i := 0
	for ; i < 10; i++ {
		fmt.Println("i", i)
	}

	//We can also put increment in the code
	for i < 10 {
		fmt.Println("i-0", i)
		i++
	}

	//Forever loop until we break
	for {
		fmt.Println("i", i)
		i++
		if i > 10 {
			break
		}
	}

	//Using with continue
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println("oddIndx", j)
	}

	//Using for range method
	var namesList = []string{"Kohli", "Vijay", "Arjun", "Pawan"}
	for i, v := range namesList {
		fmt.Println(i, v)
	}

	//If we want we can ignore the index
	//The underscore tells the compiler to ignore the value
	for _, v := range namesList {
		fmt.Println("value", v)
	}

	//Similarly we can iterate through a map
	var data = map[string]bool{
		"User":   true,
		"Exists": false,
	}
	for k, v := range data {
		fmt.Println(k, v)
	}
	//If we don't give second variable by default it takes the first variable as keys/indexes
	for k := range data {
		fmt.Println(k)
	}
	data["isActive"] = true
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range data {
			fmt.Println(k, v)
		}
	}

	var simpleString string = "Hello World!"
	for i, k := range simpleString {
		fmt.Println(i, k, string(k))
	}

	// Labelling a for loop
outer:
	for i := 0; i < 5; i++ {
	inner:
		for _, v := range simpleString {
			if string(v) == "o" {
				//Here you are specifying i need to continue this statement
				continue inner
			}
			if string(v) == " " {
				//Here you are specifying i need to break outer for loop
				break outer
			}
			fmt.Println("value of string", string(v))
		}
	}

	//Even though we can specify break and default in go switch, most of the time we don't use those values
	switch i := 10; i {
	case 0, 5:
		fmt.Println("the number is", i)
	default:
		fmt.Println("the default number is,")
	}

loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0:
			fmt.Println(i)
		case 1:
			fmt.Println(i)
		case 2:
			fmt.Println(i)
		case 3:
			fmt.Println(i)
		case 4:
			fmt.Println(i)
		case 5:
			fmt.Println(i)
		case 6:
			fmt.Println(i)
		case 7:
			fmt.Println(i)
		case 8:
			fmt.Println("Breaking the loop")
			//If we don't specify the label for loop, then it will only break the switch statement
			break loop
		case 9:
			fmt.Println(i)
		}
	}

	//Blank switches instead of co;mparision in switch statement, you compare them in case statements
	//Once executed it will break from the switch(default break)
	switch wordLen := len("Hello"); {
	case true:
		fmt.Println("Hello", wordLen)
	case wordLen == 5:
		fmt.Println("World")
	}

}
