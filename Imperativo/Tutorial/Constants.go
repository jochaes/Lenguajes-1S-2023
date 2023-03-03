package main

import(
	"fmt"
	"math"
)

//Const declare a constant value 
const s string  = "Constant"

func main(){
	fmt.Println(s)
	//const can appear anywhere a var can 
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	//Numeric cinstant has no type until it's given one
	// like an explicit conversion
	fmt.Println(int64(d))

	//A number can be given a type by using in a context that requires one, such as a variable assignment 
	// or function call.
	//In this case sin expects an float64
	fmt.Println(math.Sin(n))

}

