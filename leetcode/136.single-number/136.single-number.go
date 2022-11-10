package leetcode

// 136. 只出现一次的数字
// https://leetcode.cn/problems/single-number/

func singleNumber(nums []int) int {
	// m存储每个数字出现的次数，由于题目说“除了某个元素只出现一次以外，其余每个元素均出现两次。”所以记出现1次为真，出现2次为假
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = !m[v]
	}
	for k, v := range m {
		// 如果有找到只出现1次的数字，那么从map中找出它的序号并返回
		if v == true {
			return k
		}
	}
	// 如果没有找到，返回0
	return 0
}
