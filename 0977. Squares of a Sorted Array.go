package leetcode

/**
 * Дан целочисленный массив, отсортированный в неубывающем порядке, вернуть массив квадратов каждого числа,
 * отсортированный в неубывающем порядке
 * Пример 1:
 * Входной массив: nums = [-4,-1,0,3,10]
 * Результат: [0,1,9,16,100]
 * Пояснения: После возведения в квдрат, получился массив [16,1,0,9,100], а после сортировки стал [0,1,9,16,100].
 * Пример 2:
 * Входной массив: nums = [-7,-3,2,3,11]
 * Результат: [4,9,9,49,121]
 */

func sortedSquares(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	left := 0
	right := l - 1
	i := right
	for left <= right {
		if nums[right]*nums[right] > nums[left]*nums[left] {
			res[i] = nums[right] * nums[right]
			right--
		} else {
			res[i] = nums[left] * nums[left]
			left++
		}
		i--
	}

	return res
}
