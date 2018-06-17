package leetcode_array

import (
	"testing"
	"sort"
	"fmt"
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