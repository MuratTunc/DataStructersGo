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


func main() {

    array := []int{10,9,7,2,11,3,65,89,99}
    search_value:=11
    fmt.Println(BinarySearch(array,search_value))
    
}