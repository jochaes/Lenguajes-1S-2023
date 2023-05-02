Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068


Practica 2 Ejercicio 1 de Lenguage Lógico 

Laberinto - Video: https://www.youtube.com/watch?v=OfeJMA93M8Y

IDE: VS Code 
Lenguaje: Prolog 



solucionLaberinto(Ini,Fin,Result).
 Devuelve en Result todas las rutas que puede tomar dentro del laberinto para ir desde Ini hasta fin

Instrucciones para ejecutar el programa: 

- Para iniciar swipl desde consola y cargar el programa: 

$swipl -f Ejercicio_1.pl

- Una vez dentro de la consola de swipl, puede ejecutar esta llamadas de prueba: 
  ?- solucionLaberinto(i,f,Result).
                       | |  └>Resultado
                       | └> A dónde se quiere ir dentro del laberinto (Salida)
                       └> Casilla de Inicio



Por ejemplo: 
?- solucionLaberinto(i,f,Result) ; true.
Result = [[i, 2, 3, 4, 10, 16, 22, 21|...], [i, 2, 8, 9, 3, 4, 10|...]] [write]
Result = [[i, 2, 3, 4, 10, 16, 22, 21, 15, 14, 20, 26, 27, 28, 34, 33, 32, f], [i, 2, 8, 9, 3, 4, 10, 16, 22, 21, 15, 14, 20, 26, 27, 28, 34, 33, 32, f]] 