package main

import(
	"fmt"
	"time"
)

func main() {
	i:= 2

	fmt.Println("Write ", i , " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2: 
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//Can use commas to separate multiple expressions
	// in the same case statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekday")
	}

	//Alternate way to express if/else logic
	// case expressions can be non-constants 
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default: 
		fmt.Println("It's after noon")
	}

	//Functions inside functions 

	whatAmI := func( i interface{}) {
		switch t := i.(type)  {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm and int")
		default: 
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	//Assign functions to "variables"
	jess := whatAmI

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	jess(21)

}