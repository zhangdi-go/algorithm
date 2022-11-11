package leetcode

import "strings"

// 1704. 判断字符串的两半是否相似
// https://leetcode.cn/problems/determine-if-string-halves-are-alike/

func halvesAreAlike(s string) bool {
	// cnt是计算元音字母数量的变量
	cnt := 0
	// 从索引0切到索引(len(s)/2)-1，共len(s)/2个元素
	for _, c1 := range s[:len(s)/2] {
		// 如果c1在字符串中，计数器cnt加1
		if strings.ContainsRune("aeiouAEIOU", c1) {
			cnt++
		}
	}
	// 从索引len(s)/2切到最后，共len(s)/2个元素
	for _, c2 := range s[len(s)/2:] {
		// 如果c2在字符串中，计数器cnt减1
		if strings.ContainsRune("aeiouAEIOU", c2) {
			cnt--
		}
	}
	// 如果最终cnt变成了0，说明两边元音字母数量一样多，返回真
	return cnt == 0
}
