/**
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 3 de Lenguages Imperativos 
**/

package main

import "fmt"


func rotar( list *[]string, rotacion int, cant int ){

	//Por cada rotación
	for j := 0; j < cant; j++ {

		//Guarda los extremos del slice
		last := (*list)[len(*list)-1]
		first := (*list)[0] 

			//Hace la rotación
			for i := 0; i < len(*list)-1; i++ {

				if rotacion == 1 {				 //Rotación derecha
					temp := (*list)[i+1]
					(*list)[i+1] = first
					first = temp
				}else{										 //Rotación izquierda
					(*list)[i] = (*list)[i+1]
				}
			}

			//Asigno el primer o ultimo elemento 
			if rotacion == 1 {
				(*list)[0] = last		
			}else	{
				(*list)[len(*list)-1] = first
			}
	}
}

func main(){

	t := []string{"a","b","c","d","e","f","g","h"}

	fmt.Println("Secuencia original: ")
	fmt.Print(t)

	//rotar(slice, (izq=0, der=1), cantidad rotaciones)
	rotar(&t, 0, 3)

	fmt.Println("Secuencia final rotada: ")
	fmt.Println(t)
	
}