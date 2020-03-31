#lang racket
(define (average lst)
  (define newList (remove* (list (apply min lst)) lst))
  (define newList1 (remove* (list (apply max newList)) newList))
  (/ (apply + newList1) (length newList1)))
(average '(7 6 3 10 9))