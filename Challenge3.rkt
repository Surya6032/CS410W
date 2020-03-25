#lang racket
(define (values hours rate)
  (if (> hours 40)
      (+(* 40 rate) (* rate (* 1.5 (- hours 40))))
      (* hours rate)))
  
(values 20 8)

(values 50 8)