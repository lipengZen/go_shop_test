package leetcode

import (
	"fmt"
	"testing"
)

func moveZeroes(nums []int) {

	if nums == nil || len(nums) <= 1 {
		return
	}

	var j int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && i > j {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			j++
		}
	}

	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}

	return

}

func TestMoveZeroes(t *testing.T) {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)

	fmt.Println(arr)

}
