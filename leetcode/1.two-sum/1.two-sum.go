package leetcode

// 1. 两数之和

func twoSum(nums []int, target int) []int {
	// 从头开始遍历i
	for i := 0; i < len(nums); i++ {
		// 从i后面一个元素开始遍历j
		for j := i + 1; j < len(nums); j++ {
			// 如果和为目标值
			if nums[i]+nums[j] == target {
				// 那么就返回它们的数组下标
				return []int{i, j}
			}
		}
	}
	return nil
}
