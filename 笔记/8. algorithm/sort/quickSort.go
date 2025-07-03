package sort

func quickSort(arr []int, l, r int) {
	// 一开始就相同时，没必要进行下面的操作
	if l >= r {
		return
	}

	i, j, base := l, r, arr[(l+r)>>1]
	for i < j {
		for arr[i] < base {
			i++
		}
		for arr[j] > base {
			j--
		}
		// 相同时也要进来进行 i ++, j -- 操作，否则下面 l - j, i - r 的区间会出错，导致无限递归
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i, j = i+1, j-1
		}
	}
	quickSort(arr, l, j)
	quickSort(arr, i, r)
}
