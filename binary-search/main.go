package main

import "fmt"

// Given a sorted array of distinct integers and a target value, return the index if the target is found.*
// If not, return the index where it would be if it were inserted in order.*

// Example 1:*
// Input: nums = [1,3,5,6], target = 5*
// Output: 2*

// Example 2:*
// Input: nums = [1,3,5,6], target = 2*
// Output: 1*

// Example 3:*
// Input: nums = [1,3,5,6], target = 7*
// Output: 4*

// Example 4:*
// Input: nums = [1,3,5,6], target = 0*
// Output: 0*

func main() {
	nums := []int{1, 3, 5, 6}

	out1 := binarySearch(nums, 7)
	p := fmt.Sprintf("not found %d", out1)
	fmt.Println(p)

	out := searchReturnIndex(nums, 7)
	fmt.Println(out)
}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2 // To prevent potential overflow

		if arr[mid] == target {
			return mid // Target found at index 'mid'
		} else if arr[mid] < target {
			low = mid + 1 // Search in upper half
		} else {
			high = mid - 1 // Search in lower half
		}
	}

	return -1 // Target not found
}

// searchReturnIndex not only does binary search for a value but if the value if not
// found will returns the index in where the `target` should be inserted to keep the
// array sorted.
func searchReturnIndex(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2 // To prevent potential overflow

		if arr[mid] == target {
			return mid // Target found at index 'mid'
		} else if arr[mid] < target {
			low = mid + 1 // Search in upper half
		} else {
			high = mid - 1 // Search in lower half
		}
	}

	// normally binary search would only return -1 if not found
	// but in this case we want to return the position that the item would be
	// when not found.
	return high + 1 // Target not found
}
