package twosum_test

import (
	"leetcode/array/twosum"
	"testing"
)

func TestTwosum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	target := 3
	result := twosum.TwoSum(nums, target)
	if len(result) == 2 {
		if result[0] == 0 && result[1] == 1 || result[1] == 0 && result[0] == 1 {
			t.Log("pass")
		} else {
			t.Fatal("not pass")
		}
	} else {
		t.Fatal("not pass")
	}
}
