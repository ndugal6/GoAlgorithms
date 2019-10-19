package algoExpert

import (
	"fmt"
	"sort"
)

func TwoNumberSum(numbers []int, targetSum int) []int {
	result := make([]int, 0) /* init size to 2 since min numPairs =1 */
	sort.Ints(numbers)
	for floor, ceiling := 0, len(numbers)-1; floor < ceiling; floor += 1 {
		for (numbers[floor] + numbers[ceiling]) > targetSum {
			if floor+1 >= ceiling {
				break
			}
			ceiling -= 1
		}
		if numbers[floor]+numbers[ceiling] == targetSum {
			result = append(result, numbers[floor], numbers[ceiling])
			ceiling -= 1
		}
	}
	sort.Ints(result)
	fmt.Println(result)
	return result
}
