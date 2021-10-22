package PracticeProblemSolving

import "fmt"

var friends = []string{"minhaj", "ahmed", "sadik", "iftakher", "hanif", "mehedi", "jobeda"}
var numbers = []int{12, 10, 11, 19, 29, 33, 34, 35, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}

func BinarySearch() {
	fmt.Println("hello from Binary Search")

	start := 0
	end := len(numbers) - 1

	fmt.Println("thinking's", friends, numbers, start, end)

}

// func BinarySearch(array []int, target int) int {
// 	startIndex := 0
// 	endIndex := len(array) - 1
// 	midIndex := len(array) / 2
// 	for startIndex <= endIndex {

// 		value := array[midIndex]

// 		if value == target {
// 			return midIndex
// 		}

// 		if value > target {
// 			endIndex = midIndex - 1
// 			midIndex = (startIndex + endIndex) / 2
// 			continue
// 		}

// 		startIndex = midIndex + 1
// 		midIndex = (startIndex + endIndex) / 2
// 	}

// 	return -1
// }
