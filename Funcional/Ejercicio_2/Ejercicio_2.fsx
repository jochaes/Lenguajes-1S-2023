(*
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 2 de Lenguages Funcionales 

Construya una función que se llame general-sec A B.
  Esta función genera una secuencia de números en una lista de A hasta B,
  donde A y B son números enteros.

*)


module Ejercicio_2 =
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


open Ejercicio_2

let frases = ["otra cosa";"la rueda";"mejor la otra";"no se cuala";"la"] 
printfn "%A" (filtrarPalabras "la" frases)

