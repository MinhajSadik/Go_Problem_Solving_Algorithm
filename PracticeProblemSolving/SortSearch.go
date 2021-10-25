package PracticeProblemSolving

import (
	"fmt"
	"sort"
)

func SortSearch() {

	numbers := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] >= x
	})

	if i < len(numbers) && numbers[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, numbers)
	} else {
		fmt.Printf("%d not found in %v\n", x, numbers)
	}

}

func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}
