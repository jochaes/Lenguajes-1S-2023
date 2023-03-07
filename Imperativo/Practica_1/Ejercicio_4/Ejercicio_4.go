/**
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 4 de Lenguages Imperativos 
**/

package main

import (
	"fmt"
	"math/rand"
	"time"
	"golang.org/x/exp/slices"
)

/**
Genera un numero random 
**/
func generarId() int {

	s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

	return r1.Intn(1000)

}

type calzado struct {
	id int 				//Solo para no repetir toda la info siempre, va a calcularse de manera aleatoria
	marca string
	precio int 
	talla int 
	stock int
}

const TALLAMINIMA = 34
const TALLAMAXIMA = 44

type inventario []calzado

var zapateria inventario

/**
	Agrega calzados al inventario, si ya existe el calzado, entonces aumenta su stock
**/
func (l *inventario) agregarCalzado( marca string, precio int, talla int, stock int){

	if  t := (*l).buscarCalzado(marca, talla); t != -1 {
		stockOriginal := (*l)[t].stock
		(*l)[t].stock += stock 															//Si el calzado ya existe entonces sume el stock 
		fmt.Println("Ya existen",stockOriginal, "pares de", marca, "de talla", talla,". Aumentado el stock a:", (*l)[t].stock, "pares." )		
	}else{
		if talla < TALLAMINIMA || talla > TALLAMAXIMA {
			fmt.Println("Solo se pueden agregar tallas entre: ", TALLAMINIMA, " y ", TALLAMAXIMA)
		}else{
			fmt.Println("Agregando", stock, "pares de", marca, "de talla", talla, "al inventario. Precio:", precio)
			*l = append(*l, calzado{generarId(), marca, precio, talla, stock})
		}
	}
}

/**
	Retorna el indice del calzado que tenga misma marca y talla
**/
func (l *inventario) buscarCalzado(pMarca string, pTalla int) int {

	comp := func( cal calzado ) bool {

		return cal.marca == pMarca && cal.talla == pTalla
	}

	return slices.IndexFunc(*l, comp )
}

/**
	Vende una cantidad de articulos del stock, si el articulo existe.
**/
func (l *inventario) venderCalzado(pMarca string, pTalla int, pCantidad int){

	if idx := (*l).buscarCalzado(pMarca, pTalla); idx != -1{
		st := (*l)[idx].stock
		if st == 0 {
			fmt.Println("No se puede vender, ya no hay stock de",(*l)[idx].marca, "talla", (*l)[idx].talla )
		}else if st - pCantidad < 0 {
			fmt.Println("Solo hay" ,(*l)[idx].stock, "pares", "no se pueden vender", pCantidad)
		}else{
			(*l)[idx].stock -= pCantidad
			fmt.Println("Se vendieron:", pCantidad, "pares de", (*l)[idx].marca, "talla",(*l)[idx].talla)
			fmt.Println("El stock es de:",(*l)[idx].stock)
			fmt.Println()
		}
	}else{
		fmt.Println("El calzado no existe")
	}
}

func crearDatos(){
	zapateria.agregarCalzado("Nike", 50000, 35, 10)
	zapateria.agregarCalzado("Nike", 50000, 36, 10)

	zapateria.agregarCalzado("New Balance", 45000, 44, 20)
	zapateria.agregarCalzado("New Balance", 45000, 43, 10)

	zapateria.agregarCalzado("Converse", 60000, 42, 5)
	zapateria.agregarCalzado("Converse", 60000, 37, 15)

	zapateria.agregarCalzado("Puma", 35000, 44, 10)
	zapateria.agregarCalzado("Under_Armour", 20000, 40, 20)
	zapateria.agregarCalzado("adidas", 75000, 43, 10)
	zapateria.agregarCalzado("Reebok", 90000, 39, 35)
	zapateria.agregarCalzado("Vans", 20000, 40, 10)
	zapateria.agregarCalzado("Prada", 1000000, 35, 10)
	zapateria.agregarCalzado("LV", 2000000, 36, 5)

	
	fmt.Println("Imprimiento inventario:")
	fmt.Println(zapateria)
	fmt.Println()

}

func main(){
	crearDatos()

	//Agregando otros 50 pares de Nike Talla 35
	zapateria.agregarCalzado("Nike", 50000, 35, 50)
	
	//Buscando el calzado de Nike talla 35
	indice := zapateria.buscarCalzado("Nike", 35)
	fmt.Println(zapateria[indice])
	
	//Vendiendo 200 pares de Nike talla 35 
	zapateria.venderCalzado("Nike", 35,200)

	//Vender 30 pares de Nike talla 35
	zapateria.venderCalzado("Nike", 35,30)


	//Vender 30 pares de Nike talla 35
	zapateria.venderCalzado("Nike", 35,30)

	//Vender 1 pares de Nike talla 35
	zapateria.venderCalzado("Nike", 35,1)

	//Buscando el calzado de Nike talla 35
	indice = zapateria.buscarCalzado("Nike", 35)
	fmt.Println(zapateria[indice])

}

