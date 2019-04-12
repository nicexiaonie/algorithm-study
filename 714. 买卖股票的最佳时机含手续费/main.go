package main

import (
	"fmt"
)

func main() {
	prices := []int{345, 1, 345, 456, 456, 7687, 32, 2, 90, 4, 3, 76, 2, 8, 6, 4, 9}
	//prices := []int{1, 3, 2, 8, 4, 9}
	maxProfit(prices, 2)
}
func maxProfit(prices []int, fee int) int {
	result := make([]int, 0)
	is := true
	for key := 0; key < len(prices); key++ {
		price := prices[key]
		if is {
			for i := key + 1; i < len(prices); i++ {
				if prices[i]-2 > price {

					nextPrice := 0
					if key < len(prices)-1 {
						nextPrice = prices[key+1]
					}

					if nextPrice > price {
						result = append(result, price)
						is = false
						break
					}
				}
			}
		} else {
			maiChuKey := key
			for i := key + 1; i < len(prices); i++ {
				if prices[i]-2 > prices[maiChuKey] {
					maiChuKey = i
					continue
				}
			}
			is = true
			key = maiChuKey
			result = append(result, prices[maiChuKey])
		}
	}
	for _, value := range result {
		fmt.Printf(": %d \n", value)
	}
	return 0
}
