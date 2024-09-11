package main

import (
	"fmt"
	"sort"
)

func closestNumbers(arr []int32) []int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	minDiff := arr[1] - arr[0]
	result := []int32{arr[0], arr[1]}
	for i := int32(2); i < int32(len(arr)); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
			result = []int32{arr[i-1], arr[i]}
		} else if diff == minDiff {
			result = append(result, arr[i-1], arr[i])
		}
	}

	return result

}
func main() {
	arr := []int32{5, 4, 3, 2}
	fmt.Println(closestNumbers(arr))
}
