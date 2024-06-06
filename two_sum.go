package main

import "fmt"

func main() {
	var output_1 []int = twoSum([]int{2, 7, 11, 15}, 9)

	if output_1[0] == 0 && output_1[1] == 1 {
		fmt.Println("Test 1 passed")
	}
	var output_2 []int = twoSum([]int{3, 2, 4}, 6)

	if output_2[0] == 1 && output_2[1] == 2 {
		fmt.Println("Test 2 passed")
	}
	var output_3 []int = twoSum([]int{3, 3}, 6)

	if output_3[0] == 0 && output_3[1] == 1 {
		fmt.Println("Test 3 passed")
	}
}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

	}

	return []int{0, 1}
}
