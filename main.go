package main

import "fmt"


// LinearSearch function
func linearsearch(a []int, key int) bool {
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

func bubbleSort(arr []int) {
	n := len(arr)
	swapped := true
	m := 0

	for swapped {
		swapped = false
		
		fmt.Println("iteration number: " ,m)
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				// Swap elements if they are in the wrong order
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		n-- // The largest element is now in its correct position, 
            // so reduce the length of the unsorted part of the array
		

		m++
	}
}

//interpolationSearch
func interpolationSearch(array []int, key int) int {

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
    result := -1 // not found
	n:=len(arr) //array size
    last := arr[n - 1] // Last element of the array
    arr[n - 1] = key // Element to be searched is placed at the last index
    i := 0
 
    for arr[i] != key {
		i++
	}

    arr[n - 1] = last // Put the last element back
 
    if i < (n - 1) || arr[n - 1] == key {
		return i
	} else {
		fmt.Println("Element Not found")
	}
	
	return result
}

func main() {

    array := []int{10,9,7,2,11,3,65,89,99}
    search_value:=11

	array2 := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(interpolationSearch(array2,45))

	fmt.Println("array:",array)
	fmt.Println("search_value:",search_value)
	fmt.Println("*******************************")
	fmt.Println("BinarySearch Result:")
    fmt.Println(BinarySearch(array,search_value))

	fmt.Println("SentinelSearch Result:")
	fmt.Println(SentinelSearch(array,search_value))
    
	fmt.Println("bubbleSort Result:")
    bubbleSort(array)
	fmt.Println("The sorted array is:", array)

    fmt.Println(linearsearch(array,99))


    
}