package Week_01

// 删除数组的重复项
func removeDuplicates1(nums []int) int {
	left := 0
	for _, v := range nums {
		// 不相等 加一 并赋值
		if nums[left] != v {
			nums[left+1] = v
			left++
		}
	}
	return left+1
}


// 和方法1类似，不用range 内存占用少些
func removeDuplicates2(nums []int) int {
	left, right := 1, 1
	for right < len(nums) {
		if nums[right] != nums[right-1] {
			// 不想等 赋值 左指针移动
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
