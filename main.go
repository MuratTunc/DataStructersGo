package main

import "fmt"

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
		
		fmt.Println("Sorted array:", arr)
		m++
	}
}

func main() {

    array := []int{10,9,7,2,11,3,65,89,99}
    search_value:=11
    fmt.Println(BinarySearch(array,search_value))

    bubbleSort(array)
	fmt.Println("The sorted array is:", array)
    
}