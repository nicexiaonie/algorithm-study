package main

import "fmt"

func main() {

	arr := []int{11, 13, 4, 16, 2, 5, 12, 47, 9, 4}

	for k, v := range arr {


		fmt.Print(k)
		fmt.Print(v)
	}

}
