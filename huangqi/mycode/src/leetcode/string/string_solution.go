package string

import "sort"

/**
 * 49. 字母异位词分组
 * 将字符串数组按照字母异位词进行分组
 * @param strs 字符串数组
 * @return 结果列表
 */
func GroupAnagrams(strs []string) [][]string {
	// 使用哈希表存储分组结果
	groups := make(map[string][]string)

	for _, str := range strs {
		// 将字符串转换为字符数组并排序
		charArray := []byte(str)
		sort.Slice(charArray, func(i, j int) bool {
			return charArray[i] < charArray[j]
		})
		key := string(charArray)
		// 根据排序后的字符串作为键，将原始字符串添加到对应的分组中
		groups[key] = append(groups[key], str)
	}

	// 将分组结果转换为二维切片，并返回
	var result [][]string
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
