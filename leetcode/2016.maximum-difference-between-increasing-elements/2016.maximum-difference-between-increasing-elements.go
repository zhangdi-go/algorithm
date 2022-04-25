package leetcode

// 2016. 增量元素之间的最大差值

// https://leetcode.cn/problems/maximum-difference-between-increasing-elements/

func maximumDifference(nums []int) int {
	// min初始化为nums数组第一个元素
	min := nums[0]
	// max初始化为-1(题目要求)
	max := -1

	// 遍历
	for i := 1; i < len(nums); i++ {
		// 如果第i个元素比min大
		if nums[i] > min {
			// 且差值比之前的max大
			if nums[i]-min > max {
				// 更新max
				max = nums[i] - min
			}
		} else {
			// 更新min
			min = nums[i]
		}
	}

	// 循环结束，返回max
	return max
}
