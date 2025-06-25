package sort

func BubbleSort(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	copy(res, nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}

	return res
}
