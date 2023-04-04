(*
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 3 de Lenguages Funcionales 

Migrar el ejercicio realizado en Go sobre facturas con listas de productos haciendo énfasis 
en  la  implementación  de  las  funciones  para  calcular  el  monto  a  pagar  por  la  factura 
completa y el monto a pagar por concepto del 13% de impuesto de venta para aquellos 
productos que deban pagarlo según criterio de selección. 
*)

module Ejercicio_3 = 
  
  //Objeto que tiene nombre y precio
  type Producto = {Nombre: string; Precio: float; Impuesto: bool}

  //Lista de productos
  let factura:list<Producto> = 
    [
      {Nombre = "Coca Cola"; Precio = 1500; Impuesto = true}
      {Nombre = "Pepsi"; Precio = 2000; Impuesto = true}
      {Nombre = "Fanta"; Precio = 2500; Impuesto = true}
      {Nombre = "Sprite"; Precio = 1500; Impuesto = true}
      {Nombre = "Manzana"; Precio = 750; Impuesto = false}
      {Nombre = "Pera"; Precio = 550; Impuesto = false}
      {Nombre = "Sandia"; Precio = 800; Impuesto = false}
      {Nombre = "Melon"; Precio = 1200; Impuesto = false}
      {Nombre = "Naranja"; Precio = 300; Impuesto = false}
    ]

  //Funcion para recorrer todos los elementos de una lista de producto
  let rec recorrerListaProductos (lista: Producto list) = 
    match lista with
    | [] -> printfn "No hay productos"
    | h::t -> printfn "%s - %f" h.Nombre h.Precio; recorrerListaProductos t

  // Retorna true si el producto tiene impuesto
  let tieneImpuesto (producto: Producto) = 
    producto.Impuesto

  // Retorna una lista con los impuestos de los productos que tienen impuesto
  let productosConImpuesto (lista: Producto list) = 
    (List.map (fun x -> x.Precio * 0.13) (List.filter tieneImpuesto lista) )
  
  // Retorna el total de los impuestos 
  let totalImpuestos (lista: Producto list) = 
    List.reduce (fun x y -> x + y) (productosConImpuesto lista)


  // Retorna el precio total de los productos que no tienen impuesto
  let precioTotal (lista: Producto list) = 
    List.reduce (fun x y -> x + y) (List.map (fun x -> x.Precio) lista)



open Ejercicio_3

let total = precioTotal factura
let impuestos = totalImpuestos factura

printfn "Precio Total sin Impuesto: %f"  total
printfn "Impuestos: %f"  impuestos
printfn "Precio Total con Impuesto: %f"  (total + impuestos)



  


