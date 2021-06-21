package main

import (
	"flag"
	"fmt"
	"math"
)

type RetValues struct {
	min, max, percentage float64
}

var values = []RetValues{
	{0, 12450, 0.19},
	{12450, 20200, 0.24},
	{20200, 35200, 0.30},
	{35200, 60000, 0.37},
	{60000, 300000, 0.45},
	{300000, math.Inf(0), 0.47},
}

func main() {
	amount := flag.Float64("a", 0.0, "raw annual income")
	flag.Parse()

	ret := TotalRetention(*amount, values)
	net := *amount - ret
	month := net / 12

	fmt.Println("Raw income:", *amount, "€")
	fmt.Println("Retention:", ret, "€")
	fmt.Println("Net income:", net, "€")
	fmt.Println("Net monthly income:", month, "€")
}

func PartialRetention(amount float64, v RetValues) float64 {
	switch {
	case amount <= v.min:
		return 0
	case amount >= v.max:
		return (v.max - v.min) * v.percentage
	default:
		return (amount - v.min) * v.percentage
	}
}

func TotalRetention(amount float64, values []RetValues) float64 {
	total := 0.0
	for _, v := range values {
		total += PartialRetention(amount, v)
	}
	return total
}
