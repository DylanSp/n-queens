% load arithmetic and tools for working with finite domains
:- use_module(library(clpfd)).

not_threatened_on_row(Queen1Row-_Queen1Col, Queen2Row-_Queen2Col) :- Queen1Row #\= Queen2Row.

not_threatened_on_col(_Queen1Row-Queen1Col, _Queen2Row-Queen2Col) :- Queen1Col #\= Queen2Col.

not_threatened_on_diag(Queen1Row-Queen1Col, Queen2Row-Queen2Col) :- abs(Queen1Row - Queen2Row) #\= abs(Queen1Col - Queen2Col).

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

queen_valid_for_solution(N, QueenRow-QueenCol) :-
  QueenRow #>= 0,
  QueenRow #< N,
  QueenCol #>= 0,
  QueenCol #< N.

label_queen(QueenRow-QueenCol) :- label([QueenRow, QueenCol]).

solution(N, SortedQueens) :-
  length(Queens, N),
  is_set(Queens),
  maplist(queen_valid_for_solution(N), Queens), % make sure all queen values are valid for size N
  no_queens_threaten(Queens),
  maplist(label_queen, Queens), % "ground" our solution in actual values.
  sort(Queens, SortedQueens).

solutions_without_duplicates(N, SolutionsList) :- setof(RawSolutions, solution(N, RawSolutions), SolutionsList).

% "all" column in https://en.wikipedia.org/wiki/Eight_queens_puzzle#Counting_solutions_for_other_sizes_n
% counts reflections/rotations
count_all_solutions(N, Count) :-
  solutions_without_duplicates(N, SolutionsList),
  length(SolutionsList, Count).

fundamental_solutions(N, FundamentalSolutions) :-
  solutions_without_duplicates(N, DeduplicatedSolutions).
  % TODO - eliminate reflections/rotations
