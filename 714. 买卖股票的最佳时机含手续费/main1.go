package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []float64{345, 1, 345, 456, 456, 7687, 32, 2, 90, 4, 3, 76, 2, 8, 6, 4, 9}
	result := maxProfit1(prices, 2)
	fmt.Printf("利润: %f \n", result)
}
func maxProfit1(prices []float64, fee float64) float64 {
	n := len(prices)
	if n < 1 {
		return 0
	}
	buy := -prices[0]
	cash := 0.0
	for i := 1; i < n; i++ {
		cash = math.Max(cash, float64(buy+prices[i]-fee))
		buy = math.Max(buy, cash-prices[i])
	}
	return cash
}
