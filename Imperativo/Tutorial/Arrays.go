package main

import "fmt"

//Arrays have static lenght 

func main() {

	//By default an array is zero-valued, wich for 
	// ints means 0s
	var a [5] int
	fmt.Println("emp:", a)

	//Get and set using the array's index 
	a[4] = 100
	fmt.Println("set:",a)
	fmt.Println("set:",a[4])

	fmt.Println("len:", len(a))

	//Declare and Initialize an array
	b := [5] int {1,2,3,4,5}
	fmt.Println("dcl:", b) 

	//Array types are one-dimensional, but you can compose 
	// types to build multi-dimensional data structures.

	var twoD [2][3] int
	for i:= 0; i < 2; i++{
		for j :=0; j < 3;j++{
			twoD[i][j] = i+j
		}
	}

	fmt.Println("2d: ", twoD)
}