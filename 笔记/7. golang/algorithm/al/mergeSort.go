package al

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) >> 1
	i, j := l, mid+1
	tmp := make([]int, r-l+1)

	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)

	idx := 0
	for i <= mid && j <= r {
		if arr[i] <= arr[j] {
			tmp[idx] = arr[i]
			i++
		} else {
			tmp[idx] = arr[j]
			j++
		}
		idx++
	}
	for i <= mid {
		tmp[idx] = arr[i]
		i, idx = i+1, idx+1
	}

	for j <= r {
		tmp[idx] = arr[j]
		j, idx = j+1, idx+1
	}

	for i := 0; i < idx; i++ {
		arr[l+i] = tmp[i]
	}
}
