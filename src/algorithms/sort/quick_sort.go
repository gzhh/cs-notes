package sort

func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	target := nums[left]
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= target {
			j--
		}
		for i < j && nums[i] <= target {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	QuickSort(nums, left, i-1)
	QuickSort(nums, i+1, right)
}
