/**
 */
package leetcode

func peakIndexInMountainArray(arr []int) int {
	var mid int = 0
	left := 0
	right := len(arr) - 1
	for left < right {
		mid = left + (right-left)/2
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
