miembro(N,[N|_]).
miembro(N,[_|Tail]):- miembro(N,Tail),!.

factorial(0,1).
factorial(A,B) :- 
		A > 0,
		C is A-1,
		factorial(C,D),
		B is A*D.


insertar(E,[],[E]). 
insertar(E,[H|T],[E|[H|T]]):-  E < H,!.
insertar(E,[H|T],[H|R]):- insertar(E,T,R).

eliminar(X, [], []).
eliminar(X, [X|T], R):- eliminar(X, T, R), !.
eliminar(X, [Y|T], [Y|R]):- eliminar(X, T, R).