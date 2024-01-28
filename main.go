package main

import (
	"fmt"
	"math"
)

// LinearSearch function
func Linearsearch(array []int, key int) int {

	var index int
	var size = len(array) // Finding array size
	for index = 0; index < size; index++ {
		if array[index] == key {
			return index //Success case (key is found)
		}
	}

	return -1 //Fail case (key is not found)
}

// BinarySearch function
func BinarySearch(array []int, key int) int {
	result := -1          // Initialize result not found
	start := 0            // Start position
	end := len(array) - 1 // End index

	for start <= end {
		mid_point := (start + end) / 2
		if array[mid_point] == key {
			result = mid_point // Success case (key is found)
			break
		} else if array[mid_point] < key {
			start = mid_point + 1
		} else if array[mid_point] > key {
			end = mid_point - 1
		}
	}

	return result //Fail case (key is not found)
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
func SentinelSearch(array []int, key int) int {
	result := -1       // Initialize result not found
	n := len(array)    // Finding array size
	last := array[n-1] // Last element of the array
	array[n-1] = key   // Element to be searched is placed at the last index
	i := 0

	for array[i] != key {
		i++
	}

	array[n-1] = last // Put the last element back

	if i < (n-1) || array[n-1] == key {
		return i // Success case (key is found)
	}

	return result // Fail case (key is not found)
}

func JumpBinarySearch(array []int, key int) int {

	var n float64 = float64(len(array)) // Finding array size
	var step float64 = math.Sqrt((n))   // Finding the block size
	var prev float64 = 0

	for array[int(math.Min(step, n)-1)] < key {
		prev = step
		step += math.Sqrt((n))
		if prev >= n {
			return -1
		}

	}
	// Linear search for x in block
	for array[int(prev)] < key {
		prev++
		// If we reached next block or end of
		// array, element is not present.
		if prev == math.Min(step, float64(key)) {
			return -1
		}
	}

	if array[int(prev)] == key {
		return int(prev) // Success case (key is found)
	}

	return -1 // Fail case (key is not found)
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
	fmt.Println(JumpBinarySearch(array2, 220))

}
