/**
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 2 de Lenguages Imperativos 
**/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func draw(centerLine int){

	middle := centerLine/2
	ast := -1

	for i := middle; i > (middle*-1)-1; i-- {

		str := ""
		spaces := i

		if spaces < 0 {
			spaces *=-1	
			ast -= 2 
		}else{
			ast+=2
		}

		for j := 0; j < spaces; j++ {
			str +="  "
		}

		for j := 0; j < ast; j++ {
			str += "* "
		}

		fmt.Println(str)
	}

}


func main(){

	centerLineElemStr := os.Args[1]
	

	//Convert command line argument to Integer
	centerLineElem, err := strconv.Atoi(centerLineElemStr)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	if centerLineElem <=0 || centerLineElem%2 != 1 {
		fmt.Println("Input must be an odd number greater than 0")
		return
	}

	fmt.Println("Drawing figure with center line of:",centerLineElem)
	draw(centerLineElem)

}