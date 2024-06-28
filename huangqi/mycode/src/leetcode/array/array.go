package array

import (
	"math"
	"sort"
)

// 674. 最长连续递增序列
func FindLengthOfLCIS(nums []int) int {
	maxLen := 1
	currentLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentLen++
			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			currentLen = 1
		}
	}
	return maxLen
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算成功举办活动需要多少名主持人
 * @param n int整型 有n个活动
 * @param startEnd int整型二维数组 startEnd[i][0]用于表示第i个活动的开始时间，startEnd[i][1]表示第i个活动的结束时间
 * @return int整型
 */
func MinmumNumberOfHost(n int, startEnd [][]int) int {
	var start, end []int

	for _, v := range startEnd {
		start = append(start, v[0])
		end = append(end, v[1])
	}

	sort.Ints(start)
	sort.Ints(end)

	res := 1
	j := 0
	for i := 1; i < n; i++ {
		if start[i] >= end[j] {
			j++
		} else {
			res++
		}
	}
	return res
}

// 26. 删除有序数组中的重复项
func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}

	return index + 1
}

// 80. 删除有序数组中的重复项 II
func RemoveDuplicatesII(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	index := 1
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[index-1] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}

//189. 轮转数组
/**
 *
 * @param nums 整数数组 nums
 * @param k    k 是非负数
 */
func Rotate(nums []int, k int) {
	len := len(nums)
	k = k % len
	reverse(nums, 0, len-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len-1)
}

/**
 * 反转数组中指定范围的元素
 *
 * @param nums  数组
 * @param start 起始索引
 * @param end   结束索引
 */
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// 121. 买卖股票的最佳时机
func MaxProfit(prices []int) int {
	maxProfit := math.MinInt32
	minPrice := math.MaxInt32
	for _, price := range prices {
		minPrice = min(price, minPrice)
		maxProfit = max(price-minPrice, maxProfit)
	}
	return maxProfit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 122. 买卖股票的最佳时机 II
func MaxProfitII(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
