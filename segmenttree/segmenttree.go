package segmenttree

import "fmt"

const maxLEN = 1000

type Segtree struct {
	arrSize int
	array   []int
	tree    [maxLEN]int
}

func (stree *Segtree) Build(arr []int, size int) {
	stree.array = arr
	stree.arrSize = size
	stree.buildTree(0, 0, size-1)
}

func (stree *Segtree) PrintTree(size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("tree[%d] = %d\n", i, stree.tree[i])
	}
}

func (stree *Segtree) buildTree(node int, start int, end int) {
	if start == end {
		stree.tree[node] = stree.array[start]
	} else {
		mid := (start + end) / 2
		leftNode := 2*node + 1
		rightNode := 2*node + 2
		stree.buildTree(leftNode, start, mid)
		stree.buildTree(rightNode, mid+1, end)
		stree.tree[node] = stree.tree[leftNode] + stree.tree[rightNode]
	}
}

func (stree *Segtree) Query(left int, right int) int {
	return stree.queryTree(0, 0, stree.arrSize-1, left, right)
}

func (stree *Segtree) queryTree(node int, start int, end int, l int, r int) int {
	if r < start || l > end {
		return 0
	} else if start == end {
		return stree.tree[node]
	} else {
		mid := (start + end) / 2
		leftNode := 2*node + 1
		rightNode := 2*node + 2
		leftSum := stree.queryTree(leftNode, start, mid, l, r)
		rightSum := stree.queryTree(rightNode, mid+1, end, l, r)
		return leftSum + rightSum
	}
}

func (stree *Segtree) Update(idx int, val int) {
	stree.array[idx] = val
	stree.updateTree(0, 0, stree.arrSize-1, idx, val)
}

func (stree *Segtree) updateTree(node int, start int, end int, idx int, val int) {
	if start == end {
		stree.tree[node] = val
	} else {
		mid := (start + end) / 2
		leftNode := 2*node + 1
		rightNode := 2*node + 2
		if idx >= start && idx <= mid {
			stree.updateTree(leftNode, start, mid, idx, val)
		} else {
			stree.updateTree(rightNode, mid+1, end, idx, val)
		}
		stree.tree[node] = stree.tree[leftNode] + stree.tree[rightNode]
	}
}
