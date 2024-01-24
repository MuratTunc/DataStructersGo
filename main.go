package main

import (
	"fmt"
	"math"
)

// LinearSearch function
func Linearsearch(a []int, key int) bool {
	for _, item := range a {
		if item == key {
			return true
		}
	}
	return false
}

// BinarySearch function
func BinarySearch(a []int, x int) int {
	result := -1 // not found
	start := 0
	end := len(a) - 1

	for start <= end {
		mid := (start + end) / 2
		if a[mid] == x {
			result = mid // found
			break
		} else if a[mid] < x {
			start = mid + 1
		} else if a[mid] > x {
			end = mid - 1
		}
	}

	return result
}

// InterpolationSearch
func InterpolationSearch(array []int, key int) int {

	min, max := array[0], array[len(array)-1]

	low, high := 0, len(array)-1

	for {
		if key < min {
			return low
		}

		if key > max {
			return high + 1
		}

		// make a guess of the location
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(key-min) / float64(max-min)))
			guess = low + offset
		}

		// maybe we found it?
		if array[guess] == key {
			// scan backwards for start of value range
			for guess > 0 && array[guess-1] == key {
				guess--
			}
			return guess
		}

		// if we guessed to high, guess lower or vice versa
		if array[guess] > key {
			high = guess - 1
			max = array[high]
		} else {
			low = guess + 1
			min = array[low]
		}
	}
}

// Function to search x in the given array
func SentinelSearch(arr []int, key int) int {
	result := -1     // not found
	n := len(arr)    //array size
	last := arr[n-1] // Last element of the array
	arr[n-1] = key   // Element to be searched is placed at the last index
	i := 0

	for arr[i] != key {
		i++
	}

	arr[n-1] = last // Put the last element back

	if i < (n-1) || arr[n-1] == key {
		return i
	} else {
		fmt.Println("Element Not found")
	}

	return result
}

func JumpBinarySearch(arr []int, key int) int {

	step := int(math.Round(math.Sqrt(float64(len(arr)))))
	n := len(arr)
	// Finding the block where element is present (if it is present)
	var prev int = 0
	for arr[min(step, n)-1] < key {
		prev = step
		step += int(math.Round(math.Sqrt(float64(n))))
		if prev >= n {
			return -1
		}

	}

	// Doing a linear search for x in block
	// beginning with prev.
	for arr[prev] < key {
		prev++
		// If we reached next block or end of
		// array, element is not present.
		if prev == min(step, key) {
			return -1
		}

	}
	// If element is found
	if arr[prev] == key {
		return prev
	}

	return -1

}

func main() {

	array := []int{10, 9, 7, 2, 11, 3, 65, 89, 99}
	search_value := 11

	array2 := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	fmt.Println("array:", array)
	fmt.Println("search_value:", search_value)

	fmt.Println("*******************************")
	fmt.Println("BinarySearch Result:")
	fmt.Println(BinarySearch(array, search_value))

	fmt.Println("*******************************")
	fmt.Println("SentinelSearch Result:")
	fmt.Println(SentinelSearch(array, search_value))

	fmt.Println("*******************************")
	fmt.Println("linearsearch Result:")
	fmt.Println(Linearsearch(array, search_value))

	fmt.Println("*******************************")
	fmt.Println("interpolationSearch Result:")
	fmt.Println(InterpolationSearch(array2, 45))

	fmt.Println("*******************************")
	fmt.Println("JumpBinarySearch Result:")
	fmt.Println(JumpBinarySearch(array, search_value))

}
