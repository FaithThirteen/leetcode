package main

/**
 *
 * @param rotateArray int整型一维数组
 * @return int整型
 */
func minNumberInRotateArray(rotateArray []int) int {
	// write code here

	low, high := 0, len(rotateArray)-1
	for low < high {
		mid := (low + high) / 2
		if rotateArray[mid] > rotateArray[high]{ // 中间大于结尾，排除左区间，并且mid肯定不是最小的可以前移
			low = mid+1
		}else if rotateArray[mid] < rotateArray[high]{ // 小于，不能确定mid是否是最小，可以将右区间排除
			high = mid
		}else{ // 相等，我们确定不了最小值在哪个区间
			high--
		}
	}
	return rotateArray[low]
}
