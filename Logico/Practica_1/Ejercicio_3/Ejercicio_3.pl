/*
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 2 de Lenguage Lógico 


Implemente  la  función  aplanar.  Esta  función  recibe  una  lista  con  múltiples  listas 
anidadas como elementos y devuelve una lista con los mismos elementos de manera lineal 
(sin listas).
Ej:  
  ?-aplanar([1,2,[3,[4,5],[6,7]]],X). 
  True 
  X=[1,2,3,4,5,6,7]. 
*/


aplanar([],[]).                % Prolog una lista vacía es una lista vacía?
aplanar(H|T, LISTA) :-         % Si es una lista entonces:
  aplanar(H, HLISTA),          % Aplananos H en HLISTA
  aplanar(T, TLISTA),          % Aplanamos T en TLista
  append(HLISTA,TLISTA, LISTA).% Concatenamos las listas planas HLISTA y TLISTA en LISTA
aplanar(ELEM, [ELEM]).         % Si es un elemento, devuelve una lista con ese elemento.

