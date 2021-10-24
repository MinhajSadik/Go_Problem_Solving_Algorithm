package PracticeProblemSolving

import "fmt"

func LinearSearch() {
	fmt.Println("Hello From Linear Search")
	var friends = []string{"minhaj", "ahmed", "sadik", "iftakher", "hanif", "mehedi", "jobeda"}

	for i := 0; i < len(friends); i++ {
		names := friends[i]
		if names == "minhaj" || names == "iftakher" {
			fmt.Println("people-names", names, "positions", i)
		}

	}

	var numbers = []int{12, 10, 11, 19, 29, 33, 34, 35, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}

	for i := 0; i < len(numbers); i++ {
		number := numbers[i]
		if number == 34 || number == 41 {
			fmt.Println("number", number, "position", i)
		}

	}

	fmt.Println("slice", numbers, friends)

}
