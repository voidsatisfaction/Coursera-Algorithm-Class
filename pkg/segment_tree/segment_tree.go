package segmentTree

// This segment tree saves maximum value of certain segments
func max(nums ...int) int {
	maxNum := nums[0]
	for _, n := range nums {
		if n > maxNum {
			maxNum = n
		}
	}
	return maxNum
}

type segmentTree []int

func New(arr []int) segmentTree {
	segTree := make(segmentTree, 4*len(arr))

	st, ed := 0, len(arr)-1
	getMax(segTree, arr, st, ed, 1)
	return segTree
}

func getMax(segTree segmentTree, arr []int, st, ed, node int) int {
	if st == ed {
		segTree[node] = arr[st]
		return arr[st]
	}

	mid := (st + ed) / 2
	// You can customize this part according to function of your segment tree
	segTree[node] = max(
		getMax(segTree, arr, st, mid, 2*node),
		getMax(segTree, arr, mid+1, ed, 2*node+1),
	)
	return segTree[node]
}

func (segTree segmentTree) Get(st, ed, tst, ted, node int) int {
	if tst <= st && ted >= ed {
		return segTree[node]
	}

	mid := (st + ed) / 2
	if ted <= mid {
		return segTree.Get(st, mid, tst, ted, 2*node)
	} else if tst >= mid+1 {
		return segTree.Get(mid+1, ed, tst, ted, 2*node+1)
	} else {
		return max(segTree.Get(st, mid, tst, ted, 2*node), segTree.Get(mid+1, ed, tst, ted, 2*node+1))
	}
}

func (segTree segmentTree) Update(st, ed, i, newVal, node int) int {
	if st == i && ed == i {
		segTree[node] = newVal
		return newVal
	}

	mid := (st + ed) / 2
	if i <= mid {
		segTree[node] = max(segTree.Update(st, mid, i, newVal, 2*node), segTree[2*node+1])
	} else if i >= mid+1 {
		segTree[node] = max(segTree[2*node], segTree.Update(mid+1, ed, i, newVal, 2*node+1))
	}
	return segTree[node]
}
