package other

import (
	"sort"
	"strings"
)

// 9. 回文数 321 123
func IsPalindrome(x int) bool {
	sum := x
	var res int = 0
	for x > 0 {
		res = res*10 + x%10
		x = x / 10
	}
	return sum == res
}

// 1614. 括号的最大嵌套深度
func MaxDepth(s string) int {
	maxDepth := 0     // 最大深度
	currentDepth := 0 // 当前深度

	for _, char := range s {
		if char == '(' {
			currentDepth++               // 遇到左括号，当前深度加一
			if currentDepth > maxDepth { // 更新最大深度
				maxDepth = currentDepth
			}
		} else if char == ')' {
			currentDepth-- // 遇到右括号，当前深度减一
		}
	}

	return maxDepth
}

// 面试题 08.08. 有重复字符串的排列组合
func Permutation(S string) []string {
	// 创建一个存储结果排列组合的切片
	var resultList []string
	// 将输入字符串转换为字符数组
	chars := strings.Split(S, "")
	// 对字符数组进行排序，以确保相同字符相邻
	sort.Strings(chars)
	// 创建一个空字符串切片，用于构建当前排列组合
	var current []string
	// 创建一个布尔数组用于跟踪已使用的字符
	used := make([]bool, len(chars))

	// 调用回溯函数生成排列组合
	backtrack(chars, used, &current, &resultList)

	return resultList
}

func backtrack(chars []string, used []bool, current *[]string, resultList *[]string) {
	// 如果当前排列组合的长度与输入字符串的长度相同，则将其添加到结果切片中
	if len(*current) == len(chars) {
		*resultList = append(*resultList, strings.Join(*current, ""))
		return
	}

	// 遍历输入字符串中的字符
	for i := 0; i < len(chars); i++ {
		// 如果字符已被使用，或者是一个重复字符且前一个字符未被使用，则跳过
		if used[i] || (i > 0 && chars[i] == chars[i-1] && !used[i-1]) {
			continue
		}

		// 将当前字符添加到当前排列组合中
		*current = append(*current, chars[i])
		// 将当前字符标记为已使用
		used[i] = true

		// 递归调用回溯函数生成排列组合
		backtrack(chars, used, current, resultList)

		// 从当前排列组合中删除最后一个字符
		*current = (*current)[:len(*current)-1]
		// 将当前字符标记为未使用，以进行回溯
		used[i] = false
	}

}

// 46. 全排列
func Permute(nums []int) [][]int {
	resultList := [][]int{}              // 存储最终排列结果的切片
	deque := []int{}                     // 存储当前排列的切片
	used := make([]bool, len(nums))      // 用于记录数字是否被使用的布尔切片
	dfs(nums, used, &deque, &resultList) // 调用深度优先搜索函数生成排列
	return resultList                    // 返回排列列表
}

func dfs(nums []int, used []bool, deque *[]int, resultList *[][]int) {
	if len(nums) == len(*deque) {
		temp := make([]int, len(*deque))
		copy(temp, *deque)
		*resultList = append(*resultList, temp) // 将当前排列添加到结果列表
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] { // 如果数字在当前排列中未被使用
			*deque = append(*deque, nums[i]) // 将数字添加到当前排列的末尾
			used[i] = true                   // 将数字标记为已使用

			dfs(nums, used, deque, resultList) // 递归生成排列

			*deque = (*deque)[:len(*deque)-1] // 从当前排列中移除数字
			used[i] = false                   // 将数字标记为未使用，以便下一次迭代时使用
		}
	}
}
