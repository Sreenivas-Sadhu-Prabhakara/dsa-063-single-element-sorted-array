package main

import "fmt"

// SingleElement returns the one value that appears once in a sorted array
// where every other value appears exactly twice. Binary search on pair
// alignment: before the single, pairs start at even indexes; after, at odd.
func SingleElement(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if mid%2 == 1 {
			mid-- // inspect pairs from their even start
		}
		if nums[mid] == nums[mid+1] {
			lo = mid + 2 // pair intact: the single is to the right
		} else {
			hi = mid // alignment broken at or before mid
		}
	}
	return nums[lo]
}

func main() {
	fmt.Println(SingleElement([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
}
