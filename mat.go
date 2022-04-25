package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

var (
	x = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	y = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	degree = 1
)

func main() {
	a := Vandermonde(x, degree)
	b := mat.NewDense(len(y), 1, y)
	c := mat.NewDense(degree+1, 1, nil)

	var qr mat.QR
	qr.Factorize(a)

	const trans = false
	err := qr.SolveTo(c, trans, b)
	if err != nil {
		log.Fatalf("could not solve QR: %+v", err)
	}
	fmt.Printf("%.3f\n", mat.Formatted(c))
}

func Vandermonde(a []float64, degree int) *mat.Dense {
	d := degree + 1
	x := mat.NewDense(len(a), d, nil)
	for i := range a {
		for j, p := 0, 1.; j < d; j, p = j+1, p*a[i] {
			x.Set(i, j, p)
		}
	}
	return x
}
