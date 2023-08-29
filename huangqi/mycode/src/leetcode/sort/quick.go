package sort

// QuickSort 是对给定切片进行快速排序的入口函数
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

// quickSort 是快速排序的递归函数，用于对给定切片进行排序
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := partition(arr, left, right)
	quickSort(arr, left, pivotIndex-1)
	quickSort(arr, pivotIndex+1, right)
}

// partition 用于将切片按照基准元素进行分区，并返回基准元素的最终位置
func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[right], arr[i] = arr[i], pivot
	return i
}

func ThreeWayQuickSort(arr []int) {
	threeWayQuickSort(arr, 0, len(arr)-1)
}

func threeWayQuickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	lt, gt := low, high
	pivot := arr[low]
	i := low + 1

	for i <= gt {
		if arr[i] < pivot {
			arr[i], arr[lt] = arr[lt], arr[i]
			lt++
			i++
		} else if arr[i] > pivot {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		} else {
			i++
		}
	}

	threeWayQuickSort(arr, low, lt-1)
	threeWayQuickSort(arr, gt+1, high)
}
