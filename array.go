package leetcode_array

import "sort"

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
func FindSingleElement(nums []int) int {
	// 题目已经明确是非空数组
	// 针对数组只有一个元素的情况
	if len(nums) == 1 {
		return nums[0]
	}
	// 数组有多个元素
	tmp := 0
	for _, num := range nums {
		tmp ^= num
	}
	return tmp
}


/**
给定一个整数数组，判断是否存在重复元素。

如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
 */
 func containsDuplicate(nums []int) bool {
	// 首先判断数组长度，== 0 直接返回false
	if len(nums) == 0 {
		return false
	}
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i]^nums[i-1] == 0 {
			return true
		}
	}
 	return false
 }