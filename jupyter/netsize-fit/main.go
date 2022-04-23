package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

func main() {
	f, err := os.Open("data2.csv")
	if err != nil {
		fmt.Println("Unable to read input file", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse file as CSV for ", err)
	}

	grouped := map[int][]float64{}

	xs := []float64{}
	ys := []float64{}
	for _, record := range records {
		x, _ := strconv.ParseFloat(record[0], 64)
		y, _ := strconv.ParseFloat(record[1], 64)
		xs = append(xs, x)
		ys = append(ys, y)

		order := int(x)
		if _, ok := grouped[order]; !ok {
			grouped[order] = []float64{}
		}
		grouped[order] = append(grouped[order], y)
	}
	orders := []int{}
	for order := range grouped {
		orders = append(orders, order)
	}
	sort.Ints(orders)

	ws := []float64{}
	for _, order := range xs {
		ws = append(ws, 1.0/stat.StdDev(grouped[int(order)], nil))
	}

	alpha, beta := stat.LinearRegression(xs, ys, ws, true)
	fmt.Println("alpha: ", alpha)
	fmt.Println("beta: ", beta)

	ws = []float64{}
	xs = []float64{}
	for _, order := range orders {
		xs = append(xs, float64(order))
		ws = append(ws, 1.0/stat.StdDev(grouped[order], nil))
	}

	lhs := Vandermonde(xs, 1)

	weights := mat.NewDense(2, len(xs), append(ws, ws...))

	lhs.MulElem(lhs, weights.T())

	var scale mat.Dense
	scale.MulElem(lhs, lhs)

	row := []float64{math.Sqrt(mat.Sum(scale.ColView(0))), math.Sqrt(mat.Sum(scale.ColView(1)))}
	scales := []float64{}
	for range xs {
		scales = append(scales, row...)
	}
	scale = *mat.NewDense(len(xs), 2, scales)
	lhs.DivElem(lhs, &scale)

	var c mat.Dense
	c.Mul(lhs.T(), lhs)

	var i mat.Dense
	i.Inverse(&c)

	outer := mat.NewDense(2, 2, []float64{
		scale.At(0, 0) * scale.At(0, 0), scale.At(0, 0) * scale.At(0, 1),
		scale.At(1, 0) * scale.At(0, 0), scale.At(1, 0) * scale.At(0, 1),
	})

	fmt.Println(mat.Formatted(outer))

	i.DivElem(&i, outer)

	fmt.Println(math.Sqrt(i.At(0, 0)))
}

func Vandermonde(a []float64, degree int) *mat.Dense {
	d := degree + 1
	x := mat.NewDense(len(a), d, nil)
	for i := range a {
		for j, p := d-1, 1.; j >= 0; j, p = j-1, p*a[i] {
			x.Set(i, j, p)
		}
	}
	return x
}
