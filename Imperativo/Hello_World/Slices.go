package main

import "fmt"

func main() {

	//Slice of strings of length 3, zero-valued
	s := make([]string, 3)
	fmt.Println("emp:", s)

	//Set and get just like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	//Length
	fmt.Println("len", len(s))

	//Append returns a slice containing one or more new values
	// Note that we need to accept a return value from append as we may
	// get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	//Copy a slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	//Slice operator 

	l := s[2:5]  
	fmt.Println("sl1: ", l)

	l = s[:5]  
	fmt.Println("sl2: ", l)

	l = s[2:]  
	fmt.Println("sl3: ", l)

	//Declare and initialice a variable for slice in a single line
	t:= []string{"g", "h", "i"}
	fmt.Println("dlc: ", t)

	//Slices can be composed into multi-dimensional data structures. 
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.

	twoD := make([][]int, 3)
	for i:= 0; i < 3; i++{
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}