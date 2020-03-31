#lang racket

(define (zerocount numbers)
  (cond
    ((null? numbers)0)
    ((zero? (car numbers))(+ 1 (zerocount (cdr numbers))))
    (else (zerocount(cdr numbers)))))
  
(zerocount '(7 0 99 0 -8 44 0))