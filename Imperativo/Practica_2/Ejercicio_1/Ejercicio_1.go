package main
//TEC - Lenguajes de Programacion ->Ejercicio 1, practica 2, puntos a y b, Josué Chaves Araya - 2015094068
import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}

type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	
	var prod = l.buscarProducto(nombre) //Se fija a ver si existe el producto

	fmt.Println("Agregando ", nombre)
	if prod != -1 {
		(*l)[prod].cantidad += cantidad //Suma la nueva cantidad 
		fmt.Println("Ya existe ",nombre , " ahora hay: ", (*l)[prod].cantidad , " productos")
		if (*l)[prod].precio != precio {
			fmt.Println("El precio del ",nombre," es diferente, el nuevo precio es: ", precio)
			(*l)[prod].precio = precio //Cambia el precio cuando es diferente 
		}
	}else{
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})	
	}
	
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

//Tambien agregar la cantidad de productos que quiero vender 
func (l *listaProductos) venderProducto(nombre string, cantidad int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 {
		(*l)[prod].cantidad = (*l)[prod].cantidad - cantidad
		fmt.Println("Se vendieron: ", cantidad, " ", nombre, " Stock: ", (*l)[prod].cantidad )

		if (*l)[prod].cantidad <= 0 {
			fmt.Println("Ya no hay producto, borrando el producto")
			*l = append((*l)[:prod], (*l)[prod+1:]...)  //Borra el producto 
		}
		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar
	}
}


func (l *listaProductos) listarProductosMinimos() listaProductos {

	var prodMin listaProductos

	for _, prod := range *l {
		if prod.cantidad <= existenciaMinima{
			prodMin = append(prodMin, prod)
		}
	}

	if len(prodMin) == 0 {
		return nil 
	}else{
		return prodMin
	}
	// debe retornar una nueva lista con productos con existencia mínima
	// Devuele una nueva lista con productos con existencia minima

}

// ************  Practica 2  ************

//Practica 2 - Ejercicio 1 a
func (l *listaProductos) aumentarInventarioDeMinimos(min listaProductos)  {

	for _,prodMin := range min{
		cantidadAgregar := (existenciaMinima - prodMin.cantidad) + 1  // A cada producto le suma la cantidad minima + 1
		l.agregarProducto(prodMin.nombre, cantidadAgregar,prodMin.precio)
	}
}

//Practica 2 - Ejercicio 1 c 

type By func(p1, p2 *producto) bool //Tipo que especifica la busqueda

//funcion que ordena los productos dependiendo del tipo
func (by By) Sort(pProductos listaProductos) {
	ps := &ordenadorProductos{
		productos: pProductos,
		by: by, 
	}
	sort.Sort(ps)
}

//Struct para ordenar
type ordenadorProductos struct {
	productos listaProductos       //La lista de productos 
	by func(p1, p2 *producto) bool //La función que dice cómo los va a ordenar
}

//Funciones utilitarias para hace un sort generico
func (s *ordenadorProductos) Len() int {
	return len(s.productos)
}

func (s *ordenadorProductos) Swap(i, j int) {
	s.productos[i], s.productos[j] = s.productos[j], s.productos[i]
}

func (s *ordenadorProductos) Less(i, j int) bool {
	return s.by(&s.productos[i], &s.productos[j])
}

//Ordena por cualquier elemento del struct de productos 
func (l *listaProductos) ordenarPor(key string ){

	switch key {
	case "nombre":
			fmt.Println("Ordenando por Nombre")
			nombre := func(p1, p2 *producto) bool { //Funcion que compara los nombres 
				return p1.nombre < p2.nombre
			}
			By(nombre).Sort(*l) 										//Ordena la lista por la funcion de nombre

	case "cantidad":
		fmt.Println("Ordenando por Cantidad")
		cantidad := func(p1, p2 *producto) bool{ //Funcion que compara las cantidades  
			return p1.cantidad < p2.cantidad
		}
		By(cantidad).Sort(*l) 										//Ordena la lista por la funcion de cantidad
	case "precio":
		fmt.Println("Ordenando por Precio")
		precio := func(p1, p2 *producto) bool{ //Funcion que compara los precios de los productos  
			return p1.precio < p2.precio
		}
		By(precio).Sort(*l) 										//Ordena la lista por la funcion de precio

	default:
		fmt.Println(key, " no corresponde a ningún elemento de la estructura")
	}
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
}

func main() {
	llenarDatos()

	//Imprime la lista de productos Original
	fmt.Println("Lista de Productos Original:")
	fmt.Println(lProductos)
	fmt.Println()

	//Imprime la lista de productos Mínimos
	prodMin := lProductos.listarProductosMinimos()
	fmt.Println("Lista prod minimos: ", "\n", prodMin)
	fmt.Println()

	//Aumenta el inventario para que ya no hayan productos mínimos
	//Ejercio 1 a
	fmt.Println("Aumentando el inventario de productos minimos")
	lProductos.aumentarInventarioDeMinimos(prodMin)
	fmt.Println("Lista de Productos Actualizados:", "\n", lProductos)
	fmt.Println()


	//Ordenar 
	//Ejercio 1 b
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Lista Productos:")
	fmt.Println(lProductos)
	fmt.Println()
	lProductos.ordenarPor("nombre")
	fmt.Println(lProductos)
	fmt.Println()
	lProductos.ordenarPor("cantidad")
	fmt.Println(lProductos)
	fmt.Println()
	lProductos.ordenarPor("precio")
	fmt.Println(lProductos)
	fmt.Println()

}
