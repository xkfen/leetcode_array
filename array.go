package leetcode_array

import (
	"sort"
)

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
	for i := 1; i < len(nums); i++ {
		// 如果临时变量tmp与当前循环的元素相等，continue
		if tmp == nums[i] {
			continue
		} else {
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

/**
加一
给定一个非负整数组成的非空数组，在该数的基础上加一，返回一个新的数组。

最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头

*/
/**
解法的关键在于弄明白什么情况下会产生进位
​
想让个位+1进位，那么个位必须为9
​
想让十位+1进位，那么十位必须为9，想要产生进位carry，那么必须由个位进位而来。想让个位进位，个位必须为9.
​
想让百位+1进位，那么百位必须为9，想要产生进位carry，那么必须由十位进位而来，想让十位进位，那么十位必须为9，想要产生进位，个位必须为9。
​
根据以上可以推论得出两种情况：
1.最高位进位
2.最高位不进位
​
最高位进位
//若最高位进位，那么比他低的位数字都为9，且加1后都为0，需要初始化一个长度为(lenght+1)的新数组，0位置为1代表进位
​
最高位不进位
//若最高位不进位，那么不需要产生新数组，后续数字由更低位计算而来
 */
func plusOne(digits []int) ([]int) {
	carry := 1
	// 倒着循环
	for i := len(digits) - 1; i >= 0; i-- {
		// 最后一个元素+1
		tmp := digits[i] + carry
		// 最后一个元素 +1 以后整除10
		carry = tmp / 10
		digits[i] = tmp % 10
	}
	if carry != 0 {
		array := make([]int, len(digits)+1)
		array[0] = 1
		return array
	}
	return digits
}

func plusOne2(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i -- {
		// 非9加1
		if digits[i] < 9 {
			// 如果是前一位需要进位的话那么前一位置为0,就在后一位(也就是此时的当前位+1)
			digits[i]++
			// 直接把改变之后的数组返回(因为当前位不是9就不需要进位了,前面该进的位都已经进了所以直接把数组返回就行了)
			return digits
		}
		// 缝9进0
		digits[i] = 0
	}
	// 全部为9,则需要数组扩充1位
	/**
	 * 如果位数全部都是9的话,那么遍历原来=数组上的所有的数,都是让数组上的所有的数变成0
	 * 也就是说明上面的if代码里面的内柔都是没有走的,就需要进行数组扩容
	 */
	 // 新建的数组，没有初始化默认为0
	result := make([]int, n+1)
	result[0] = 1
	return result
}
