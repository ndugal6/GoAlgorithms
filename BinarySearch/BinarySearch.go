package program

func BinarySearch(array []int, target int) int {
	index := len(array)/2 
	if array[index] == target {
		return index;
	} else if len(array) < 2 {
		return -1
	} else if array[index] > target {
		return BinarySearch(array[:index], target)
	} else {
		result := BinarySearch(array[index:], target)
		if result == -1 {
			return -1
		} else {
		return index + result
		}
	}
}
