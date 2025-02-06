package main

import "fmt"

func sliceSum() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	const numGorutines = 3
	ch := make(chan int, numGorutines)
	partSize := len(arr) / numGorutines
	for i := 0; i < numGorutines; i++ {
		start := i * partSize
		end := start + partSize
		go sumPart(arr[start:end], ch)
	}
	totalSum := 0
	for i := 0; i < numGorutines; i++ {
		totalSum += <-ch
	}
	fmt.Println("Total sum:", totalSum)
}

func sumPart(part []int, ch chan int) {
	sum := 0
	for _, num := range part {
		sum += num
	}
	ch <- sum
}
