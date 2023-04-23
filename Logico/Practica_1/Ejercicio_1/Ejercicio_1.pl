/*
Instituto Tecnologico de Costa Rica
Escuela Computacion 
Sede San Carlos 

Curso: Lenguages de Programacion
Profesor: Oscar Mario Viquez Acuna

Estudiante: Josue Chaves Araya - 2015094068

Ejercicio 1 de Lenguage Lógico 

Implemente  un  predicado  que,  a  partir  de  una  lista  de  cadenas  de  parámetro,  filtre 
aquellas que contengan una subcadena que el usuario indique en otro argumento.

Ej: 
 
 sub_cadenas(“la”, [“la casa, “el perro”, “pintando la cerca”],Filtradas). 
 True 
 Filtradas = [“la casa, “pintando la cerca”]

*/

sub_cadenas(_, [], []).                      % Si la subcadena, no se puede comparar con nada más entonces []
sub_cadenas(SUBSTR, [STR|T], [STR|LIST]) :-  % Una cadena puede pertenecer a la lista, sí y sólo sí :
  sub_atom(STR, _, _, _, SUBSTR),            % Una cadena es subcadena de otra cadena
  sub_cadenas(SUBSTR, T, LIST),!.            % Una cadena es subcadena del resto de cadenas 
sub_cadenas(SUBSTR, [_|T], LIST) :-          % Para las otras cadenas que no sean subcadenas de otras cadenas,
  sub_cadenas(SUBSTR, T, LIST).              %  sólo se ignoran. 
