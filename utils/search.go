package utils

// BinarySearch 二分查找
func BinarySearch(array []int, needle int) int {
	var haystack []int
	for _, item := range array {
		haystack = append(haystack, item)
	}
	return handleBinarySearch(haystack, needle, 0, len(haystack)-1)
}

func handleBinarySearch(haystack []int, needle, start, end int) int {
	m := end - start
	mid := start + m/2
	//如果起始与结束下标相同，并且该下标的值与查找的值n相同，则直接返回该下标
	if m == 0 && haystack[start] == needle {
		return start
	} else if m < 0 {
		//特殊情况，未找到n的下标，直接返回-1
		return -1
	} else {
		//中间值与查找的值n相同，返回中间值的下标
		if haystack[mid] == needle {
			return mid
		} else if haystack[mid] < needle {
			// n大于中间值，说明要查找的n值在右半部分，递归查找右半部分
			return handleBinarySearch(haystack, needle, mid+1, end)
		} else {
			// n小于中间值，说明要查找的n值在左半部分，递归查找左半部分
			return handleBinarySearch(haystack, needle, start, mid-1)
		}
	}
}
