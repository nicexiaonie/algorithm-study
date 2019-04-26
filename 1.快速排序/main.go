package main

import (
	"fmt"
	"os"
	"time"
	"bufio"
	"strconv"
	"io"
)

func main() {

	//arr := getFileContent("1.快速排序/demo.txt")

	//arr := []int{21, 4, 5, 12, 47, 9, 4, 13, 10, 11}
	//arr := []int{11, 4, 5, 12, 47, 9, 4, 13, 10, 21}
	arr := []int{11, 13, 4, 16, 2, 5, 12, 47, 9, 4}
	left := 0
	right := len(arr) - 1
	fmt.Printf("len  %d\n", len(arr))
	start := time.Now().UnixNano() / 1000000
	fmt.Printf("start 	sort %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("start 	sort %d\n", start)
	sort(arr, left, right)
	end := time.Now().UnixNano() / 1000000
	fmt.Printf("end		sort %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("end		sort %d\n", end)
	fmt.Printf("time:	%d\n", end-start)
	//fmt.Printf("%s", arr)

}

func sort(arr []int, left int, right int) {
	pivot := 0
	//fmt.Printf("%d, %d, %d \n", left, right, pivot)
	if left < right {
		fmt.Println(arr)
		pivot = partition(arr, left, right)
		fmt.Println(arr)
		fmt.Println(pivot)
		os.Exit(9)
		sort(arr, left, pivot-1)
		fmt.Println(arr)
		sort(arr, pivot+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	key := arr[left]
	for left < right {
		for left < right && arr[right] >= key {
			right --
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= key {
			left ++
		}
		arr[right] = arr[left]
	}
	arr[left] = key
	return left
}

func getFileContent(p string) []int {
	file, err := os.OpenFile(p, os.O_RDWR, 0766)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	br := bufio.NewReader(file)
	r := make([]int, 0)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		v, _ := strconv.Atoi(string(a))
		r = append(r, v)

	}
	//fmt.Println(r)
	//fmt.Println(len(r))
	//os.Exit(9)
	return r
}
