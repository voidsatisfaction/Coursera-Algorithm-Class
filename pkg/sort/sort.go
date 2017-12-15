package sort

func SelectionSort(data Comparable) {
	N := data.Len()
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}

func InsertionSort(data Comparable) {
	N := data.Len()
	for i := 0; i < N; i++ {
		for j := i; j > 0; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}

func ShellSort(data Comparable) {
	N := data.Len()
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && data.Less(j, j-h); j -= h {
				data.Swap(j, j-h)
			}
		}
		h /= 3
	}
}
