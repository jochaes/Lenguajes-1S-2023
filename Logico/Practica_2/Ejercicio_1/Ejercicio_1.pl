/*
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 1 Practica 2  de Lenguage Lógico 


Explicación del Ejercicio: 

https://www.youtube.com/watch?v=OfeJMA93M8Y
*/

% Laberinto
conectado(i,2).
conectado(2,3).
conectado(2,8).
conectado(8,9).
conectado(9,3).
conectado(3,4).
conectado(4,10).
conectado(10,16).
conectado(16,22).
conectado(22,21).
conectado(21,15).
conectado(15,14).
conectado(14,13).
conectado(13,7).
conectado(7,1).
conectado(14,20).
conectado(20,26).
%conectado(22,28).
conectado(26,27).
conectado(27,28).
conectado(28,29).
conectado(29,23).
conectado(23,17).
conectado(17,11).
conectado(11,5).
%conectado(1,7). %por eliminar
conectado(5,6).
conectado(28,34).
conectado(34,35).
conectado(35,36).
conectado(36,30).
conectado(30,24).
conectado(24,18).
conectado(18,12).
conectado(32,31).
conectado(31,25).
conectado(25,19).
conectado(34,33).
conectado(33,32).
conectado(32,f).


conectados(X,Y):- conectado(X,Y).
conectados(X,Y):- conectado(Y,X).

rutaLaberinto(Fin, Fin, Camino, Res):-
  reverse([Fin|Camino], Res).  %Reversa el resultado obtenido 

rutaLaberinto(Ini,Fin,Camino,Res):-    %Es solucion sí 
  conectados(Ini,Z),              % Ini está conectado con Z        - Si Z ya estuviera, acá se devuelve a buscar otro conectado
  not(member(Z,Camino)),          % Z no es miembro del Camino
  rutaLaberinto( Z, Fin, [Ini|Camino], Res). % Se puede ir de Z a fin 

solucionLaberinto(_,_,Result) :-
  findall(Res, rutaLaberinto(i,f,[],Res), Result).
