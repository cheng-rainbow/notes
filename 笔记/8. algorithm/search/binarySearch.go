package search

func binarySearch1(arr []int, l, r, t int) int {
	i, j := l, r
	for i < j {
		mid := (i + j) >> 1
		if arr[mid] >= t {
			j = mid
		} else if arr[mid] < t {
			i = mid + 1
		}
	}

	if arr[i] == t {
		return i
	}
	return -1
}

func binarySearch2(arr []int, l, r, t int) int {
	i, j := l, r
	for i < j {
		mid := (i + j + 1) >> 1
		if arr[mid] <= t {
			i = mid
		} else if arr[mid] > t {
			j = mid - 1
		}
	}

	if arr[i] == t {
		return i
	}
	return -1
}
