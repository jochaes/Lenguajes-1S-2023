/**
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 1 de Lenguages Imperativos 

Haga un programa que cuente e indique el numero de caracteres, el numero de palabras y 
el numero de lineas de un texto cualquiera (obtenido de cualquier forma que considere). 
Considere hacer el calculo de las tres variables en el mismo proceso. 

**/

package main
import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {

    //Variable Initialization
    var numCharacters, numWords, numLines int

    //Opens a file 
    f, err := os.Open("text.txt")

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)

    fmt.Println("Printing Text: ")
    
    //Goes line by line
    for scanner.Scan() {
        numLines ++                                         //Every line
        numCharacters += len(scanner.Text())                //Every character in the line
        numWords += len(strings.Split(scanner.Text(), " ")) //Every word in the line
        fmt.Println(scanner.Text())
    }
    numCharacters += (numLines - 1)

    fmt.Println("Result (The number of characters includes new lines and spaces) : ")
    fmt.Println("Chars: ", numCharacters)
    fmt.Println("Words: ", numWords)
    fmt.Println("Lines: ", numLines)

    //Closes the file 
    //Can also be done using defer at the begining of the code
    f.Close() 
}

