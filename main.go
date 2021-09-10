package main

import (
	"flag"
	"fmt"
	"math"
)

// RetValue contains the percentage to be retained for an amount in the [min,max] interval
type RetValue struct {
	min, max, percentage float64
}

// values are the current IRPF intervals in Spain
var values = []RetValue{
	{0, 12450, 0.19},
	{12450, 20200, 0.24},
	{20200, 35200, 0.30},
	{35200, 60000, 0.37},
	{60000, 300000, 0.45},
	{300000, math.Inf(0), 0.47},
}

func main() {
	amount := flag.Float64("a", 0.0, "raw annual income")
	payments := flag.Int("p", 12, "number of individual payments")
	flag.Parse()

	ret := TotalRetention(*amount, values)
	net := *amount - ret
	retPercent := ret * 100 / *amount
	month := net / float64(*payments)

	fmt.Println("Raw income:", *amount, "€")
	fmt.Println("Retention:", ret, "€")
	fmt.Println("Retention percentage:", retPercent, "%")
	fmt.Println("Net income:", net, "€")
	fmt.Println("Net income per payment (", *payments, "pays):", month, "€")
}

// PartialRetention returns the retention to be applied to an amount given a RetValue
func PartialRetention(amount float64, v RetValue) float64 {
	switch {
	case amount <= v.min:
		return 0
	case amount >= v.max:
		return (v.max - v.min) * v.percentage
	default:
		return (amount - v.min) * v.percentage
	}
}

// TotalRetention returns the sum of the retentions to be applied to an amount given a list of RetValue
func TotalRetention(amount float64, values []RetValue) float64 {
	total := 0.0
	for _, v := range values {
		total += PartialRetention(amount, v)
	}
	return total
}
