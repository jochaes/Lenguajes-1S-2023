/*
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 2 de Lenguage Lógico 


Construya una función que se llame sub_conjunto. Esta función recibe dos listas y debe 
retornar True cuando el primer argumento es subconjunto completo del segundo y #f en 
caso contrario. 

Por ejemplo:  
  sub_conjunto([],[1,2,3,4,5]). 
  True 
  sub_conjunto([1,2,3],[1,2,3,4,5]). 
  True 
  sub_conjunto([1,2,6],[1,2,3,4,5]). 
  False 
*/


sub_conjunto([],_).          % Un conjunto vacío es subconjunto de cualquier conjunto 
sub_conjunto([H|T], R) :-    % Cada elemento de [H|T] es subconjunto de R sí y sólo sí: 
  member(H, R),             % H es miembro de R
  sub_conjunto(T, R),!.      % y el resto de los elementos de T también son subconjunto de R.  
