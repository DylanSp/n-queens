% load arithmetic
:- use_module(library(clpfd)).

not_threatened_on_row([Queen1Row, _Queen1Col], [Queen2Row, _Queen2Col]) :- Queen1Row #\= Queen2Row.

not_threatened_on_col([_Queen1Row, Queen1Col], [_Queen2Row, Queen2Col]) :- Queen1Col #\= Queen2Col.

% not_threatened_on_diag([Queen1Row, Queen1Col], [Queen2Row, Queen2Col]) :- abs(Queen1Row - Queen2Row) #\= abs(Queen1Col - Queen2Col).
not_threatened_on_diag([Queen1Row, Queen1Col], [Queen2Row, Queen2Col]) :-
  (Queen1Row - Queen2Row) #\= (Queen1Col - Queen2Col),
  (Queen1Row - Queen2Row) #\= (-1 * (Queen1Col - Queen2Col)).


not_threatened(Queen1, Queen2) :-
  not_threatened_on_row(Queen1, Queen2),
  not_threatened_on_col(Queen1, Queen2),
  not_threatened_on_diag(Queen1, Queen2).
  
no_queens_threaten([]).
no_queens_threaten([_SingleQueen]).
no_queens_threaten([Queen1|[Queen2|Queens]]) :-
  not_threatened(Queen1, Queen2),
  no_queens_threaten([Queen1|Queens]),
  no_queens_threaten([Queen2|Queens]).

queen_valid_for_solution(N, [QueenRow, QueenCol]) :-
  QueenRow #>= 0,
  QueenRow #< N,
  QueenCol #>= 0,
  QueenCol #< N.

solution(N, Queens) :-
  length(Queens, N),
  is_set(Queens),
  maplist(queen_valid_for_solution(N), Queens),
  no_queens_threaten(Queens),
  maplist(label, Queens). % "ground" our solution in actual values.
