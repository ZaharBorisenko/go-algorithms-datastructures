package main

import (
	"fmt"
	selection "github.com/ZaharBorisenko/go-algorithms-datastructures/pkg/algorithms/sorting/Selection"
)

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", arr)

	sorted := selection.Sort(arr)
	fmt.Println("Sorted array:", sorted)
}
