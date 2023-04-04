(*
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 1 de Lenguages Funcionales - Practica 2

*)

module Ejercicio_1 =

  let grafo = [
    [('i', 0);('a', 1);('b', 2)];
    [('a', 1);('i', 0);('c', 3);('d', 2)];
    [('b', 2);('i', 0);('c', 1);('d', 4)];
    [('c', 3);('a', 1);('b', 1);('x', 2)];
    [('d', 2);('a', 1);('b', 4);('f', 3)];
    [('f', 3);('d', 2)];
    [('x', 2);('c', 3)]
]

let miembro e (lista: ('a * 'b) list) =
    lista
    |> List.map (fun (x,_) -> x = e)
    |> List.reduce (fun x y -> x || y)

let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | ((head, peso)::adyacencias)::tail when nodo = head -> adyacencias
    | _::tail -> vecinos nodo tail 


let extender (ruta: _ list) grafo =
    (vecinos (fst ruta.Head) grafo)
    |> List.map (fun (x,peso) -> if (miembro x ruta) then [] else (x,peso)::ruta )
    |> List.filter (fun x -> x <> [])

let rec prof_aux fin (rutas: _ list list) grafo =
    if rutas = [] then
        []
    elif fin = (fst rutas.Head.Head) then
        //List.rev rutas.Head
        List.append
            ([List.rev rutas.Head])
            (prof_aux fin rutas.Tail grafo)
    else
        prof_aux fin (List.append (extender rutas.Head grafo) rutas.Tail) grafo
        
let prof ini fin grafo =
    prof_aux fin [[(ini,0)]] grafo
    


//**********************************************************************
(*
   Funciones que corresponden al ejercicio 2
*)

//Calcula el peso total de una ruta
let rec pesoRuta (lista: ('a * int) list) =
    match lista with
    | [] -> 0
    | (x,y)::tail -> y + (pesoRuta tail)

//Da la ruta con menor peso
let menorRuta (lista: ('a * int) list list) =
    lista
    |> List.reduce (fun x y -> if pesoRuta x < pesoRuta y then x else y)


open Ejercicio_1

printf "Grafo: \n%A" grafo
printfn "\n"
printfn "Rutas ( i->f ): \n%A" (prof 'i' 'f' grafo)
printfn "\n"
printf "Menor Ruta ( i->f ):  %A" (menorRuta (prof 'i' 'f' grafo))
