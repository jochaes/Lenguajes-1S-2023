Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Practica 1 Ejercicio 4 de Lenguages Imperativos 


Instrucciones para ejecutar el programa: 

- En la función main del archivo "Ejercicio_4.go" está el código de prueba, puede cambiarlo y jugar con las diferentes funciones del slice
-Las funciones sobre el inventario son: 
- agregarCalzado(marca, precio, talla, stock)
- buscarCalzado(marca, talla)
- venderCalzado(marca, talla, cantidad)

- En el código de main, pasa lo siguiente: 
linea 130 -> Se crean los datos de prueba, ingresando diferentes marcas y tallas de calzado
linea 136 -> Se busca un calzado especifico 
linea 140 -> Se intenta vender una cantidad mayor al stock de un calzado, mostrando un mensaje de eror 
linea 143 -> Se vende la mitad del stock de un calzado
linea 147 -> Se vende la otra mitad del stock de un calzado, dejando el stock en 0
linea 150 -> Se vende un par del calzado, lo cual muestra un mensaje de error, y que la venta no puede suceder

- Para ejecutar el programa 
$ go run Ejercicio_4.go



