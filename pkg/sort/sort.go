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

func InsertionSortFromTo(from int, to int, data Comparable) {
	for i := from; i <= to; i++ {
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

// This is Top-Down merge sort
func MergeSortTD(data Comparable) {
	aux := make([]int, data.Len())
	var sort func(data Comparable, low, high int)
	merge := func(data Comparable, low, mid, high int) {
		li, ls := low, mid
		ri, rs := mid, high
		cursor := 0
		for ; li < ls && ri < rs; cursor++ {
			if data.Less(li, ri) {
				aux[li-low] = cursor
				li++
			} else {
				aux[ri-low] = cursor
				ri++
			}
		}
		for ; li < ls; li++ {
			aux[li-low] = cursor
			cursor++
		}
		for i := range aux[:cursor] {
			for j := aux[i]; j != i; {
				data.Swap(low+i, low+j)
				aux[i], aux[j], j = aux[j], aux[i], aux[j]
			}
		}
	}
	sort = func(data Comparable, low, high int) {
		size := high - low
		if size < 2 {
			return
		}
		mid := (low + high) / 2
		sort(data, low, mid)
		sort(data, mid, high)
		merge(data, low, mid, high)
	}
	sort(data, 0, data.Len())
}

// This is Bottom-Up merge sort
func MergeSortBU(data Comparable) {

}

func QuickPartitionSort(data Comparable) {
	var quickSort func(low, high int, data Comparable)
	var partition func(low, high int, data Comparable) int
	var medianOf3 func(a, b, c int, data Comparable) int
	const CUTOFF = 6

	medianOf3 = func(a, b, c int, data Comparable) int {
		if data.Less(a, b) {
			if data.Less(a, c) {
				if data.Less(b, c) {
					return b
				}
				return c
			}
			return a
		}
		if data.Less(b, c) {
			if data.Less(a, c) {
				return a
			}
			return c
		}
		return b
	}

	partition = func(low, high int, data Comparable) int {
		pibot, i, j := low, low+1, high
		for {
			for data.Less(i, pibot) {
				i++
				if i >= high {
					break
				}
			}

			for data.Less(pibot, j) {
				j--
				if j <= low {
					break
				}
			}

			if i >= j {
				break
			}

			data.Swap(i, j)
		}
		data.Swap(pibot, j)

		return j
	}

	quickSort = func(low, high int, data Comparable) {
		if high <= low+CUTOFF {
			InsertionSortFromTo(low, high, data)
			return
		}

		m := medianOf3(low, high, (high-low)/2, data)
		data.Swap(low, m)

		pibot := partition(low, high, data)
		quickSort(low, pibot-1, data)
		quickSort(pibot+1, high, data)
	}

	data.Shuffle()
	quickSort(0, data.Len()-1, data)
}

func QuickSort(data Comparable) {
	var quickSort func(low, high int, data Comparable)
	var medianOf3 func(a, b, c int, data Comparable) int
	const CUTOFF = 10

	medianOf3 = func(a, b, c int, data Comparable) int {
		if data.Less(a, b) {
			if data.Less(a, c) {
				if data.Less(b, c) {
					return b
				}
				return c
			}
			return a
		}
		if data.Less(b, c) {
			if data.Less(a, c) {
				return a
			}
			return c
		}
		return b
	}

	quickSort = func(low, high int, data Comparable) {
		if high <= low+CUTOFF {
			InsertionSortFromTo(low, high, data)
			return
		}

		m := medianOf3(low, high, (high-low)/2, data)
		data.Swap(low, m)

		i, lt, gt := low+1, low, high
		for i <= gt {
			if data.Less(i, lt) {
				data.Swap(lt, i)
				lt++
				i++
			} else if data.Less(lt, i) {
				data.Swap(i, gt)
				gt--
			} else {
				i++
			}
		}

		quickSort(low, lt-1, data)
		quickSort(gt+1, high, data)
	}

	data.Shuffle()
	quickSort(0, data.Len()-1, data)
}
