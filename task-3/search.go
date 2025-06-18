package task3

import "fmt"

var (
	ErrEmptyArr = fmt.Errorf("empty array")
	ErrCantFindTarget = fmt.Errorf("target not found")
)

func BinSearch(arr []int, target int) (int, error) {
	if len(arr) == 0 {
		return -1, ErrEmptyArr 
	}

	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid, nil
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, ErrCantFindTarget
}