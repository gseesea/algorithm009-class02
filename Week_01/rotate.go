package Week_01

// 翻转数组


// 方法1 暴力破解
func rotate1(nums []int, k int) {
	l := len(nums)
	for i := 0; i <k; i++ {
		// 移动最后一个数
		temp := nums[l-1]
		// 所有数向后移动一位 每次均需执行l-1次
		for j:= 0; j < l -1; j++ {
			nums[l-j-1] = nums[l-j-1-1]
		}
		nums[0] = temp
	}
}


// 方法2 利用反转
func rotate2(nums []int, k int)  {
	k %= len(nums)
	// 全部翻转
	reverse(nums, 0, len(nums) -1)
	// 翻转前k个
	reverse(nums, 0, k-1)
	// 翻转k到末尾
	reverse(nums, k, len(nums)-1)
}
func reverse(nums []int, start, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
