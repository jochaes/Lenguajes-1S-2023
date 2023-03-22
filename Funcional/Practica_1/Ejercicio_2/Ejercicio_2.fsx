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

  let generar_sec A B =
    if A < B then [for i in A .. B -> i]
    else []
    
  

open Ejercicio_2
open System

let args:string[] = Environment.GetCommandLineArgs()

let mutable A = 0
let mutable B = 0

//printfn "%A" args

if Int32.TryParse(args[(args.Length - 2)], &A) &&  Int32.TryParse(args[(args.Length - 1)], &B) then printfn "Secuencia Generada: %A" (generar_sec A B)
else printfn "Error tratando de parsear los numeros A y B \n debe agregar el rango de la siguiente manera:\n $ dotnet fsi Ejercicio_2.fsx A B"
