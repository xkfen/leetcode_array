package leetcode_array

/**
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
 */
func RemoveDuplicates(nums []int) (int) {
	// 数组为nil或者长度为0 的情况
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// 数组只有一个元素的情况
	if len(nums) == 1 {
		return 1
	}
	// 数组有多个元素
	// 临时变量，数组第一个元素
	tmp := nums[0]
	// 用来统计新数组的长度
	length := 0
	for i := 1 ; i < len(nums); i++ {
		// 如果临时变量tmp与当前循环的元素相等，continue
		if tmp == nums[i] {
			continue
		}else {
			//如果不相等，重新赋值当前元素给tmp
			tmp = nums[i]
			nums[length] = nums[i]
			length++
		}
	}
	return length + 1
}

/**
只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 */