package leetcode_array

import (
	"testing"
	"sort"
	"fmt"
	"github.com/stretchr/testify/assert"
)
//go:generate
// 测试Int数组·排序
func TestRortIntArray(t *testing.T){
	nums := []int{0,0,4,1,3,2,2,3,3,4}
	sort.Ints(nums)
	fmt.Println(nums)
}

func TestRemoveDuplicates(t *testing.T){
	nums := []int{0, 0, 1, 2 ,2 ,3 ,3, 3 ,4 ,4}
	length:= RemoveDuplicates(nums)
	fmt.Println(length)
}

// 按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或
func TestMove(t *testing.T){
	num1 := 1
	num2 := 1
	fmt.Println(num1 & num2)
}

func TestFindSingleElement(t *testing.T) {
	nums := []int{4,1,2,1,2}
	tmp := FindSingleElement(nums)
	fmt.Println(tmp)
}

//
func TestContainsDuplicate(t *testing.T){
	nums1 := []int{1,1,1,3,3,4,3,2,4,2}
	flag1 := containsDuplicate(nums1)
	assert.Equal(t, true, flag1)
	fmt.Println(flag1)
	nums2 := []int{1,2,3,1}
	flag2 := containsDuplicate(nums2)
	assert.Equal(t, true, flag2)
	fmt.Println(flag2)

	nums3 := []int{1,2,3,4}
	flag3 := containsDuplicate(nums3)
	assert.Equal(t, false, flag3)
	fmt.Println(flag3)
}