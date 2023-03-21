(*
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 1 de Lenguages Funcionales 

Haciendo uso de la función filter, implemente una función que a partir de una
  lista de cadenas representando frases de parámetro, filtre aquellas que
  contengan una palabra que el usuario indique en otro argumento
  (palabra completa contenida).

*)

module Ejercicio_1 =
  //Función recursiva que se encarga de buscar un match en la lista 
  // ele: elemento a buscar en la lista
  // list: Lista en dónde puede estar el elemeto
  let rec checkaux ele list =
    match list with    
          | [] -> false
          | head :: tail when ele = head -> true
          | _ :: tail -> checkaux ele tail

  //Filtrar textos por contenido de palabras
  // pal: Palabra a filtrar
  // lista: Lista de frases a filtrar que pueden contener pal
  let filtrarPalabras (pal:string) (lista:string list) =
      lista
      |> List.filter (fun str -> checkaux pal (Array.toList (str.Split())) )


open Ejercicio_1

let frases = ["otra cosa";"la rueda";"mejor la otra";"no se cuala";"la"] 
printfn "%A" (filtrarPalabras "la" frases)