#lang racket
(define (member a lst)
  (cond
    ((null? lst) #f)
    ((eq? a (car lst)) #t)
    (else (member a (cdr lst)))))
(member 9 '(1 2 3))
    