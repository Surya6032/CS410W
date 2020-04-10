male(surya).
male(nitish).
male(rajender).
male(ravinder).

female(nidhi).
female(alka).
female(anita).

parents(surya,rajender,alka).
parents(nidhi,rajender,alka).
parents(nitish,ravinder,anita).

parents(rajender,balbir,santosh).
parents(ravinder,balbir,santosh).

sister_of(X,Y) :- female(X),parents(X,Mother,Father),parents(Y,Mother,Father),X\=Y.
brother_of(X,Y):- male(X),parents(X,Mother,Father),parents(Y,Mother,Father),X\=Y.
