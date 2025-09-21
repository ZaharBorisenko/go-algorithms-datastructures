package main

import (
	"fmt"
	"github.com/ZaharBorisenko/go-algorithms-datastructures/pkg/algorithms/sorting/bubble"
)

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", arr)

	sorted := bubble.Sort(arr)
	fmt.Println("Sorted array:", sorted)
}
