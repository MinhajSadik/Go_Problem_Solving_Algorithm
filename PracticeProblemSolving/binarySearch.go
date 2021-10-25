package PracticeProblemSolving

// var friends = []string{"minhaj", "ahmed", "sadik", "iftakher", "hanif", "mehedi", "jobeda"}
// var numbers = []int{12, 10, 11, 19, 29, 33, 34, 35, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}

// func BinarySearch(numbers []int, target int) int {
// 	fmt.Println("hello from Binary Search")
// 	start := 0
// 	end := len(numbers) - 1

// 	for start <= end {
// 		midIndex := len(numbers) / 2

// 		if midIndex == target {
// 			return midIndex
// 		}
// 		if midIndex < target {
// 			start = midIndex + 1
// 		}
// 		if midIndex > target {
// 			start = midIndex - 1
// 		}
// 	}

// 	return -1
// }

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

// func BinarySearch(a []int, x int) int {
//     r := -1 // not found
//     start := 0
//     end := len(a) - 1
//     for start <= end {
//         mid := (start + end) / 2;
//         if a[mid] == x {
//             r = mid // found
//             break
//         } else if a[mid] < x {
//             start = mid + 1
//         } else if a[mid] > x {
//             end = mid - 1
//         }
//     }
//     return r
// }
