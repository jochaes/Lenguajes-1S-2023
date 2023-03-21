printfn "Hello World from F#" //Se llama a la funcieon printfn y le mando argumentos

let lista = [3;4;5;6;7] //Debe ser homogenea y es inmutable... También se puede hacer let lista = [1..10]


let cuadrado = (fun x -> x*x)

//Sistanxis de patrones 

let listaMap =
    lista
    |> List.map cuadrado
    |> List.filter (fun x -> (x % 2) = 0)

//Crear funcion con parametro 

let FuncionMap (lista: int List) =
    lista
    |> List.map cuadrado
    |> List.filter (fun x -> (x % 2) = 0)


let rec factorial x = //Hay que poner rec para "activarle" que pueda utilizar recursion
  if( x = 0 ) then 
    1
  elif (x = 1) then 
    1
  else 
    x * (factorial (x-1))


let x = cuadrado 10


printfn "El valor entero es: %A" (factorial 10)


