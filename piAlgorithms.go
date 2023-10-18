package pialgorithms

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func CalcRatio(sampleSize int) int {
	returnValue := 0

	for i := 0; i <= sampleSize; i++ {
		x := rand.Float32()
		y := rand.Float32()

		if x*x+y*y <= 1 {
			returnValue++
		}
	}
	return returnValue
}

func DetermineSampleSize(s string) int {
	switch s {
	case "IntMax":
		return math.MaxInt32
	case "Int32":
		return math.MaxInt32
	case "Int16":
		return math.MaxInt16
	case "Int8":
		return math.MaxUint8
	}
	returnValue, err := strconv.Atoi(s)
	if err == nil {
		return returnValue
	}
	return 0
}

func MonteCarloSim() {
	fmt.Print("Sample Size?: ")
	var i string
	fmt.Scanln(&i)
	sampleSize := DetermineSampleSize(i)
	sampleTop := CalcRatio(sampleSize)
	piAprox := 4.0 * float32(sampleTop) / float32(sampleSize)
	fmt.Printf("Pi Aproximated with %v samples: %v\n", sampleSize, piAprox)
}

func LeibnizSim() {
	fmt.Print("Sample Size?: ")
	var input string
	fmt.Scanln(&input)
	sampleSize := DetermineSampleSize(input)

	sum, term := 0.0, 0.0
	for i := 0; i <= sampleSize; i++ {
		term = math.Pow(-1.0, float64(i)) / (2.0*float64(i) + 1.0)
		sum += term
	}
	pi := sum * 4.0

	fmt.Printf("Pi aproximated with %v samples: %v\n", sampleSize, pi)
}

func GaussLegendreMethod() {
	fmt.Print("Terms?: ")
	var input string
	fmt.Scanln(&input)
	sampleSize := DetermineSampleSize(input)

	a_n := 1.0
	b_n := 1.0 / math.Sqrt(2.0)
	t_n := 1.0 / 4.0
	p_n := 1.0

	for i := 0; i < sampleSize; i++ {
		a_n1 := (a_n + b_n) / 2.0
		b_n1 := math.Sqrt(a_n * b_n)
		t_n1 := t_n - p_n*((a_n-a_n1)*(a_n-a_n1))
		p_n1 := 2 * p_n
		pi := (math.Pow(float64(a_n1+b_n1), 2) / (4 * t_n1))
		fmt.Printf("Pi at index %v: %v\n", i, pi)
		a_n = a_n1
		b_n = b_n1
		t_n = t_n1
		p_n = p_n1
	}
}
