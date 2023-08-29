package sort

// MergeSort 使用归并排序算法对整数切片进行排序
func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

// mergeSort 是归并排序的递归函数，对指定范围内的元素进行排序
func mergeSort(arr []int, left, right int) {
	// 如果左边界大于等于右边界，表示只有一个元素或无元素，直接返回
	if left >= right {
		return
	}

	// 计算中间位置
	var mid int = left + (right-left)/2

	// 递归排序左半部分
	mergeSort(arr, left, mid)
	// 递归排序右半部分
	mergeSort(arr, mid+1, right)

	// 创建临时数组，用于合并两个有序子数组
	var tempArr []int = make([]int, len(arr))

	// 合并两个有序子数组
	merge(arr, tempArr, left, mid, right)
}

// merge 合并两个有序子数组
func merge(arr []int, tempArr []int, left, mid, right int) {
	// 左子数组的起始位置
	l_pos := left
	// 右子数组的起始位置
	r_pos := mid + 1
	// 合并后数组的当前位置
	pos := left

	// 比较左右子数组的元素，将较小的元素放入临时数组
	for l_pos <= mid && r_pos <= right {
		if arr[l_pos] > arr[r_pos] {
			tempArr[pos] = arr[r_pos]
			r_pos++
		} else {
			tempArr[pos] = arr[l_pos]
			l_pos++
		}
		pos++
	}

	// 将剩余的左子数组元素复制到临时数组
	for l_pos <= mid {
		tempArr[pos] = arr[l_pos]
		l_pos++
		pos++
	}

	// 将剩余的右子数组元素复制到临时数组
	for r_pos <= right {
		tempArr[pos] = arr[r_pos]
		r_pos++
		pos++
	}

	// 将临时数组中的元素复制回原始数组
	for i := left; i <= right; i++ {
		arr[i] = tempArr[i]
	}
}
