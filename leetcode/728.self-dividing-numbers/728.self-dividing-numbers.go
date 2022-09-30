package leetcode

// 728. 自除数

// https://leetcode.cn/problems/self-dividing-numbers/

// 遍历整数的每一位的方法：每次将当前整数对10取模即可得到当前整数的最后一位，然后将当前整数除以10，重复该操作，直到当前整数变成0时即遍历了原始整数的每一位
func isSelfDividing(num int) bool {
	// 传入的原始整数num拷贝给当前整数x，每次循环结束后x除以10，x就是个从num中取出每一个d的工具
	for x := num; x > 0; x /= 10 {
		// d是当前整数x对10取模得到的最后一位数
		d := x % 10
		// 判断num是否可以被它包含的每一位数d整除
		// 若自除数包含0 || 若num不能被它包含的每一位数d整除
		if d == 0 || num%d != 0 {
			// 就不是自除数
			return false
		}
	}
	// 否则就是自除数
	return true
}

// 遍历范围[left, right]内的所有整数，判断是否为自除数。如果是就添加到答案的数组里
func selfDividingNumbers(left int, right int) (answer []int) {
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			answer = append(answer, i)
		}
	}
	return answer
}
