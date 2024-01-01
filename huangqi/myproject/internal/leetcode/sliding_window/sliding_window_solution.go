package sliding_window

//滑动窗口

// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	var last [128]int
	for i := 0; i < 128; i++ {
		last[i] = -1
	}

	res, start := 0, 0

	for i := 0; i < len(s); i++ {
		index := s[i]
		//更新窗口的起始位置 last[i]初始值为 = -1， last[i]+1 > start 则重复了，start = last[i]+1 比如 a,b,c,d ,start  =0,若第二个a 出现 last[a] = 0 , start = 0+1
		start = max(start, last[index]+1)
		//更新最长不重复子串的长度
		res = max(res, i-start+1)
		last[index] = i //字符作为下表，值存索引i
	}
	return res
}
